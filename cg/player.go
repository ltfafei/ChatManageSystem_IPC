//-*- coding: UTF-8 -*-
// Author: afei00123

package cg

import "fmt"

//构造动作结构体
type Player struct {
	//游戏玩家昵称
	Name string "name"
	Level int "Level"
	Exp int "exp"
	//房间号
	Room int "room"

	//等待收取的信息
	mq chan *Message
}

type Room struct {
	Room int "room"
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}

	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "received message: ", msg.Content)
		}
	}(player)
	return player
}