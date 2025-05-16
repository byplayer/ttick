package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/byplayer/ttick/internal/cmd/ttick"
)




func main() {
    baseName := filepath.Base(os.Args[0])

    err := ttick.NewApp(baseName).Run(os.Args)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
