package main

import (
	"fmt"
	"hids/agent/client"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: agent serverApi")
		fmt.Println("Example: agent 8.8.8.8:8080")
		return
	}
	var agent client.Agent
	agent.Init()
	agent.Run()
}
