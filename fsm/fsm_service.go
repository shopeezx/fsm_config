/**
 * @Copyright: 2025 Shopee
 * @Author: xiang.zhong
 * @Email: xiang.zhong@shopee.com
 * @Date: 2025/01/24 15:41
 */
package fsm

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// LoadConfigFromFile 从文件加载配置
func LoadConfigFromFile(filePath string) (*FsmConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config FsmConfig
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
