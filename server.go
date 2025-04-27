package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strings"
)

// everything need to communicate to a server
type NNTPServer struct {
	server   string        // name of the usenet server
	userid   string        // when usenet server needs a userid
	password string        // when usenet server needs a password
	conn     net.Conn      // handle of the connexion to the server
	rdr      *bufio.Reader // buffer reader
}

// initiate connexion to the server
func NewNNTPServer(server string, userid string, password string, ssl bool) *NNTPServer {
	var conn net.Conn
	var err error

	// establish connexion to server
	if ssl {
		conn, err = tls.Dial("tcp", server, nil)
	} else {
		conn, err = net.Dial("tcp", server)
	}

	if err != nil {
		log.Fatalf("error '%v' connecting to server '%s'", err, server)
	}

	// get server greetings
	rdr := bufio.NewReader(conn)
	log.Printf("greeting from server: %s", readLine(rdr))

	return &NNTPServer{
		server, userid, password, conn, rdr,
	}
}

// when server needs an authentication
func (s *NNTPServer) Authenticate() {
	// send userid
	sendCommand(s.conn, fmt.Sprintf("AUTHINFO USER %s\r\n", s.userid))
	resp := readLine(s.rdr)
	if !strings.HasPrefix(resp, "381") {
		log.Fatalf("AUTHINFO USER failed: %s", resp)
	}

	// send password
	sendCommand(s.conn, fmt.Sprintf("AUTHINFO PASS %s\r\n", s.password))
	resp = readLine(s.rdr)
	if !strings.HasPrefix(resp, "281") {
		log.Fatalf("AUTHINFO PASS failed: %s", resp)
	}

}

// list articles from a group
func (s *NNTPServer) ListArticlesFromGroup(group string) {
	// Select a group
	sendCommand(s.conn, fmt.Sprintf("LISTGROUP %s\r\n", group))
	resp := readLine(s.rdr)
	fmt.Println("LISTGROUP response:", resp)

	if !strings.HasPrefix(resp, "211") {
		panic("Failed to select group")
	}

	// Now read article numbers
	fmt.Println("=== Article Numbers ===")
	for {
		line := readLine(s.rdr)
		if line == "." {
			break
		}
		fmt.Println(line)
	}
}

// private function to just send NNTP command to server
func sendCommand(conn net.Conn, cmd string) {
	_, err := conn.Write([]byte(cmd))
	if err != nil {
		log.Fatalf("Error: '%v' sending command '%s'", cmd, err)
	}
}

// read a single line from I/O with server
func readLine(r *bufio.Reader) string {
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("error '%v' reading line from server", err)
	}
	return strings.TrimRight(line, "\r\n")
}
