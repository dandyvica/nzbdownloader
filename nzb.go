package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/net/html/charset"
)

// Structs matching NZB file structure
type NZB struct {
	XMLName xml.Name  `xml:"nzb"`
	Files   []NZBFile `xml:"file"`
}

type NZBFile struct {
	Poster   string       `xml:"poster,attr"`
	Date     string       `xml:"date,attr"`
	Subject  string       `xml:"subject,attr"`
	Groups   []string     `xml:"groups>group"`
	Segments []NZBSegment `xml:"segments>segment"`
}

type NZBSegment struct {
	Bytes  uint32 `xml:"bytes,attr"`
	Number int    `xml:"number,attr"`
	ID     string `xml:",chardata"` // The message ID inside the tag
	Offset uint32 // offset of the segment within the file
}

// Implement the String() method
// func (z *NZB) String() string {
// 	return fmt.Sprintf("bytes=%d numner=")
// }

func (z *NZBSegment) Download(s *NNTPServer) {
	// Request the article by Message-ID
	resp := s.SendCommand("BODY", z.ID)

	if !strings.HasPrefix(resp, "222") {
		log.Fatalf("Failed to retrieve article ID: %s", z.ID)
	}

	// Create a file to save binary data
	outFile, err := os.Create("segment_output.part")
	if err != nil {
		log.Fatalf("Unable to create file")
	}
	defer outFile.Close()

	// Start reading lines (headers first)
	for {
		line := readLine(s.rdr)
		if line == "" {
			break // end of headers
		}
		// optional: you can print or skip headers
	}

	// Now, body starts (this is the encoded binary part usually yEnc)
	for {
		line, err := s.rdr.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\r\n")

		if line == "." {
			break // End of article body
		}

		if strings.HasPrefix(line, "..") {
			line = line[1:] // RFC3977 dot-stuffing: unescape lines starting with ".."
		}

		outFile.WriteString(line + "\n")
	}

	fmt.Println("Segment saved to segment_output.part")

}

func NewNZB(nzbFile string) *NZB {
	// Load the NZB file
	nzbData, err := os.ReadFile(nzbFile)
	if err != nil {
		log.Println("Error reading NZB file:", err)
		return nil
	}

	// handle errors like: "xml: encoding "iso-8859-1" declared but Decoder.CharsetReader is nil"
	decoder := xml.NewDecoder(strings.NewReader(string(nzbData)))
	decoder.CharsetReader = charset.NewReaderLabel

	var nzb NZB
	// err = xml.Unmarshal(nzbData, &nzb)
	err = decoder.Decode(&nzb)
	if err != nil {
		log.Println("Error parsing NZB:", err)
		return nil
	}

	return &nzb
}

// Implement the String() method
func (z *NZB) String() string {
	s := ""

	// Example: print all files and their segments
	for _, file := range z.Files {
		s = fmt.Sprintln("Subject:", file.Subject)
		for _, segment := range file.Segments {
			s += fmt.Sprintf("- Segment %d (%d bytes): %s\n", segment.Number, segment.Bytes, segment.ID)
		}
	}

	return s
}

// NZB segments should be sorted by number but don't assume that
func (z *NZB) Sort() {
	for _, file := range z.Files {
		sort.Slice(file.Segments, func(i, j int) bool {
			return file.Segments[i].Number < file.Segments[j].Number
		})
	}
}

// Assign offsets to each segment to manage multi-threading
func (z *NZB) AssignOffset() {
	for i := range z.Files {
		offset := uint32(0)
		for j := range z.Files[i].Segments {
			z.Files[i].Segments[j].Offset = offset
			offset += z.Files[i].Segments[j].Bytes
		}
	}
}
