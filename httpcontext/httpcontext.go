package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

func Process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	//req, err := http.NewRequest("GET", "http://google.com", nil)
	req, err := http.NewRequest("GET", "https://www.baidu.com/", nil)
	if err != nil {
		fmt.Println("http new request failed", err.Error())
		return
	}
	go func() {
		resp, err1 := client.Do(req)
		pack := Result{r: resp, err: err1}
		c <- pack
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		packtmp := <-c
		fmt.Println("Time out!")
		fmt.Println("Request error is ", packtmp.err.Error())

	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response : %s \n", out)
	}
}

func main() {
	Process()
}
