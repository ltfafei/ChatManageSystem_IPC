//-*- coding: UTF-8 -*-
// Author: afei00123

package ipc

import "testing"

//定义一个空结构体。空结构体的初衷只有一个：节省内存
type EchoServer struct {
}

func (server *EchoServer) Handle(request string) string {
	return "Echo: " + request
}

func (server *EchoServer) Name() string {
	return "EchoServer successfully"
}

//单元测试，测试IPC通信
func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1 := client1.Call("From Client1")
	resp2 := client2.Call("From Client2")

	if resp1 != "Echo: From Client1" || resp2 != "Echo: From Client2" {
		t.Error("IpcClient Call failed. resp1: ", resp1, "resp2: ", resp2)
	}
	client1.Close()
	client2.Close()
}