package commands

import (
	"flag"
	"fmt"
	"net/http"
)

type Start struct {
	Port    int
	Version string
	HelpF   bool
}

const helpTextStart = `Responsável por inicializar um servidor simples`
const helpLongTextStart = `
	Responsável por inicializar um servidor simples liberando acesso a todos
	os serviços configurados.

	expertscli start --port [serverPort] --version [serverVersion]
	- serverPot: int
	- serverVersion: string
`

const exampleTextStart = `
	expertscli start --port 3030 --version 1.0.0
	expertscli start --version 1.0.0         
`

func (cmd *Start) Name() string { return "start" }

func (cmd *Start) Example() string { return exampleTextStart }

func (cmd *Start) Help() string { return helpTextStart }

func (cmd *Start) LongHelp() string { return helpLongTextStart }

func (cmd *Start) Register(fs *flag.FlagSet) {
	fs.IntVar(&cmd.Port, "port", 8080, "porta do servidor")
	fs.StringVar(&cmd.Version, "version", "", "versão do servidor")
	fs.BoolVar(&cmd.HelpF, "help", false, "comando de help")
}

func (cmd *Start) Run() {
	if cmd.HelpF {
		fmt.Println(cmd.LongHelp())
		return
	}
	if cmd.Version == "" {
		fmt.Println("[--version] é obrigatório")
		return
	}

	fmt.Printf("Servidor v%s rodando na porta %v", cmd.Version, cmd.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", cmd.Port), nil)
}
