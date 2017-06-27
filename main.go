package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/oldtree/utils/ServeReport"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	se := &ServeReport.ReportServe{}
	se.Start("127.0.0.1:8899")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-sc
}
