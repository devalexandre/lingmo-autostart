package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const serviceContent = `[Unit]
Description=Lingmo Service

[Service]
ExecStart=/usr/local/bin/lingmoos run
Restart=always
User=%i

[Install]
WantedBy=default.target
`

func main() {
	if len(os.Args) > 1 && os.Args[1] == "install" {
		install()
	} else if len(os.Args) > 1 && os.Args[1] == "run" {
		run()
	} else {
		fmt.Println("Uso: lingmoos [install|run]")
	}
}

func install() {
	// Caminho do binário
	binaryPath := "/usr/local/bin/lingmoos"

	// Compile o binário
	cmd := exec.Command("go", "build", "-o", binaryPath)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Erro ao compilar o código Go: %v", err)
	}

	// Crie o diretório de configuração, se não existir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Erro ao obter o diretório home do usuário: %v", err)
	}
	configDir := filepath.Join(homeDir, ".config", "lingmoos")
	os.MkdirAll(configDir, 0755)

	// Crie o arquivo de serviço systemd
	servicePath := "/etc/systemd/system/lingmo.service"
	err = os.WriteFile(servicePath, []byte(serviceContent), 0644)
	if err != nil {
		log.Fatalf("Erro ao escrever o arquivo de serviço systemd: %v", err)
	}

	// Recarregue o systemd
	err = exec.Command("sudo", "systemctl", "daemon-reload").Run()
	if err != nil {
		log.Fatalf("Erro ao recarregar o systemd: %v", err)
	}

	// Habilite e inicie o serviço
	err = exec.Command("sudo", "systemctl", "enable", "lingmo.service").Run()
	if err != nil {
		log.Fatalf("Erro ao habilitar o serviço: %v", err)
	}
	err = exec.Command("sudo", "systemctl", "start", "lingmo.service").Run()
	if err != nil {
		log.Fatalf("Erro ao iniciar o serviço: %v", err)
	}

	fmt.Println("Lingmo Service instalado e iniciado com sucesso.")
}

func run() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Erro ao obter o diretório home do usuário: %v", err)
	}

	configFilePath := filepath.Join(homeDir, ".config", "lingmoos", "lingmoos")

	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo de configuração: %v", err)
	}

	fmt.Printf("Conteúdo do arquivo de configuração:\n%s\n", content)
}
