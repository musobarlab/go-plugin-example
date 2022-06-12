package main

import (
	"os"
)

func main() {}

type console int

func (c console) Dump(data string) error {
	_, err := os.Stdout.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}

var Dumper console