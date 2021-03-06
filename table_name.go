package nodecom

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var TABLE_NAME map[string]interface{}

// 加载数据表名配置文件
//
// @param file
// @return err
//
func LoadTableFromFile(file string) (err error) {
	if nil != TABLE_NAME {
		return nil
	}

	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)
	err = json.Unmarshal([]byte(data), &TABLE_NAME)
	if nil != err {
		return
	}

	return
}

// 通过key获取表的真实名字
//
// @param key 	表的key（主要是所需要表名的全大写）
// @return string
//
func TableName(key string) string {
	if _, ok := TABLE_NAME[key]; !ok {
		return "NOT_TABLE"
	}

	return TABLE_NAME[key].(string)
}
