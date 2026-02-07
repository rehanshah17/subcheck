package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func testSpinner() {
	spinner := NewSpinner("Testing spinner…")
	time.Sleep(3 * time.Second)
	spinner.Stop("Spinner test complete")
	status("Spinner works!")
}

func runBuild() {
	ensureEnv()
	checkFootguns()

	status("Running make release…")
	runInContainer("make", "release")
	status("Build succeeded (CAEN-compatible)")
}

func runValgrind() {
	ensureEnv()
	checkFootguns()

	executable := getExecutableName()
	status("Running under valgrind…")
	runInContainer("valgrind", "--leak-check=full", "./"+executable)
}

func runInContainer(args ...string) {
	wd, _ := os.Getwd()

	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", wd+":/work",
		"-w", "/work",
		imageTag(),
	)
	cmd.Args = append(cmd.Args, args...)

	runCmd(cmd)
}

func checkFootguns() {
	data, err := os.ReadFile("Makefile")
	if err != nil {
		fatalError("No Makefile found")
	}

	if !strings.Contains(string(data), "release:") {
		fatalError(`Missing "release" target in Makefile.

CAEN autograder runs:
  make release`)
	}
}

func getExecutableName() string {
	data, err := os.ReadFile("Makefile")
	if err != nil {
		fatalError("No Makefile found")
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "EXECUTABLE") && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[1])
				return name
			}
		}
	}

	fatalError(`Missing "EXECUTABLE" variable in Makefile`)
	return ""
}

func runDoctor() {
	status("Running diagnostics…")

	check("Docker installed", exec.Command("docker", "--version"))
	check("Docker daemon running", exec.Command("docker", "info"))
	check("Makefile present", exec.Command("test", "-f", "Makefile"))

	status("Enviornment ready for subcheck")
}

func check(name string, cmd *exec.Cmd) {
	if cmd.Run() != nil {
		fmt.Printf("\033[31m✗ %s\033[0m\n", name)
		os.Exit(1)
	}
	fmt.Printf("\033[32m✓ %s\033[0m\n", name)
}
