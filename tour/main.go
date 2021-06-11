package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Name string

func (i *Name) String()string{
	return fmt.Sprint(*i)
}

func (i *Name) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}

	*i = Name("name: " + value)
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	var name Name

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.Var(&name, "name", "Go语言")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.Var(&name, "n", "PHP语言")
		_ = phpCmd.Parse(args[1:])
	}

	log.Println(name)
}
