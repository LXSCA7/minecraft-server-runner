package main

import (
	"fmt"
	"os"
	"script-server/server"
	"script-server/utils"
)

// Códigos ANSI para cores
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

var serverRunning = server.CheckServerInstance()

type Config struct {
	VMUser      string `json:"vm_user"`
	VMIP        string `json:"vm_ip"`
	KeyPath     string `json:"key_path"`
	ServerPath  string `json:"server_path"`
	JavaCommand string `json:"java_command"`
}

func main() {
	args := os.Args
	if len(args) > 1 {
		if args[1] == "start" {
			server.StartServer()
			os.Exit(0)
		}

		if args[1] == "stop" {
			server.StopServer()
			os.Exit(0)
		}

		if args[1] == "help" {
			utils.Help()
			os.Exit(0)
		}
	}

	serverRunning = server.CheckServerInstance()
	var option int
	for {
		utils.Menu(server.CheckServerInstance())
		fmt.Print("Opção: ")
		fmt.Scanln(&option)
		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

		if option == 1 {
			fmt.Print("\n")
			server.StartServer()
			fmt.Print("\n")
		} else if option == 2 {
			fmt.Print("\n")
			server.StopServer()
			fmt.Print("\n")
		} else if option == 3 {
			fmt.Println("Saindo...")
			os.Exit(0)
		} else {
			fmt.Println("Opção inválida.")
		}
	}
}
