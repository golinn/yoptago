package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/kvizyx/yoptago/internal/transpiler"
	"github.com/kvizyx/yoptago/internal/yoptalog"
)

const helpMessage = `YoptaGo CLI

Usage: yoptago [COMMAND] SRC_DIR

Commands:
 run - Run without output files
 build - Build project to .go files

`

const (
	defaultPath = "."

	yoptaSuffix = ".yo"
)

func main() {
	if len(os.Args) < 2 {
		yoptalog.Log("too few arguments. See 'yoptago help' for details.")
		os.Exit(1)
	}

	var (
		cmd  = os.Args[1]
		path = defaultPath
	)

	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	switch cmd {
	case "run":
		var mainPath string

		if len(os.Args) < 3 {
			yoptalog.Log("missing source directory")
			os.Exit(1)
		}

		mainPath = os.Args[2]

		if strings.HasSuffix(mainPath, yoptaSuffix) {
			mainPath = strings.TrimSuffix(mainPath, yoptaSuffix) + ".go"
		}

		tmpFiles, err := transpiler.TranspileDirectory(path)
		if err != nil {
			yoptalog.WithError("failed to transpile files", err)
			os.Exit(1)
		}
		defer clearTempFiles(tmpFiles)

		cmd := exec.Command("go", "run", mainPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			yoptalog.WithError("failed to run code", err)
			os.Exit(1)
		}

	case "build":
		_, err := transpiler.TranspileDirectory(path)
		if err != nil {
			yoptalog.WithError("failed to transpile files", err)
			os.Exit(1)
		}

	case "help":
		fmt.Print(helpMessage)

	default:
		yoptalog.Log("unknown command. See 'yoptago help' for details.")
		os.Exit(1)
	}
}

func clearTempFiles(paths []string) {
	for _, path := range paths {
		err := os.Remove(path)
		if err != nil {
			yoptalog.WithError("failed to remove tmp file", err)
		}
	}
}
