package main

import (
	"github.com/cnbattle/drone-upx/cmd"
	"os"
)

var (
	level        string
	saveFile     string
	originalFile string
)

func init() {
	level = os.Getenv("PLUGIN_LEVEL")
	saveFile = os.Getenv("PLUGIN_SAVE_FILE")
	originalFile = os.Getenv("PLUGIN_ORIGINAL_FILE")
}

func main() {
	cmd.Cmd("/bin/upx", "-"+level, "-o", saveFile, originalFile)
}
