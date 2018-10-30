package main

import (
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
	if len(os.Args) > 1 {
		if decode {
			decodedText, err := base64.StdEncoding.DecodeString(os.Args[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "can't decode %s due to %v\n", os.Args[1], err)
				os.Exit(1)
			}
			fmt.Fprintln(os.Stdout, decodedText)
		} else {
			encodedText := base64.StdEncoding.EncodeToString([]byte(os.Args[1]))
			fmt.Fprintln(os.Stdout, encodedText)
		}
	} else {
		fmt.Println("No args where specified.")
		os.Exit(0)
	}
}
