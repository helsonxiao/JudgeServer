package main

import (
	"github.com/helsonxiao/JudgeServer/configs"
	"github.com/helsonxiao/JudgeServer/server"
)

func main() {
	// TODO: setup logger
	configs.InitEnv()
	r := server.SetupRouter()
	r.Run(":8000")
}
