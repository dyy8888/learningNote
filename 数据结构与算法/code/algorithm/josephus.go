package algorithm

import (
	"container/ring"
	"fmt"
)

// 设定游戏固定配置
const (
	totalPlayer   = 100
	startPosition = 1
	deadline      = 3
)

type Player struct {
	alive bool
	pos   int
}

func StartPlay() {
	rings := ring.New(totalPlayer)
	for index := startPosition; index <= totalPlayer; index++ {
		rings.Value = &Player{
			pos:   index,
			alive: true,
		}
		rings = rings.Next()
	}
	if startPosition > 1 {
		rings.Move(startPosition - 1)
	}
	count := 1
	deadcount := 0
	rev := rings.Prev()
	for deadcount < totalPlayer {
		if count == deadline {
			rev.Link(rev.Move(2))
			fmt.Println("死亡", rings.Value)
			rings = rev.Next()
			deadcount++ //死亡人数
			count = 0
			count++
		} else {
			rev = rings
			rings = rings.Next()
			count++
		}
	}
}
