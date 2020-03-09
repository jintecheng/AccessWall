package main

import (
	"fmt"
	"os"

	R "github.com/jintecheng/accesswall/routes"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage : <app> <databas ip> <database port>")
		os.Exit(0)
	}
	R.InitDB(args[1], args[2])
	R.Server()
}
