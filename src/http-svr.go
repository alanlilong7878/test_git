

package main

import (
    "io"
    "net/http"
    "log"
    "io/ioutil"
	"fmt"
)


// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	
	//请求消息
	fmt.Println(req.Header)
	recv, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(recv))
	
	
	
	w.WriteHeader(200)
    io.WriteString(w, "http-svr1:hello, world!\n")
}


func main() {
    http.HandleFunc("/in-hello", HelloServer)
    err := http.ListenAndServe(":12347", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


