package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

var decode bool

func init() {
	flag.BoolVar(&decode, "d", false, "Decode mode")
	flag.Parse()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if decode {
			decodedText, err := base64.StdEncoding.DecodeString(scanner.Text())
			if err != nil {
				fmt.Fprintf(os.Stderr, "can't decode %s due to %v\n", scanner.Text(), err)
				os.Exit(1)
			}
			fmt.Fprintln(os.Stdout, string(decodedText))
		} else {
			encodedText := base64.StdEncoding.EncodeToString([]byte(scanner.Text()))
			fmt.Fprintln(os.Stdout, encodedText)
		}
	}
}
