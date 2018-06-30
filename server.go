package main

import (
	"awesomeProject/judgerServer/consumer"
	_ "awesomeProject/judgerServer/consumer"
)

func main() {
	consumer.Consumer()
}
