package main

import (
	_ "fmt"
	cmd "go-todo/cmd"
	cnf "go-todo/config"
	"log"
)

func main() {
	config, err := cnf.GetConfig()
	if err != nil {
		log.Fatalf("Config loading failed : %v", err)
		return
	}
	cmd.Execute(config)
}