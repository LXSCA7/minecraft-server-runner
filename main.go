package main

import (
	"fmt"
	"os"
	"script-server/server"
	"script-server/utils"
)

var serverRunning = server.CheckServerInstance()
var version = "development"

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

		if args[1] == "version" {
			fmt.Println("Versão: " + version)
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
