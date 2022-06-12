package main

import (
	"os"
)

func main() {}

type file int

func (c file) Dump(data string) error {
	f, err := os.Create("out.txt")
	if err != nil {
		return err
	}

	defer func() { f.Close() }()

	_, err = f.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

var Dumper file