package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CuriousHet/D-CAS/cli"
	"github.com/CuriousHet/D-CAS/network"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: store <file-path> | get <hash> [-port=<port>]")
		os.Exit(1)
	}

	// Default port
	port := 3000

	// Parse command-line arguments
	args := os.Args[1:]
	for _, arg := range args {
		if len(arg) > 6 && arg[:6] == "-port=" {
			p, err := strconv.Atoi(arg[6:])
			if err != nil {
				fmt.Println("Invalid port number")
				os.Exit(1)
			}
			port = p
		}
	}

	// Command: Start a node
	if args[0] == "node" {
		tr := network.NewTCPTransport(fmt.Sprintf(":%d", port))
		fmt.Printf("Listening on :%d\n", port)
		if err := tr.ListenAndAccept(); err != nil {
			fmt.Println("Error starting node:", err)
			os.Exit(1)
		}
		select {}
	}

	// Command: Store a file
	if args[0] == "store" && len(args) > 1 {
		cli.StoreFile(args[1])
		return
	}

	// Command: Get a file
	if args[0] == "get" && len(args) > 1 {
		cli.RetrieveFile(args[1])
		return
	}

	// If no valid command, print usage
	fmt.Println("Usage: store <file-path> | get <hash> | node [-port=<port>]")
}
