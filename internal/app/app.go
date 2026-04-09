package app

import (
	"os"
)

func Main(args []string, stdout, stderr *os.File) int {
	return Run(args, stdout, stderr)
}
