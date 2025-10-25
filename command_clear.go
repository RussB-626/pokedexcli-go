package main

import (
	"os"
	"os/exec"
)

func commandClear(config *Config) error {
	cmd := exec.Command("clear") // Or "printf", "tput clear"
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}
