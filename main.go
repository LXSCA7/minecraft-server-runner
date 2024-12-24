package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

// Códigos ANSI para cores
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

var serverRunning = false

type Config struct {
	VMUser      string `json:"vm_user"`
	VMIP        string `json:"vm_ip"`
	KeyPath     string `json:"key_path"`
	ServerPath  string `json:"server_path"`
	JavaCommand string `json:"java_command"`
}

func connect() (*ssh.Client, *ssh.Session, error) {
	configFile, err := os.Open("settings.json")
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao abrir o arquivo de configuração: %v", err)
	}
	defer configFile.Close()

	var config Config
	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao ler o arquivo de configuração: %v", err)
	}
	if err := json.Unmarshal(byteValue, &config); err != nil {
		return nil, nil, fmt.Errorf("Erro ao analisar o arquivo de configuração: %v", err)
	}

	key, err := ioutil.ReadFile(config.KeyPath)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao ler a chave privada: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao processar a chave privada: %v", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: config.VMUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", config.VMIP+":22", sshConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("Erro ao conectar ao servidor: %v", err)
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("Erro ao criar a sessão: %v", err)
	}

	return client, session, nil
}

func startServer() {
	if checkServerInstance() {
		fmt.Println(Yellow, "O servidor já está rodando.", Reset)
		return
	}
	client, session, err := connect()
	serverRunning = checkServerInstance()
	if err != nil {
		fmt.Println(Red, "Erro ao conectar ao servidor:", err, Reset)
		return
	}
	defer client.Close()
	defer session.Close()

	command := "screen -dmS minecraft bash -c 'cd /home/ubuntu/server && java -Xmx16G -Xms16G -jar server.jar nogui'"
	fmt.Println("Iniciando o servidor...")

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	if err := session.Run(command); err != nil {
		fmt.Println(Red, "Erro ao iniciar o servidor:", err, Reset)
		return
	}

	serverRunning = true
	fmt.Println(Green, "Servidor iniciado com sucesso!", Reset)
}

func stopServer() {
	if !checkServerInstance() {
		fmt.Println(Yellow, "O servidor já está parado.", Reset)
		return
	}

	client, session, err := connect()
	if err != nil {
		fmt.Println(Red, "Erro ao conectar ao servidor:", err, Reset)
		return
	}
	defer client.Close()
	defer session.Close()

	command := "screen -S minecraft -X quit"
	fmt.Println("Parando o servidor...")

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	if err := session.Run(command); err != nil {
		fmt.Println(Red, "Erro ao parar o servidor:", err, Reset)
		return
	}

	serverRunning = false
	fmt.Println(Green, "Servidor parado com sucesso!", Reset)
}

func printServerStatus() {
	if serverRunning {
		fmt.Println("STATUS DO SERVIDOR:", Green, "ONLINE", Reset)
	} else {
		fmt.Println("STATUS DO SERVIDOR:", Red, "OFFLINE", Reset)
	}
}

func checkServerInstance() bool {
	client, session, err := connect()
	if err != nil {
		fmt.Println(Red, "Erro ao criar a sessão para verificar instância:", err, Reset)
		return false
	}
	defer client.Close()
	defer session.Close()

	command := "screen -list"

	output, _ := session.CombinedOutput(command)
	serverRunning = strings.Contains(string(output), "minecraft")
	return strings.Contains(string(output), "minecraft")
}

func skipLineConsole() {
	fmt.Print("\n")
}

func menu() {
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	printServerStatus()
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("Selecione uma opção:")
	fmt.Println("[1] Iniciar servidor\n[2] Parar servidor\n[3] Sair")
}

func main() {
	serverRunning = checkServerInstance()
	var option int
	for {
		menu()
		fmt.Print("Opção: ")
		fmt.Scanln(&option)
		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

		if option == 1 {
			skipLineConsole()
			startServer()
			skipLineConsole()
		} else if option == 2 {
			skipLineConsole()
			stopServer()
			skipLineConsole()
		} else if option == 3 {
			fmt.Println("Saindo...")
			os.Exit(0)
		} else {
			fmt.Println("Opção inválida.")
		}
	}
}
