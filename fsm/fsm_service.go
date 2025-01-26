package fsm

import (
	"encoding/json"
	"io"
	"os"
)

// LoadConfigFromFile 从文件加载配置
func LoadConfigFromFile(filePath string) (*FsmConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config FsmConfig
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
