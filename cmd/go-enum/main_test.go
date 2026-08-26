package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain_TypeOverrideFlag(t *testing.T) {
	// Build the main binary first
	tmpDir, err := ioutil.TempDir("", "go-enum-main-test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	binaryPath := filepath.Join(tmpDir, "go-enum")
	buildCmd := exec.Command("go", "build", "-o", binaryPath, "main.go")
	buildCmd.Dir = "."
	err = buildCmd.Run()
	require.NoError(t, err)

	// Create a dummy Go file containing the enum
	fileContent := []byte(`package testpkg

//go:generate go-enum --type int16
const (
	StatusActive   Status = 1
	StatusInactive Status = 2
)
`)
	dummyFilePath := filepath.Join(tmpDir, "status.go")
	err = ioutil.WriteFile(dummyFilePath, fileContent, 0644)
	require.NoError(t, err)

	// Run the built go-enum binary with the --type flag override
	cmd := exec.Command(binaryPath, "--type", "int16", "--line", "4", "--with-value", dummyFilePath)
	cmd.Dir = tmpDir
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "failed to run go-enum: %s", string(output))

	// The generated file should be status.gen.go
	genFilePath := filepath.Join(tmpDir, "status.gen.go")
	genContent, err := ioutil.ReadFile(genFilePath)
	require.NoError(t, err)

	// Check if the generated file contains the overridden type
	assert.Contains(t, string(genContent), "type Status int16")
	assert.Contains(t, string(genContent), "func (v Status) Value() int16")
}
