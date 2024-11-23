package main

import (
	"who-ate-my-space/cli"
	"who-ate-my-space/cmd"
	"who-ate-my-space/presenters"
	"who-ate-my-space/utils"
)

func main() {

	args := cli.ParseArgs()

	var dirSizeMap = make(map[string]int64)

	cmd.ScanDir(args.Path, &dirSizeMap)

	topFiles := utils.GetHeaviestFiles(dirSizeMap, args.Limit)

	presenters.PresentAsJson(topFiles)
}
