package main

import (
	"who-ate-my-space/cli"
	"who-ate-my-space/cmd"
	"who-ate-my-space/presenters"
	"who-ate-my-space/utils"
)

func main() {

	args := cli.ParseArgs()

	var fileMap = make(map[string]int64)

	cmd.ScanDir(*args, &fileMap)

	topFiles := utils.GetHeaviestFiles(fileMap, args.Limit)

	presenters.PresentAsJson(topFiles)
}
