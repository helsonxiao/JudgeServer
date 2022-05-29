package main

import (
	"github.com/helsonxiao/JudgeServer/configs"
	"github.com/helsonxiao/JudgeServer/server"
)

func main() {
	// TODO: setup logger
	configs.SetupEnv()
	r := server.SetupRouter()
	r.Run(":8080")
}
