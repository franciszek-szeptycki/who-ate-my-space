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
	path := flag.String("path", "", "Directory to scan for files")
	ignoreSize := flag.Int64("ignore-size", 1024*1024, "Ignore files smaller than this size in bytes")
	limit := flag.Int("limit", 10, "Number of files to display")
	flag.Parse()

	if *path == "" {
		fmt.Println("The \"-path\" flag is required")
		os.Exit(1)
	}

	if *ignoreSize <= 0 {
		fmt.Println("The \"-ignore-size\" flag must be greater than 0")
		os.Exit(1)
	}

	if *limit <= 0 {
		fmt.Println("The \"-limit\" flag must be greater than 0")
		os.Exit(1)
	}

	return &Args{
		Path:       *path,
		IgnoreSize: *ignoreSize,
		Limit:      *limit,
	}
}
