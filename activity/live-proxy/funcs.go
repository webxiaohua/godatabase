package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ApiProxyParams struct {
	Params   string `json:"params"`
	AppId    string `json:"appId"`
	Module   string `json:"module"`
	GrpcPath string `json:"grpcPath"`
}

func GetFormParams(apiProxyParams *ApiProxyParams) string {
	return "params=" + apiProxyParams.Params + "&appId=" + apiProxyParams.AppId + "&module=" + apiProxyParams.Module + "&grpcPath=" + apiProxyParams.GrpcPath
}

func HttpPost(url string, params string) string {
	req, _ := http.NewRequest("POST", url, strings.NewReader(params))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[error]", err)
		return "error"
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
	//fmt.Println(string(body))
}
