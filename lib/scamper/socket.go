/*
 Copyright (c) 2015, Northeastern University
 All rights reserved.

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are met:
     * Redistributions of source code must retain the above copyright
       notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above copyright
       notice, this list of conditions and the following disclaimer in the
       documentation and/or other materials provided with the distribution.
     * Neither the name of the Northeastern University nor the
       names of its contributors may be used to endorse or promote products
       derived from this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 DISCLAIMED. IN NO EVENT SHALL Northeastern University BE LIABLE FOR ANY
 DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Package scamper is a library to work with scamper control sockets
package scamper

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/NEU-SNS/ReverseTraceroute/lib/util"
	"github.com/golang/glog"
)

type stringReadWriter interface {
	io.Reader
	io.Writer
	ReadString(delim byte) (line string, err error)
}

var (
	// ErrorCmdNotFound returned when no cmd is found in the cmdMap
	ErrorCmdNotFound = fmt.Errorf("No command found matching given Id")
	// ErrorDupCommand returned when a socket as a cmd with the same id already
	// running
	ErrorDupCommand = fmt.Errorf("Command already exists with the give Id")
)

type cmdMap struct {
	sync.Mutex
	cmds map[uint32]cmdResponse
}

type cmdResponse struct {
	cmd  *Cmd
	done <-chan Response
}

func (cm *cmdMap) getCmd(id uint32) (*Cmd, error) {
	cm.Lock()
	defer cm.Unlock()
	if cmd, ok := cm.cmds[id]; ok {
		return cmd.cmd, nil
	}
	return nil, ErrorCmdNotFound
}

func (cm *cmdMap) rmCmd(id uint32) {
	cm.Lock()
	defer cm.Unlock()
	delete(cm.cmds, id)
}

func (cm *cmdMap) addCmd(c cmdResponse) error {
	cm.Lock()
	defer cm.Unlock()
	if _, ok := cm.cmds[c.cmd.userID]; ok {
		return ErrorDupCommand
	}
	cm.cmds[c.cmd.userID] = c
	return nil
}

func newCmdMap() *cmdMap {
	m := make(map[uint32]cmdResponse)
	return &cmdMap{cmds: m}
}

type userID struct {
	UserID uint32 `json:"user_id"`
}

type unmarshal func(data []byte, v interface{}) error

// Socket represents a scamper control socket
type Socket struct {
	fname         string
	ip            string
	port          string
	converterPath string
	closeChan     chan struct{}
	errChan       chan error
	cmdChan       chan *Cmd
	respChan      chan Response
	cmds          *cmdMap
	con           net.Conn
	wartsHeader   [2]Response
	rc            uint32
	userID        uint32
	unmarsh       unmarshal
}

// NewSocket creates a new scamper socket
func NewSocket(fname, cPath string, um unmarshal) *Socket {
	con, err := net.Dial("unix", fname)
	if err != nil {
		return nil
	}
	cc := make(chan *Cmd, 10)
	rc := make(chan Response, 10)
	clc := make(chan struct{})
	sock := &Socket{
		fname:         fname,
		cmds:          newCmdMap(),
		cmdChan:       cc,
		respChan:      rc,
		closeChan:     clc,
		con:           con,
		converterPath: cPath,
		unmarsh:       um,
	}

	go sock.monitorConn()
	return sock
}

// Stop closes the connection the socket represents
func (s *Socket) Stop() {
	close(s.closeChan)
}

func (s *Socket) monitorConn() {
	rw := bufio.NewReadWriter(bufio.NewReader(s.con), bufio.NewWriter(s.con))

	for {
		select {
		case c := <-s.cmdChan:
			err := c.issueCommand(s.con)
			glog.Errorf("Error issuing command %s", c.marshal())
			c.Err = err
		case <-s.closeChan:
			s.con.Close()
			return
		default:
			line, err := rw.ReadString('\n')

			if err != nil {
				glog.Errorf("Error reading from socket: %v", err)
				s.con.Close()
				if s.con, err = net.Dial("unix", s.fname); err != nil {
					glog.Error("Failed to reopen connection")
					return
				}
				s.rc = 0
				rw = bufio.NewReadWriter(bufio.NewReader(s.con), bufio.NewWriter(s.con))
				continue
			}
			resp, err := parseResponse(line, rw)
			if err != nil {
				glog.Errorf("Error parsing response: %s", line)
				continue
			}
			if resp.rType != DATA {
				continue
			}
			s.rc++
			if s.rc < 2 {
				s.wartsHeader[s.rc] = resp
				continue
			}
			dec := &util.UUDecodingWriter{}
			for i := 0; i < 2; i++ {
				s.wartsHeader[i].WriteTo(dec)
			}
			uid := &userID{}
			resp.WriteTo(dec)
			err = s.unmarsh(dec.Bytes(), uid)
			if err != nil {
				glog.Error("Could not parse UserId from response")
				continue
			}
			s.respChan <- resp
		}
	}
}

func (s *Socket) getID() uint32 {
	id := s.userID
	s.userID++
	return id
}

// DoMeasurement perform the measurement described by arg
func (s *Socket) DoMeasurement(arg interface{}) (<-chan Response, error) {
	cmd, err := newCmd(arg, s.getID())
	if err != nil {
		return nil, err
	}
	cr := cmdResponse{cmd: &cmd, done: make(<-chan Response, 1)}
	err = s.cmds.addCmd(cr)
	if err != nil {
		return nil, err
	}
	s.cmdChan <- &cmd
	return cr.done, err
}

// IP Gets the ip of the remote machine that is connected to the socket
func (s *Socket) IP() string {
	if s.ip == "" {
		s.ip = strings.Split(path.Base(s.fname), ":")[util.IP]
		return s.ip
	}
	return s.ip
}

// Port gets the port of the remote machine that is connected to the socket
func (s *Socket) Port() string {
	if s.port == "" {
		s.port = strings.Split(path.Base(s.fname), ":")[util.PORT]
		return s.port
	}
	return s.port
}

func parseResponse(r string, rw *bufio.ReadWriter) (Response, error) {
	resp := Response{}
	switch {
	case strings.Contains(r, string(OK)):
		resp.rType = OK
		r = strings.TrimSpace(r)
		split := strings.Split(r, " ")
		idsp := strings.Split(split[1], "-")
		_, err := strconv.Atoi(idsp[1])
		if err != nil {
			return resp, ErrorBadResponse
		}
		return resp, nil
	case strings.Contains(r, string(ERR)):
		resp.rType = ERR
		return resp, nil
	case strings.Contains(r, string(DATA)):
		resp.rType = DATA
		split := strings.Split(r, " ")
		if len(split) != 2 {
			return resp, ErrorBadDataResponse
		}
		n, err := strconv.Atoi(split[1][:len(split[1])-1])
		if err != nil {
			return resp, err
		}
		resp.ds = n
		buff := make([]byte, n)
		_, err = io.ReadFull(rw, buff)
		if err != nil {

			return resp, err
		}
		resp.data = buff
		return resp, nil
	case strings.Contains(r, string(MORE)):
		resp.rType = MORE
		return resp, nil
	}
	return resp, ErrorBadResponse
}