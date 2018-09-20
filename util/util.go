package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"strings"
)

func Parse(memberID uint64) string {
	hexID := fmt.Sprintf("%x", memberID)
	nameMap := map[string]string{
		"8211f1d0f64f3269": "etcd1",
		"91bc3c398fb3c146": "etcd2",
		"fd422379fda50e48": "etcd3",
	}
	if parsed, ok := nameMap[hexID]; ok {
		return parsed
	}
	return "unknown"
}


func PrintObject(o interface{}, bytesArr ...[]byte) string {
	b, err := json.MarshalIndent(o, "|", "\t")
	jsonStr := ""
	if err != nil {
		return fmt.Sprintf("[PrintObject Error:%v]", err)

	}
	jsonStr = string(b)

	for _, bytes := range bytesArr {
		b64Str := base64.StdEncoding.EncodeToString(bytes)
		jsonStr = strings.Replace(jsonStr, b64Str, string(bytes), -1)
	}

	return jsonStr
}

func PrintGetResponse(response *clientv3.GetResponse) string {
	var bytesArr [][]byte
	for _, kv := range response.Kvs {
		bytesArr = append(bytesArr, kv.Key)
		bytesArr = append(bytesArr, kv.Value)
	}
	return PrintObject(response, bytesArr...)
}