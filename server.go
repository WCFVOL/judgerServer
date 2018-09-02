package main

import (
	"awesomeProject/judgerServer/httpserver"
	"awesomeProject/judgerServer/work"
)

func main() {
	work.DoWork()
	httpserver.Run("", "8080")
}
