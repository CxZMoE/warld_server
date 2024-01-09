package main

import (
	"bytes"
	"fmt"
	"log"
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
	init_dispatch_pool()
	init_player_management()

	conn, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
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
			init_player_dispatch(p_id)
			// 检查数据帧
			{
				frame := buf[:n]
				log.Println("received frame:", frame)
				// 检查帧格式正确
				if bytes.HasPrefix(frame, []byte{FRAME_HEAD}) && bytes.HasSuffix(frame, []byte{FRAME_END}) {
					log.Println("precess frame:", frame)
					cmd_id := get_frame_cmd(frame)
					log.Println("process frame command:", cmd_id)
					switch cmd_id {
					case WD_CMD_GET_MATCH_INFO:
						log.Println("WD_CMD_GET_MATCH_INFO")
						err := process_get_match_info(p_id)
						if err != nil {
							panic(err)
						}
					}
				}
			}

			// 分发游戏信息
			dispatch_game_msg()
		}
	}()
}

func main() {
	NewGameServer(9010)
	log.Println("server started")
	fmt.Scanln()
}
