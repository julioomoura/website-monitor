package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Júlio"
	versao := 1.1
	fmt.Println("Olá, sr.", nome, "sua idade é:")
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			fmt.Println("Testando site", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, erro := http.Get(site)

	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso.")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, erro := os.Open("sites.txt")
	if erro != nil {
		fmt.Println("Ocorreu um erro:", erro)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, erro := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if erro == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, erro := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if erro != nil {
		fmt.Println(erro)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, erro := ioutil.ReadFile("log.txt")

	if erro != nil {
		fmt.Println(erro)
	}

	fmt.Println(string(arquivo))
}
