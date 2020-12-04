# nodecom
node公共的部分


## 装载数据库表文件
所有的数据库表通过配置文件设置
```
/**
* @param 	file 	配置文件路径
*/
func LoadTableFromFile(file string)
```

## 获取数据表名
```
/**
* 获取表的key
*/
func TableName(key string) string
```