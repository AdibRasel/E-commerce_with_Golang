package main

import (
	"BackEnd/cmd"
	"BackEnd/config"
	"fmt"
)

func main() {
	cnf := config.GetConfig()
	fmt.Println(cnf.HttpPort)
	fmt.Println(cnf.ServiceName)
	fmt.Println(cnf.Version)

	cmd.Serve()
}
