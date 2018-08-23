package network

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"slg_game_server/server/global"
	"slg_game_server/config/dataConfig"
	"slg_game_server/server/game/user"
)

func Help(args []string) int32 {
	fmt.Println(`
	Commands:
		help
		playerid
		playernum 
		getplayer <int32>
		attrconfig <int32>
		setplayer <int32><int32>
		broadcast
		kickplayer <int32>
		killserver
		
	`)
	return 0
}

func GetCommandHandlers() map[string]func(args []string) int32 {
	return map[string]func([]string) int32{
		"help":       Help,
		"playerid":   ListPlayerId,
		"playernum":  GetPlayersNum,
		"getplayer":  GetPlayer,
		"setplayer":  SetPlayer,
		"attrconfig": AttrConfig,
		"kickplayer": KickPlayer,
		"broadcast":  BroadCast,
		"killserver": KillServer,
	}
}

func AttrConfig(args []string) int32 {
	if len(args) != 2 {
		return 0
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("<id> should be an integer.")
		return 0
	}
	data := dataConfig.GetCfgPlayerAttr(int32(id)).InitData
	fmt.Println(data)
	return 0
}

func SetPlayer(args []string) int32 {
	if len(args) != 3 {
		return 0
	}
	userId, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("<userid> should be an integer.")
		return 0
	}
	power, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("<power> should be an integer.")
		return 0
	}
	player := global.GloInstance.GetPlayer(int32(userId))
	if player != nil {
		p := player.(*user.Player)
		fmt.Println("-------------- player: ", p.Attr.Power)
		p.SetPower(int32(power))
	}
	return 0
}

func ListPlayerId(args []string) int32 {
	idList := global.GloInstance.GetPlayerIdList()
	fmt.Println("-------------- idlist: ", idList)
	return 0
}

func GetPlayersNum(args []string) int32 {
	playerNum, caleNum := global.GloInstance.GetPlayersNum()
	fmt.Printf("--------- mapNum: %d, caleNum: %d\n", playerNum, caleNum)

	return 0
}

func GetPlayer(args []string) int32 {
	if len(args) != 2 {
		return 0
	}
	userId, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("<userid> should be an integer.")
		return 0
	}
	player := global.GloInstance.GetPlayer(int32(userId))
	if player != nil {
		p := player.(*user.Player)
		fmt.Printf("playerinfo: userid: %v, username: %v, level: %v\n", p.Attr.UserId, p.Attr.AcctName, p.Attr.Level)
	}
	return 0
}

func KickPlayer(args []string) int32 {
	if len(args) != 2 {
		return 0
	}
	userId, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("<userid> should be an integer.")
		return 0
	}
	player := global.GloInstance.GetPlayer(int32(userId))
	if player != nil {
		p := player.(*user.Player)
		p.DestoryPlayer()
	}
	return 0

}

func BroadCast(args []string) int32 {
	global.GloInstance.BroadCast()
	return 0
}

func KillServer(args []string) int32 {
	close(Server.Done)
	return 1
}

func Command() {
	Help(nil)
	r := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandlers()
	for {
		fmt.Print("Command > ")
		b, _, _ := r.ReadLine()
		line := string(b)

		tokens := strings.Split(line, " ")

		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("", tokens[0])
		}
	}
}
