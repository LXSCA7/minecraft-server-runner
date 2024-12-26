package utils

import (
	"fmt"
)

func Menu(serverRunning bool) {
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	printServerStatus(serverRunning)
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("Selecione uma opção:")
	fmt.Println("[1] Iniciar servidor\n[2] Parar servidor\n[3] Sair")
}

func printServerStatus(serverRunning bool) {
	if serverRunning {
		fmt.Println("STATUS DO SERVIDOR:", Green, "ONLINE", Reset)
	} else {
		fmt.Println("STATUS DO SERVIDOR:", Red, "OFFLINE", Reset)
	}
}
