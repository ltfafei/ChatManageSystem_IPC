//-*- coding: UTF-8 -*-
// Author: afei00123

package main

import (
	"ChatManageSystem_IPC/cg"
	"ChatManageSystem_IPC/ipc"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//调用cg.CenterClient
var centerClient *cg.CenterClient

func startCenterService() error {
	server := ipc.NewIpcServer(&cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{client}
	return nil
}

//帮助菜单
func Help(arge []string) int {
	fmt.Println(`
Commands:
	login <username> <level> <exp>
	logout <username>
	send <message>
	listplayer
	quit(q)
	help(h)
	`)
		return 0
}

//退出
func Quit(args []string) int {
	return 1
}

//登出
func Logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("USAGE: logout <username>")
		return 0
	}
	//调用登出方法
	centerClient.RemovePlayer(args[1])
	return 0
}

//登录
func Login(args []string) int {
	if len(args) != 4 {
		fmt.Println("USAGE: login <username> <level> <exp>")
		return 0
	}

	//Atoi()函数用于将字符串类型level的整数转换为int类型level
	level, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("无效的参数: <level>应该是一个int类型")
		return 0
	}
	//将字符串类型exp的整数转换为int类型exp
	exp, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("无效的参数: <exp>应该是一个int类型")
		return 0
	}

	player := cg.NewPlayer()
	player.Name = args[1]
	player.Level = level
	player.Exp = exp

	//添加用户异常处理
	err = centerClient.AddPlayer(player)
	if err != nil {
		fmt.Println("添加用户失败", err)
	}
	return 0
}

//列出用户
func ListPlayer(args []string) int {
	//调用centerClient.ListPlayer()方法并做异常处理
	ps, err := centerClient.ListPlayer("")
	if err != nil {
		fmt.Println("列出当前用户失败", err)
	} else {
		for i, v := range ps {
			fmt.Println(i + 1, ":", v)
		}
	}
	return 0
}

//发送消息
func Send(args []string) int {
	//将字符串切片中存在的所有元素连接为单个字符串
	message := strings.Join(args[1:], " ")
	err := centerClient.Broadcast(message)
	if err != nil {
		fmt.Println("发送消息失败", err)
	}
	return 0
}

//将指令和处理函数一一对应
func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func([]string) int{
		"help":       Help,
		"h":          Help,
		"quit":       Quit,
		"q":          Quit,
		"login":      Login,
		"logout":     Logout,
		"listplayer": ListPlayer,
		"send":       Send,
	}
}

func main() {
	fmt.Println("休闲服务器解决方案Demo")
	startCenterService()
	Help(nil)
	//接收标准输入
	r := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandlers()

	//循环读取用户输入
	for {
		//定义指令执行终端标签
		fmt.Print("Command>")
		b, _, _ := r.ReadLine()
		line := string(b)
		tokens := strings.Split(line, " ")
		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0{
				break
			}
		} else {
			fmt.Println("未知指令，请重新输入", tokens[0])
		}
	}
}