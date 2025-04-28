package main

import (
	"bufio"
	"os"
	//"bytes"
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

// send a NNTP command to the server
func (s *NNTPServer) SendCommand(cmd ...string) string {
	var fullCmd string

	if len(cmd) == 1 {
		fullCmd = cmd[0]
	} else if len(cmd) == 2 {
		fullCmd = fmt.Sprintf("%s %s\r\n", cmd[0], cmd[1])
	}

	log.Printf("sending command: %s", fullCmd)

	_, err := s.conn.Write([]byte(fullCmd))
	if err != nil {
		log.Fatalf("Error: '%v' sending command '%s'", fullCmd, err)
	}

	ret := readLine(s.rdr)
	log.Print(ret)

	return ret
}

// when server needs an authentication
func (s *NNTPServer) Authenticate() {
	// send userid
	resp := s.SendCommand("AUTHINFO USER", s.userid)

	// AUTHINFO USER should return "381 Password required" on success
	if !strings.HasPrefix(resp, "381") {
		log.Fatalf("AUTHINFO USER failed: %s", resp)
	}

	// send password
	resp = s.SendCommand("AUTHINFO PASS", s.password)

	// AUTHINFO PASS should return "281 Authentication accepted" on success
	if !strings.HasPrefix(resp, "281") {
		log.Fatalf("AUTHINFO PASS failed: %s", resp)
	}

}

// select a group before downloading
func (s *NNTPServer) SelectGroup(grp string) {
	resp := s.SendCommand("GROUP", grp)
	log.Println(resp)
}

// retrieve an article
func (s *NNTPServer) Download(id string) {
	// Request the article by Message-ID
	resp := s.SendCommand("ARTICLE", fmt.Sprintf("<%s>", id))
	//fmt.Println(resp)

	if !strings.HasPrefix(resp, "220") {
		log.Fatalf("Failed to retrieve article ID: %s", id)
	}

	err := os.WriteFile("output.bin", []byte(resp), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// // Read the article headers and body
	// var articleBody bytes.Buffer
	// for {
	// 	line, _ := s.rdr.ReadString('\n')
	// 	if strings.HasPrefix(line, ".") {
	// 		break // End of article
	// 	}
	// 	articleBody.WriteString(line)
	// }

	// // The body of the article might be encoded (e.g., yEnc or Base64)
	// // Here we assume Base64 encoding for simplicity. For yEnc, you'd need to decode it separately.
	// encodedBody := articleBody.String()
	// fmt.Println("Encoded Article Body:", encodedBody)

	// // Decode the Base64 encoded body
	// decodedBody, err := base64.StdEncoding.DecodeString(encodedBody)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Save the decoded binary data to a file (or process it further)
	// file, err := io.WriteFile("downloaded_file", decodedBody, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("File downloaded successfully!")

}

// // list articles from a group
// func (s *NNTPServer) ListArticlesFromGroup(group string) {
// 	// Select a group
// 	sendCommand(s.conn, fmt.Sprintf("LISTGROUP %s\r\n", group))
// 	resp := readLine(s.rdr)
// 	fmt.Println("LISTGROUP response:", resp)

// 	if !strings.HasPrefix(resp, "211") {
// 		panic("Failed to select group")
// 	}

// 	// Now read article numbers
// 	fmt.Println("=== Article Numbers ===")
// 	for {
// 		line := readLine(s.rdr)
// 		if line == "." {
// 			break
// 		}
// 		fmt.Println(line)
// 	}
// }

// // private function to just send NNTP command to server
// func sendCommand(conn net.Conn, cmd string) {
// 	_, err := conn.Write([]byte(cmd))
// 	if err != nil {
// 		log.Fatalf("Error: '%v' sending command '%s'", cmd, err)
// 	}
// }

// read a single line from I/O with server
func readLine(r *bufio.Reader) string {
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("error '%v' reading line from server", err)
	}
	return strings.TrimRight(line, "\r\n")
}
