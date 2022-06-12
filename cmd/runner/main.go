package main

import (
	"plugin"
	"os"
	"fmt"
	"github.com/musobarlab/go-plugin-example/internal/dumper"
)

func main() {
	dumperType := "console"
	if len(os.Args) == 2 {
		dumperType = os.Args[1]
	}

	var plugSo string

	switch dumperType {
	case "console":
		plugSo = "./plugins/consoledumper/consoledumper.so"
		break
	case "file":
		plugSo = "./plugins/filedumper/filedumper.so"
		break
	default:
		fmt.Println("invalid plugin type")
		os.Exit(1)
		
	}

	plug, err := plugin.Open(plugSo)
	if err != nil {
		fmt.Printf("error loading plugin %v\n", err)
		os.Exit(1)
	}

	symDumper, err := plug.Lookup("Dumper")
	if err != nil {
		fmt.Printf("error loading symbol %v\n", err)
		os.Exit(1)
	}

	d, ok := symDumper.(dumper.Dumper)
	if !ok {
		fmt.Println("invalid Dumper type")
		os.Exit(1)
	}

	err = d.Dump("wuriyanto")
	if err != nil {
		fmt.Printf("error loading symbol %v\n", err)
		os.Exit(1)
	}

}