//-*- coding: UTF-8 -*-
// Author: afei00123

package ipc

import (
	"encoding/json"
	"fmt"
)

//声明结构体
type Request struct {
	Method string "method"
	Params string "params"
}
type Response struct {
	Code string "code"
	Body string "body"
}

//定义接口
type Server interface {
	//定义方法以及声明方法类型
	Name() string
	Handle(methot, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer)Connect() chan string {
	session := make(chan string, 0)
	go func(c chan string) {
		for {
			request := <-c
			//关闭连接标志符
			if request == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("无效的请求：", request)
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			c <- string(b)
		}
		fmt.Println("Session closed.")
	}(session)
	fmt.Println("一个新的会话创建成功.")
	return session
}