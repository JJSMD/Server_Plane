package main

import (
	"./core/INet"
	"./server/app"
	_ "./server/game"
	"./server/gameutils"
	"./utils"
	"os"
)


func main() {

	if len(os.Args) > 1 {
		cfgPath := os.Args[1]
		utils.GlobalObject.Load(cfgPath)
	}else{
		utils.GlobalObject.Load("conf/game.json")
	}

	//db.Init()

	s := INet.NewServer()
	//s.AddRouter(&gameutils.STS)
	//s.AddRouter(&game.Enter)

	s.SetOnConnStart(gameutils.ClientConnStart)
	s.SetOnConnStop(gameutils.ClientConnStop)
	app.SetShutDownFunc(gameutils.ShutDown)
	app.SetServer(s)

	//go app.MasterClient(proto.ServerTypeGame)

	s.Running()
}
