package httplogic

import (
	"fmt"
	"net/http"
)

func UsrInfoReq(w http.ResponseWriter, r *http.Request) {

	//打印请求的方法

	fmt.Println("method", r.Method)

	if r.Method == "GET" {
		w.Write([]byte("server receive get method, message is hello"))
	} else {

		//否则走打印输出post接受的参数username和password

		fmt.Println(r.PostFormValue("username"))

		fmt.Println(r.PostFormValue("password"))
		fmt.Println("server receive post method, message is hello")
	}

}

func RegUsrInfo(pattern string) {

	http.HandleFunc(pattern, UsrInfoReq)
}
