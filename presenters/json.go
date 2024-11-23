package presenters

import (
	"encoding/json"
	"fmt"
	"who-ate-my-space/utils"
)

func PresentAsJson(fileInfos []utils.FileInfo) {
	jsonFileInfos := make([]map[string]interface{}, 0)
	for _, fileInfo := range fileInfos {
		jsonFileInfo := map[string]interface{}{
			"absolute_path": fileInfo.AbsolutePath,
			"size":          fileInfo.Size,
		}
		jsonFileInfos = append(jsonFileInfos, jsonFileInfo)
	}

	jsonBody := map[string]interface{}{
		"files": jsonFileInfos,
	}

	jsonOutput, err := json.Marshal(jsonBody)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(string(jsonOutput))
}
