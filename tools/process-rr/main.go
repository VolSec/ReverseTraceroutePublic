package main

import (
	"os"

	"github.com/VolSec/ReverseTraceroutePublic/tools/process-rr/processrr"
)

func main() {
	os.Exit(processrr.Main())
}
