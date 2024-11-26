package updater

// import (
// 	_ "embed"
// 	"encoding/json"
// 	"fmt"
// )

// // read my version
// // read latest version
// // compare
// // if using old
// //     download new into temp folder
// //     replace .app before use
// //     open new .app
// // else
// //     open .app normally

// //go:embed wails.json
// var wailsJSON []byte

// func ParseWailsConfigToMap() (map[string]interface{}, error) {
// 	var config map[string]interface{}
// 	if err := json.Unmarshal(wailsJSON, &config); err != nil {
// 		return nil, fmt.Errorf("failed to parse wails.json to map: %w", err)
// 	}
// 	return config, nil
// }
