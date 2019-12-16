package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Contractb struct {
	UserId    int64  `json:"userId,omitempty"`
	Projectid int64  `json:"projectid,omitempty"`
	Name      string `json:"name,omitempty"`
	Desc      string `json:"desc,omitempty"`
}
type ContractReq struct {
	Account  string     `json:"account,omitempty"`
	Path     string     `json:"path,omitempty"`
	Nodetype int32      `json:"nodetype,omitempty"`
	Detail   *Contractb `json:"detail,omitempty"`
}

func main() {

	pContractb := new(Contractb)
	pContractb.UserId = 1
	pContractb.Projectid = 1
	pContractb.Name = "testContract"
	pContractb.Desc = "test Contract desc"
	contractReq := &ContractReq{Account: "eosio", Path: "/data/cai/contract", Nodetype: 0, Detail: pContractb}
	data, err := json.Marshal(contractReq)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))

	url := "https://0.0.0.0:8195/contract/create"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		panic(err)

	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)

	fmt.Println("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

}
