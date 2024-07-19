package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeoutMilliseconds int = 5000

// barrierResp contains a response message.
type barrierResp struct {
	Err  error
	Resp string
}

// barrier implements a barrier pattern using channels.
func barrier(endpoints ...string) {
	requestNumber := len(endpoints)
	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)

	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("Error: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}
}

// makeRequest makes a GET request for a single endpoint.
func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	res.Resp = string(byt)
	out <- res
}
