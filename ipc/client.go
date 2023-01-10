//-*- coding: UTF-8 -*-
// Author: afei00123

package ipc

import "encoding/json"

type IpcClient struct {
	//声明conn转换为chan string
	conn chan string
}

//连接Server服务器
func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

//定义Call函数，将消息封装成JSON格式的字符串发送给Server
func (client *IpcClient)Call(method, params string) (resp *Response, err error) {
	req := &Request{method, params}
	var b []byte
		b, err = json.Marshal(req)
		if err != nil {
			return
		}
		client.conn <- string(b)
		//等待返回值
		str := <- client.conn

		var resp1 Response
		err = json.Unmarshal([]byte(str), &resp1)
		resp = &resp1
		return
}

//客户端关闭连接
func (client *IpcClient)Close() {
	client.conn <- "CLOSE"
}