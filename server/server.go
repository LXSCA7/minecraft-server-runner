package server

import (
	"fmt"
	"os"
	"script-server/config"
	"script-server/utils"

	"strings"

	"golang.org/x/crypto/ssh"
)

func Connect() (*ssh.Client, *ssh.Session, config.Config, error) {
	cfg, err := config.GetConfigFile("settings.json")
	if err != nil {
		return nil, nil, config.Config{}, fmt.Errorf("erro ao carregar o arquivo de configurações: %v", err)
	}

	signer, err := config.GetSigner(cfg)
	if err != nil {
		return nil, nil, config.Config{}, fmt.Errorf("erro ao processar a chave privada: %v", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: cfg.VMUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", cfg.VMIP+":22", sshConfig)
	if err != nil {
		return nil, nil, config.Config{}, fmt.Errorf("erro ao conectar ao servidor: %v", err)
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, config.Config{}, fmt.Errorf("erro ao criar a sessão: %v", err)
	}

	return client, session, cfg, nil
}

func CheckServerInstance() bool {
	client, session, _, err := Connect()
	if err != nil {
		fmt.Println(utils.Red, "Erro ao criar a sessão para verificar instância:", err, utils.Reset)
		return false
	}
	defer client.Close()
	defer session.Close()

	command := "screen -list"

	output, _ := session.CombinedOutput(command)
	return strings.Contains(string(output), "minecraft")
}

func StartServer() {
	serverRunning := CheckServerInstance()
	if serverRunning {
		fmt.Println(utils.Yellow, "O servidor já está rodando.", utils.Reset)
		return
	}
	client, session, cfg, err := Connect()
	if err != nil {
		fmt.Println(utils.Red, "Erro ao conectar ao servidor:", err, utils.Reset)
		return
	}
	defer client.Close()
	defer session.Close()

	command := fmt.Sprintf("screen -dmS minecraft bash -c 'cd %s && %s", cfg.ServerPath, cfg.JavaCommand+"'")
	fmt.Println("Iniciando o servidor...")

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	if err := session.Run(command); err != nil {
		fmt.Println(utils.Red, "Erro ao iniciar o servidor:", err, utils.Reset)
		return
	}

	serverRunning = true
	fmt.Print(utils.Green)
	fmt.Println("Servidor iniciado com sucesso!", utils.Reset)
}

func StopServer() {
	if !CheckServerInstance() {
		fmt.Println(utils.Yellow, "O servidor já está parado.", utils.Reset)
		return
	}

	client, session, _, err := Connect()
	if err != nil {
		fmt.Println(utils.Red, "Erro ao conectar ao servidor:", err, utils.Reset)
		return
	}
	defer client.Close()
	defer session.Close()

	command := "screen -S minecraft -X quit"
	fmt.Println("Parando o servidor...")

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	if err := session.Run(command); err != nil {
		fmt.Println(utils.Red, "Erro ao parar o servidor:", err, utils.Reset)
		return
	}

	fmt.Print(utils.Green)
	fmt.Println("Servidor parado com sucesso!", utils.Reset)
}
