package main

import (
	"errors"
	"fmt"
	"github.com/zzhaolei/go-programming-tour-book/tour/cmd"
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
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v, %s", err, err)
	}
}
