package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flagSet()
}

// 1
func flagDefault() {
	textAddr := flag.String("text", "sem texto", "define um texto para ser printado")
	flag.Parse()

	fmt.Println(textAddr)
	fmt.Println(*textAddr)
}

// 2
type CLI struct {
	Text string
}

func flagWithStruct() {
	myCli := new(CLI)

	flag.StringVar(&myCli.Text, "text", "sem texto", "define um texto para ser printado")

	flag.Parse()

	fmt.Println(myCli)
	fmt.Println(myCli.Text)

}

//3
type App struct {
	Port string
	Name string
}

func flagSet() {
	myApp := new(App)

	if len(os.Args) >= 2 {
		fs := flag.NewFlagSet("app", flag.ExitOnError)

		fs.StringVar(&myApp.Name, "name", "default", "define nome do app")
		fs.StringVar(&myApp.Port, "port", "3001", "define port do app")

		fs.Parse(os.Args[2:])

		fmt.Printf("%s - Ouvindo na porta %s\n", myApp.Name, myApp.Port)
		return
	}

	fmt.Println("Informe no minimo um parametro")
}
