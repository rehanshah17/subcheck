package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	imageName = "subcheck:280"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "build":
		runBuild()
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println("subcheck build")
	fmt.Println()
	fmt.Println("Runs `make release` in a CAEN-like environment to verify that")
	fmt.Println("your project will compile before submission, saving you a submission :D")
}

func runBuild() {
	checkDocker()
	buildImage()
	runMakeRelease()
}

func checkDocker() {
	cmd := exec.Command("docker", "--version")
	if err := cmd.Run(); err != nil {
		fatal("Docker is not installed or not available in PATH.\n")
	}
}

func buildImage() {
	fmt.Println("Building subcheck environment...")

	exePath, err := os.Executable()
	if err != nil {
		fatal("Failed to determine executable path.")
	}

	exeDir := filepath.Dir(exePath)
	dockerfilePath := filepath.Join(exeDir, "docker", "Dockerfile")

	cmd := exec.Command(
		"docker", "build",
		"--platform=linux/amd64",
		"-t", imageName,
		"-f", dockerfilePath,
		exeDir,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Capture output instead of streaming it
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println()
		fmt.Println("Docker image build failed.")
		fmt.Println("This is an internal subcheck error.")
		fmt.Println()
		fmt.Println(string(output))
		os.Exit(1)
	}
}

func runMakeRelease() {
	fmt.Println("Running `make release` in subcheck environment...")

	wd, err := os.Getwd()
	if err != nil {
		fatal("Failed to get working directory.")
	}

	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/work", wd),
		"-w", "/work",
		imageName,
		"make", "release",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println()
		fmt.Println("Build failed.")
		fmt.Println("This project would fail to compile on submission.")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("Build succeeded.")
	fmt.Println("This project will compile on submission.")
}

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, "Error:", msg)
	os.Exit(1)
}
