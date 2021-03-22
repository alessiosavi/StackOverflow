package main

import (
	"fmt"
	"github.com/zhksoftGo/Packet"
	"github.com/zhksoftGo/Packet/protocol/Cactus"
)

func main() {

	var p Packet.Packet
	var m Cactus.MapIntInt = make(Cactus.MapIntInt)
	m.Write(&p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%+v\n", m)
}
