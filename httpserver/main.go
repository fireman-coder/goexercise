package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	//"strings"
	"github.com/golang/glog"
	_ "net/http/pprof"
)

func main() {

	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/header", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/version", GetVeriosn)
	// mux.HandleFunc("/debug/pprof/", pprof.Index)
	// mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}
func healthz(w http.ResponseWriter, r *http.Request) {
	glog.V(2).Info("接口信息：healthz"," ip:",r.RemoteAddr," 状态码",http.StatusOK)
	io.WriteString(w, "200\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	glog.V(2).Info("接口信息：获取请求头"," ip:",r.RemoteAddr," 状态码",http.StatusOK)
	info:=""
	for k, v := range r.Header {
		w.Header().Set(k,strings.Join(v,","))
		info+=strings.Join(v,",")
	}
	io.WriteString(w, info)
	//w.WriteHeader(http.StatusBadRequest)
}

func GetVeriosn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering GetVeriosn")
	glog.V(2).Info("接口信息：GetVeriosn"," ip:",r.RemoteAddr," 状态码",http.StatusOK)
	path:=os.Getenv("GOPATH")
	w.Header().Set("GOPATH", path)
	io.WriteString(w, path)
}