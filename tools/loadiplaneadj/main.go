package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/VolSec/ReverseTraceroutePublic/dataaccess"
	"github.com/VolSec/ReverseTraceroutePublic/iplanetraceroute"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: loadiplaneadj <dir>")
		os.Exit(1)
	}
	var dir = os.Args[1]
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var dc dataaccess.DbConfig
	var conf dataaccess.Config
	conf.Host = "localhost"
	conf.Db = "revtr"
	conf.Password = "password"
	conf.Port = "3306"
	conf.User = "revtr"
	dc.ReadConfigs = append(dc.ReadConfigs, conf)
	dc.WriteConfigs = append(dc.WriteConfigs, conf)
	da, err := dataaccess.New(dc)
	if err != nil {
		panic(err)
	}
	defer da.Close()
	for _, file := range files {
		fmt.Println(file)
		if file.IsDir() {
			continue
		}
		err := processFile(file, da)
		if err != nil {
			fmt.Println("Error in file: ", file.Name())
			fmt.Println(err)
		}
	}
}

func processFile(file os.FileInfo, da *dataaccess.DataAccess) error {
	f, err := os.Open(file.Name())
	if err != nil {
		return err
	}
	defer f.Close()
	scan := iplane.NewTracerouteScanner(f)
	for scan.Scan() {
		trace := scan.Traceroute()
		for i, hop := range trace.Hops {
			if i+1 >= len(trace.Hops) {
				break
			}
			if hop.IP.String() == "0.0.0.0" {
				continue
			}
			if trace.Hops[i+1].IP.String() == "0.0.0.0" {
				continue
			}
			err := da.StoreAdjacency(hop.IP, trace.Hops[i+1].IP)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return scan.Err()
}
