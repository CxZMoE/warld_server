package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	// 添加玩家
	WD_CMD_ADD_PLAYER = 0x01
	// 删除玩家
	WD_CMD_DEL_PLAYER = 0x02
	// 玩家退出游戏
	WD_CMD_PLAYER_QUIT = 0x03
	// 玩家更新游戏状态
	WD_CMD_PLAYER_UPDATE = 0x04
	// 获取游戏(战局)信息
	WD_CMD_GET_MATCH_INFO = 0x05
)

const (
	FRAME_HEAD = 0xff
	FRAME_END  = 0xfb
)

var dispatch_pool map[string](chan []byte)

func init_dispatch_pool() {
	dispatch_pool = make(map[string]chan []byte)
}

func init_player_dispatch(p_id string) {
	dispatch_pool[p_id] = make(chan []byte, 0xff)
}

func get_frame_cmd(frame []byte) int {
	if len(frame) > 1 {
		return int(frame[1])
	}
	return -1
}

func make_package(cmd_id uint8, data []byte) []byte {
	var frame []byte
	var data_size = len(data)
	switch cmd_id {
	case WD_CMD_GET_MATCH_INFO:
		frame = make([]byte, 0, 256)
		frame = append(frame,
			FRAME_HEAD,
			WD_CMD_GET_MATCH_INFO,
			byte(data_size<<8),
			byte(data_size&0x0f),
		)
		if data_size != 0 {
			frame = append(frame, data...)
		}
		frame = append(frame, FRAME_END)
	}
	return frame
}

func process_get_match_info(p_id string) error {
	if _, ok := dispatch_pool[p_id]; !ok {
		return fmt.Errorf("p_id: %s not found in dispatch pool", p_id)
	}
	if len(dispatch_pool[p_id]) < 0xff {
		player_list_data, err := json.Marshal(players)
		if err != nil {
			return err
		}

		log.Println("dispatch msg:", string(player_list_data))
		dispatch_pool[p_id] <- make_package(WD_CMD_GET_MATCH_INFO, player_list_data)
	} else {
		log.Println("length of dispatch_pool:", len(dispatch_pool[p_id]))
	}
	return nil
}

func dispatch_game_msg() {
	for _, player := range players {
		dispatch := dispatch_pool[player.Name]
		if dispatch != nil {
			if len(dispatch) > 0 {
				serverSocket.WriteToUDP(<-dispatch, &player.connection)
				// TODO: record errors
			}
		}
	}
}
