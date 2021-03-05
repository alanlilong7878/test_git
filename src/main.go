

package main

import (
    "io"
    "net/http"
    "log"
    "fmt"
	"io/ioutil"
	"os"
)


// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	
	//打印请求消息
	fmt.Println(req.Header)
	fmt.Println(req.Header["Accept"])
	
	recv, err := ioutil.ReadAll(req.Body)
	fmt.Println(string(recv))
	
	incoming_headers := []string{
        "X-Request-Id",
        "x-ot-span-context",
        "x-datadog-trace-id",
        "x-datadog-parent-id",
        "x-datadog-sampling-priority",
        "traceparent",
        "tracestate",
        "x-cloud-trace-context",
        "grpc-trace-bin",
        "X-B3-Traceid",
        "X-B3-Spanid",
        "X-B3-Parentspanid",
        "X-B3-Sampled",
        "x-b3-flags",
		"user-agent" }
	  
    
	//调用其它http服务  Post
	conn := os.Getenv("MY_POD_IP")
    if conn == "" {
		conn = "10.96.173.240"
    }
    rawUrl := "http://"+conn+":12347/in-hello"
    
    client := &http.Client{}
	reqClient, err := http.NewRequest("POST", rawUrl,nil)
    if err != nil {
        // handle error
    }
    reqClient.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    
    for i, s := range incoming_headers {
    	fmt.Println(i, s)
    	if req.Header.Get(s) != "" {
		    reqClient.Header.Set(s,req.Header.Get(s))
    	}		
    }
    resp, err := client.Do(reqClient)
    defer resp.Body.Close()
 
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
 
    fmt.Println(string(body))
	
	w.WriteHeader(200)
    io.WriteString(w, string(body))
}


func main() {
    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServe(":12346", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


