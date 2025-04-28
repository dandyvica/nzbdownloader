package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//───────────────────────────────────────────────────────────────────────────────────
	// get cli options
	//───────────────────────────────────────────────────────────────────────────────────
	opts := CliArgs()
	defer opts.logFile.Close()

	log.Printf("options: %+v", opts)

	//───────────────────────────────────────────────────────────────────────────────────
	// connect to server and authenticate
	//───────────────────────────────────────────────────────────────────────────────────
	srvName := fmt.Sprintf("%s:%d", opts.settings.Server.Name, opts.settings.Server.Port)
	log.Printf("connecting to server %s\n", srvName)

	srv := NewNNTPServer(srvName, opts.settings.Server.Userid, opts.settings.Server.Password, opts.settings.Server.Ssl)
	defer srv.conn.Close()

	srv.Authenticate()

	// don't go further if just checking the server
	if opts.check {
		fmt.Printf("connexion to server %s OK\n", srvName)
		os.Exit(1)
	}

	//───────────────────────────────────────────────────────────────────────────────────
	// interactive: useful to understand protocol
	//───────────────────────────────────────────────────────────────────────────────────
	if opts.interActive {
		// loop to ask for user input until "exit" is entered
		for {
			reader := bufio.NewReader(os.Stdin)

			// prompt the user for input
			fmt.Print("Enter NNTP command: ")

			// read the entire line of input, including spaces
			command, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			// Remove the trailing newline character (if any)
			command = command[:len(command)-1] + "\r\n"

			// check if the user wants to exit the loop
			if command == "exit" {
				fmt.Println("Exiting the program.")
				os.Exit(0)
			}

			// send command to server
			resp := srv.SendCommand(command)
			fmt.Println(resp)
		}
	}

	srv.SelectGroup("alt.binaries.test")
	srv.Download("pan$769c1$4988e27$83431a3c$92767679$1@none.com")
}
