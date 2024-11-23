package presenters

import (
	"encoding/json"
	"fmt"
	"os"
	"who-ate-my-space/utils"
)

func roundToTwoDecimalPlaces(num float64) float64 {
	return float64(int(num*100)) / 100
}

func PresentAsJson(dirInfos []utils.ExtendedDirInfo) {
	formattedDirData := make([]map[string]interface{}, 0)
	for _, dirInfo := range dirInfos {
		jsonDirInfo := map[string]interface{}{
			"absolute_path":  dirInfo.AbsolutePath,
			"size_gigabytes": dirInfo.SizeGB,
			"size_megabytes": dirInfo.SizeMB,
			"size_kilobytes": dirInfo.SizeKB,
			"size_bytes":     dirInfo.SizeB,
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
