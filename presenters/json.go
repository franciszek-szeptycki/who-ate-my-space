package presenters

import (
	"encoding/json"
	"fmt"
	"who-ate-my-space/utils"
)

func PresentAsJson(fileInfos []utils.FileInfo) {
	jsonData, err := json.Marshal(fileInfos)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(string(jsonData))
}
