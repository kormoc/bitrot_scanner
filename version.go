package main

import (
	"fmt"
	"os"
)

var Version = "__BITROT_SCANNER_VERSION__"

func versionFlag() {
	if config.Version {
		fmt.Printf("Version: %v\n", Version)
		os.Exit(0)
	}
}
