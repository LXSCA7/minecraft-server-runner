package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

type Config struct {
	VMUser      string `json:"vm_user"`
	VMIP        string `json:"vm_ip"`
	KeyPath     string `json:"key_path"`
	ServerPath  string `json:"server_path"`
	JavaCommand string `json:"java_command"`
}

func GetConfigFile(path string) (Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("erro ao abrir o arquivo de configuração: %v", err)
	}
	defer configFile.Close()

	var config Config
	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("erro ao ler o arquivo de configuração: %v", err)
	}
	if err := json.Unmarshal(byteValue, &config); err != nil {
		return Config{}, fmt.Errorf("erro ao analisar o arquivo de configuração: %v", err)
	}

	return config, nil
}

func GetSigner(cfg Config) (ssh.Signer, error) {
	key, err := ioutil.ReadFile(cfg.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler a chave privada: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("erro ao processar a chave privada: %v", err)
	}

	return signer, nil
}
