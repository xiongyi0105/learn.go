package main

/*
1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
当访问 localhost/healthz 时，应返回200
*/
import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func index(response http.ResponseWriter, r *http.Request) {

	io.WriteString(response, "<h1>Welcome to cloud native</h1>")
	for k, v := range r.Header {
		for _, vv := range v {
			response.Header().Set(k, vv)
		}
	}
	os.Setenv("VERSION", "0.0.1")
	Version := os.Getenv("VERSION")
	fmt.Println(Version)
	response.Header().Set("Version", Version)
	fmt.Println(response.Header())
	//log.Printf("ClientIp is : %s",r.RemoteAddr)
	log.Printf("ClientIp is : %s\n", getClientIp(r))
	log.Printf("Status: %d", http.StatusOK)
}

func getClientIp(r *http.Request) string {
	XForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(XForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func main() {
	if err := flag.Set("v", "4"); err == nil {
		glog.V(4).Info("Start http server success")
	}
	//起一个http服务为8080端口
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("start server failed, %s\n", err.Error())
	}
}
