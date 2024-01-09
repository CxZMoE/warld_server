package main

import (
	"bytes"
	"net"

	"github.com/CxZMoE/warld_server/m/nova"
)

type Sprite struct {
	Name string `json:"name"`

	Scale    nova.Vector2D
	Position nova.Vector2D
	Rotate   nova.Vector2D
	Speed    nova.Vector2D
}

type Player struct {
	connection  net.UDPAddr
	Sprite      Sprite `json:"sprite"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var (
	sprites      []*Sprite
	serverSocket *net.UDPConn
)

func NewGameServer(port int) {
	init_player_management()

	conn, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	})
	if err != nil {
		panic(err)
	}
	serverSocket = conn
	go func() {
		buf := make([]byte, 1024)
		for {
			n, addr, err := serverSocket.ReadFromUDP(buf)
			if err != nil {
				panic(err)
			}
			if n == 0 {
				continue
			}

			// 检查并加入玩家
			p_id := addr.IP.String()
			check_register_player(p_id, *addr)

			// 检查数据帧
			{
				// 检查帧格式正确
				if bytes.HasPrefix(buf, []byte{FRAME_HEAD}) && bytes.HasSuffix(buf, []byte{FRAME_END}) {
					cmd_id := get_frame_cmd(buf)
					switch cmd_id {
					case WD_CMD_GET_MATCH_INFO:
						process_get_match_info(p_id)

					}
				}
			}

			// 分发游戏信息
			dispatch_game_msg()
		}
	}()
}
