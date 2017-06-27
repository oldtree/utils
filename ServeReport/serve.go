package ServeReport

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/exec"
	"runtime"
	rp "runtime/pprof"
	"time"
)

type DebugServe struct {
	StopChan chan struct{}
}

func (de *DebugServe) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	url := req.URL
	path := url.Path
	switch path {
	case "/debug/pprof/cmdline":
		pprof.Cmdline(resp, req)
	case "/debug/pprof/profile":
		pprof.Profile(resp, req)
	case "/debug/pprof/symbol":
		pprof.Symbol(resp, req)
	case "/debug/pprof/trace":
		pprof.Trace(resp, req)
	default:
		pprof.Index(resp, req)
	}
}

func (de *DebugServe) Start(address string) {
	var err = make(chan error, 1)
	select {
	case err <- http.ListenAndServe(address, http.Handler(de)):
		log.Fatalln(err)
	}

}

var StartTimeStamp = time.Now()
var Pid = os.Getpid()
var Path, _ = os.Getwd()

type ReportTable struct {
	Hostname  string      `json:"hostname,omitempty"`
	TimeStamp int64       `json:"time_stamp,omitempty"`
	Metric    string      `json:"metric,omitempty"`
	PID       int         `json:"pid,omitempty"`
	Dir       string      `json:"dir,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type ReportServe struct {
	StopChan chan struct{}
}

func (re *ReportServe) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	url := req.URL
	path := url.Path
	switch path {
	case "/report/pprof/cpu":
		ReportCpuProfile(w, req)
	case "/report/pprof/heap":
		ReportHeapProfile(w, req)
	case "/report/pprof/block":
		ReportBlockProfile(w, req)
	case "/report/pprof/goroutine":
		ReportGoroutineProfile(w, req)
	case "/report/pprof/env":
		ReportEnvProfile(w, req)
	case "/report/pprof/threadcreate":
		ReportThreadProfile(w, req)
	case "/report/pprof/memery":
		ReportMemeryProfile(w, req)
	default:
		w.Write([]byte("not support commond"))
	}
}

func (re *ReportServe) Start(address string) {
	var err = make(chan error, 1)
	select {
	case err <- http.ListenAndServe(address, http.Handler(re)):
		log.Fatalln(err)
	}
}

func ReportMemeryProfile(w http.ResponseWriter, req *http.Request) {
	filename := fmt.Sprintf("./mem-%d.mempprof", Pid)
	fi, err := os.Create(filename)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	runtime.GC()
	rp.WriteHeapProfile(fi)
	fi.Sync()
	fi.Close()
	toolsPProf := exec.Command("go", "tool", "pprof", "-text", filename)
	toolsPProf.Stderr = w
	toolsPProf.Stdout = w
	err = toolsPProf.Run()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	return
}

func ReportCpuProfile(w http.ResponseWriter, req *http.Request) {
	sec := 30
	filename := fmt.Sprintf("./cpu-%d.pprof", Pid)
	fi, err := os.Create(filename)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	rp.StartCPUProfile(fi)
	time.Sleep(time.Second * time.Duration(sec))
	rp.StopCPUProfile()
	fi.Sync()
	fi.Close()
	toolsPprof := exec.Command("go", "tool", "pprof", "-png", filename)
	toolsPprof.Stdout = w
	toolsPprof.Stderr = w
	err = toolsPprof.Run()
	if err != nil {
		log.Println("report cpu error")
	}
	return
}

func ReportHeapProfile(w http.ResponseWriter, req *http.Request) {
	pro := rp.Lookup("heap")
	pro.WriteTo(w, 2)
	return
}

func ReportBlockProfile(w http.ResponseWriter, req *http.Request) {
	pro := rp.Lookup("block")
	pro.WriteTo(w, 2)
	return
}

func ReportGoroutineProfile(w http.ResponseWriter, req *http.Request) {
	pro := rp.Lookup("goroutine")
	pro.WriteTo(w, 2)
	return
}

func ReportEnvProfile(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("GOPATH : " + os.Getenv("GOPATH") + "GOROOT: " + os.Getenv("GOROOT")))
	return
}

func ReportThreadProfile(w http.ResponseWriter, req *http.Request) {
	pro := rp.Lookup("threadcreate")
	pro.WriteTo(w, 2)
	return
}
