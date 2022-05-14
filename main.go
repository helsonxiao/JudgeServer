package main

import "github.com/helsonxiao/JudgeServer/server"

func main() {
	r := server.SetupRouter()
	r.Run(":8000")
}
