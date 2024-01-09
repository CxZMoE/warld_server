package main

import (
	"fmt"
	"log"
	"net"

	"github.com/CxZMoE/warld_server/m/nova"
)

var players map[string]Player

// 初始化玩家玩家管理系统
func init_player_management() {
	players = make(map[string]Player)
}

// 检测并添加新的玩家
func check_register_player(p_id string, addr net.UDPAddr) (bool, error) {
	if _, ok := players[p_id]; !ok {
		players[p_id] = Player{
			connection: addr,
			Sprite: Sprite{
				Scale:    nova.NewVector2D(1, 1),
				Position: nova.NewVector2D(0, 0),
				Rotate:   nova.NewVector2D(0, 0),
				Speed:    nova.NewVector2D(5, 5),
			},
			Name: p_id,
		}
		log.Println("add player:", p_id)
		return true, nil
	}
	return false, fmt.Errorf("player with name: %s already exists", p_id)
}
