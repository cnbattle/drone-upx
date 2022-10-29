// Package cmd cmd
package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Cmd exec.Command
func Cmd(name string, arg ...string) error {
	command := exec.Command(name, arg...)
	var buffer bytes.Buffer
	command.Stdout = &buffer
	if err := command.Start(); err != nil {
		return err
	}
	if err := command.Wait(); err != nil {
		return err
	}
	fmt.Printf("%s", buffer.Bytes())
	fmt.Println("===================================================")
	fmt.Println("                Successfully                       ")
	fmt.Println("===================================================")
	return nil
}
