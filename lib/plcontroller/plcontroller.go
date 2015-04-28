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
package plcontroller

import (
	dm "github.com/NEU-SNS/ReverseTraceroute/lib/datamodel"
	"github.com/NEU-SNS/ReverseTraceroute/lib/mproc"
	"github.com/NEU-SNS/ReverseTraceroute/lib/scamper"
	"github.com/NEU-SNS/ReverseTraceroute/lib/util"
	"github.com/go-fsnotify/fsnotify"
	"github.com/golang/glog"
	"net"
	"os"
	"sync"
	"time"
)

type plControllerT struct {
	port      int
	ip        net.IP
	ptype     string
	startTime time.Time
	spid      int
	sc        scamper.ScamperConfig
	mp        mproc.MProc
	w         *fsnotify.Watcher
	mu        sync.Mutex
	//the mutex protects the following
	requests int64
	time     time.Duration

	rw sync.RWMutex
	//rwmutex protext the socks
	socks map[string]scamper.Socket
}

func handleScamperStop(err error, ps *os.ProcessState) bool {
	switch err.(type) {
	default:
		return false
	case *os.PathError:
		return true
	}

}

var plController plControllerT

func (c *plControllerT) getStatsInfo() (t time.Duration, req int64) {
	c.mu.Lock()
	t, req = c.time, c.requests
	c.mu.Unlock()
	return
}

func (c *plControllerT) getStats() dm.Stats {
	utime := time.Since(c.startTime)
	t, req := c.getStatsInfo()
	var tt time.Duration
	if t == 0 {
		tt = 0
	} else {
		tt = time.Duration(req / int64(t))
	}
	s := dm.Stats{StartTime: c.startTime,
		UpTime: utime, Requests: req,
		TotReqTime: t, AvgReqTime: tt}
	return s
}

func (c *plControllerT) addSocket(sock scamper.Socket) {
	c.rw.Lock()
	c.socks[sock.IP()] = sock
	c.rw.Unlock()
}

func (c *plControllerT) removeSocket(sock scamper.Socket) {
	c.rw.Lock()
	delete(c.socks, sock.IP())
	c.rw.Unlock()
}

func Start(n, laddr string, sc scamper.ScamperConfig) chan error {
	errChan := make(chan error, 1)
	port, ip, err := util.ParseAddrArg(laddr)

	if err != nil {
		glog.Error("Failed to parse addr string")
		errChan <- err
		return errChan
	}
	err = scamper.ParseScamperConfig(sc)
	if err != nil {
		glog.Errorf("Invalid scamper args: %v", err)
		errChan <- err
		return errChan
	}
	plController.startTime = time.Now()
	plController.ip = ip
	plController.port = port
	plController.mp = mproc.New()
	plController.sc = sc
	plController.startScamperProc()

	go util.StartRpc(n, laddr, errChan, new(PlControllerApi))
	return errChan
}

func (c *plControllerT) startScamperProc() {
	sp := scamper.GetProc(c.sc.Path, c.sc.Port, c.sc.ScPath)
	plController.mp.ManageProcess(sp, true, 10, handleScamperStop)
}

func HandleSig(s os.Signal) {
	plController.handleSig(s)
}

func (c *plControllerT) handleSig(s os.Signal) {
	if c.mp != nil {
		c.mp.KillAll()
	}
}
