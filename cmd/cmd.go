package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Cmd exec.Command
func Cmd(name string, arg ...string) {
	command := exec.Command(name, arg...)
	var buffer bytes.Buffer
	command.Stdout = &buffer
	if err := command.Start(); err != nil {
		panic(fmt.Sprintf("Error: The first command can not be startup %s\n", err))
	}
	if err := command.Wait(); err != nil {
		panic(fmt.Sprintf("Error: Couldn't wait for the second command: %s\n", err))
	}
	fmt.Printf("%s", buffer.Bytes())
}
