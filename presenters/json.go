package presenters

import (
	"encoding/json"
	"fmt"
	"who-ate-my-space/utils"
)

func PresentAsJson(dirInfos []utils.DirInfo) {
	jsonFormattedDirData := make([]map[string]interface{}, 0)
	for _, dirInfo := range dirInfos {
		jsonDirInfo := map[string]interface{}{
			"absolute_path": dirInfo.AbsolutePath,
			"size":          dirInfo.Size,
		}
		jsonFormattedDirData = append(jsonFormattedDirData, jsonDirInfo)
	}

	jsonOutput, err := json.Marshal(jsonFormattedDirData)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(string(jsonOutput))
}
