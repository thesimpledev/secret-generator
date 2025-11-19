package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <count> <length>\n", os.Args[0])
		os.Exit(1)
	}

	count, err := strconv.Atoi(os.Args[1])
	if err != nil || count <= 0 {
		fmt.Fprintln(os.Stderr, "Invalid count")
		os.Exit(1)
	}

	length, err := strconv.Atoi(os.Args[2])
	if err != nil || length <= 0 {
		fmt.Fprintln(os.Stderr, "Invalid length")
		os.Exit(1)
	}

	file, err := os.Create("secrets.md")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	for range count {

		buf := make([]byte, length)
		rand.Read(buf)
		key := base64.StdEncoding.EncodeToString(buf)
		key = key[:length]
		_, err = fmt.Fprintln(file, key)
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to write secret:", err)
			os.Exit(1)
		}

	}
}
