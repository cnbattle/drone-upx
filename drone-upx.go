package main

import (
	"errors"
	"github.com/cnbattle/drone-upx/cmd"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	// Version Version
	Version string
)

func main() {
	app := &cli.App{
		Name:      "drone-upx",
		Usage:     "UPX - the Ultimate Packer for eXecutables.",
		Copyright: "Copyright (c) 2020 Qi-ai Li",
		Version:   Version,
		Authors: []*cli.Author{
			{
				Name:  "Qi-ai Li",
				Email: "qiaicn@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "level, l",
				Value:   "5",
				Usage:   "upx level",
				EnvVars: []string{"PLUGIN_LEVEL"},
			},

			&cli.StringFlag{
				Name:    "originalFile, of",
				Usage:   "original file",
				EnvVars: []string{"PLUGIN_ORIGINAL_FILE"},
			},

			&cli.StringFlag{
				Name:    "saveFile, sf",
				Usage:   "save file name",
				EnvVars: []string{"PLUGIN_SAVE_FILE"},
			},
		},
		Action: run,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx *cli.Context) error {
	var level, originalFile, saveFile = ctx.String("level"), ctx.String("originalFile"), ctx.String("saveFile")
	if len(originalFile) == 0 {
		return errors.New("miss originalFile")
	}
	if len(saveFile) == 0 {
		return errors.New("miss saveFile")
	}
	return cmd.Cmd("/bin/upx", "-"+level, "-o", saveFile, originalFile)
}
