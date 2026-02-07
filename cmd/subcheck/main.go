package main

import (
	"fmt"
	"os"
)

const baseImageName = "subcheck"

var verbose bool

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	if os.Args[1] == "--verbose" {
		verbose = true
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-spin":
		testSpinner()
	case "env":
		ensureEnv()
	case "build":
		runBuild()
	case "valgrind":
		runValgrind()
	case "doctor":
		runDoctor()
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  subcheck env")
	fmt.Println("  subcheck build")
	fmt.Println("  subcheck valgrind")
	fmt.Println("  subcheck doctor")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --verbose   Show Docker output")
	fmt.Println("  -spin       Test the spinner")
}
