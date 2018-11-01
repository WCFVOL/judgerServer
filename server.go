package main

import (
	"awesomeProject/judgerServer/httpserver"
	"awesomeProject/judgerServer/work"
)

func main() {
	go work.DoWork()
	httpserver.Run("", "8000")
}
