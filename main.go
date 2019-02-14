package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)
	flag.Parse()
}

func main() {
	dec := json.NewDecoder(os.Stdin)
	for {
		var msg interface{}
		if err := dec.Decode(&msg); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read message: %v", err)
		}
		if err := writeMessage(msg, os.Stdout); err != nil {
			log.Printf("write message: %v", err)
		}
	}
}

func writeMessage(msg interface{}, w io.Writer) error {
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err := writeMessageLength(jsonBytes, w); err != nil {
		return err
	}

	var buf bytes.Buffer
	n, err := buf.Write(jsonBytes)
	if err != nil {
		return err
	}
	if n != len(jsonBytes) {
		return fmt.Errorf("message not fully written")
	}

	nWritten, err := buf.WriteTo(w)
	if nWritten != int64(len(jsonBytes)) {
		return fmt.Errorf("message not fully written")
	}
	return err
}

func writeMessageLength(msg []byte, w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, uint32(len(msg)))
}
