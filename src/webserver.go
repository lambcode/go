package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(output http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
	if(request.URL.Path == "/"){
		request.URL.Path = "/index.html"
	}
	file, error := os.Open("html"+request.URL.Path)
	if error != nil && error !=io.EOF {
		
		panic(error)
	}
	buffer := make([]byte, 2048)
	
	for {
		n,error := file.Read(buffer)
		if error != nil && error !=io.EOF {
			panic(error)
		}
		if n == 0 {
			break
		}
		io.WriteString(output, string(buffer[:n]))
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
