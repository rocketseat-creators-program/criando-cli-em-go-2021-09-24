package message

import (
	"flag"
	"fmt"
)

type Message struct {
	Text  string
	Helpf bool
}

const helpTextStart = `Responsável por printar um texto simples`
const helpLongTextStart = `
Responsável por printar um texto simples

expertscli message --text [printText]
- printText: string
`

const exampleTextStart = `
	expertscli message --text ola_mundo        
`

func (cmd *Message) Name() string {
	return "message"
}

func (cmd *Message) Example() string {
	return exampleTextStart
}

func (cmd *Message) Help() string {
	return helpTextStart
}

func (cmd *Message) LongHelp() string {
	return helpLongTextStart
}

func (cmd *Message) Register(fs *flag.FlagSet) {
	fs.StringVar(&cmd.Text, "text", "", "texto printado")
	fs.BoolVar(&cmd.Helpf, "help", false, "comando de help")
}

func (cmd *Message) Run() {
	if cmd.Helpf {
		fmt.Println(cmd.LongHelp())
		return
	}

	if cmd.Text == "" {
		fmt.Println("[--text] é obrigatório")
		return
	}

	fmt.Printf("O texto enviado foi: %s", cmd.Text)

}
