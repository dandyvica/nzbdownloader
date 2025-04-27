package main

import (
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
}
