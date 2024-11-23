package presenters

import (
	"encoding/json"
	"fmt"
	"os"
	"who-ate-my-space/utils"
)

func PresentAsJson(dirInfos []utils.DirInfo) {
	formattedDirData := make([]map[string]interface{}, 0)
	for _, dirInfo := range dirInfos {
		jsonDirInfo := map[string]interface{}{
			"absolute_path": dirInfo.AbsolutePath,
			"size":          dirInfo.Size,
		}
		formattedDirData = append(formattedDirData, jsonDirInfo)
	}

	jsonOutput, err := json.MarshalIndent(formattedDirData, "", "  ")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	os.Stdout.Write(jsonOutput)
}
