package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func status(msg string) {
	fmt.Println(msg)
}

func fatalError(msg string) {
	fmt.Fprintf(os.Stderr, "\033[31mError:\033[0m %s\n", msg)
	os.Exit(1)
}

func runCmd(cmd *exec.Cmd) {
	if verbose {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stdout = &bytes.Buffer{}
		cmd.Stderr = &bytes.Buffer{}
	}

	if err := cmd.Run(); err != nil {
		if !verbose {
			fmt.Println(cmd.Stderr)
		}
		os.Exit(1)
	}
}

func printBanner(tag string) {
	fmt.Println()
	fmt.Println("subcheck environment")
	fmt.Println("OS: Rocky Linux 9")
	fmt.Println("Compiler: g++ 11.3.0")
	fmt.Println("Image:", tag)
	fmt.Println()
}
