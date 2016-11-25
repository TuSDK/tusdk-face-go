package api

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
	// "fmt"
)

//参数签名
func sign(params map[string]string, key string) string {
	//排序
	keys := []string{}
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//拼接key value, 转换key为小写
	signstr := ""
	for _, k := range keys {
		signstr += strings.ToLower(k) + params[k]
	}
	//拼接私有key
	signstr += key

	//返回md5字符串
	hash := md5.New()
	hash.Write([]byte(signstr))
	return hex.EncodeToString(hash.Sum(nil))
}
