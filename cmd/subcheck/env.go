package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func imageTag() string {
	h := sha256.Sum256([]byte(dockerfileCAEN))
	return fmt.Sprintf("%s:%x", baseImageName, h[:8])
}

func ensureEnv() {
	tag := imageTag()
	if imageExists(tag) {
		status("Using cached CAEN environment")
		printBanner(tag)
		return
	}

	spinner := NewSpinner("Preparing CAEN environment…")
	buildImage(tag)
	spinner.Stop("CAEN environment ready")

	printBanner(tag)
}

func imageExists(tag string) bool {
	return exec.Command("docker", "image", "inspect", tag).Run() == nil
}

func buildImage(tag string) {
	dir := mustTempDockerfile()
	defer os.RemoveAll(dir)

	cmd := exec.Command("docker", "build", "-t", tag, dir)
	runCmd(cmd)
}

func mustTempDockerfile() string {
	dir, err := os.MkdirTemp("", "subcheck-env-*")
	if err != nil {
		fatalError("Failed to create temp dir")
	}

	path := filepath.Join(dir, "Dockerfile")
	if err := os.WriteFile(path, []byte(dockerfileCAEN), 0644); err != nil {
		fatalError("Failed to write Dockerfile")
	}

	return dir
}
