package cli

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	Path       string
	IgnoreSize int64
	Limit      int
}

func ParseArgs() *Args {
	path := flag.String("path", "/var", "Directory to scan for files")
	ignoreSize := flag.Int64("ignore-size", 1024*1024, "Ignore files smaller than this size in bytes")
	limit := flag.Int("limit", 10, "Number of files to display")
	flag.Parse()

	if _, err := os.Stat(*path); os.IsNotExist(err) {
		fmt.Printf("The provided path does not exist: %s\n", *path)
		os.Exit(1)
	}

	return &Args{
		Path:       *path,
		IgnoreSize: *ignoreSize,
		Limit:      *limit,
	}
}
