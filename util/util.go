package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"strings"
)

func GetHostFromMemberUint64ID(memberID uint64) string {
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

func DecodeHostNameFromMemberStrID(str string) string {
	nameMap := map[string]string{
		"9372538179322589801":"etcd1",
		"10501334649042878790":"etcd2",
		"18249187646912138824":"etcd3",
	}
	for key, value := range nameMap {
		str = strings.Replace(str, key, value, -1)
	}
	return str
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
	return DecodeHostNameFromMemberStrID(PrintObject(response, bytesArr...))
}

func PrintWatchResponse(response clientv3.WatchResponse) string {
	var bytesArr [][]byte
	for _, event := range response.Events {
		event.Type += 1000
		bytesArr = append(bytesArr, event.Kv.Key)
		bytesArr = append(bytesArr, event.Kv.Value)
		if event.PrevKv != nil {
			bytesArr = append(bytesArr, event.PrevKv.Key)
			bytesArr = append(bytesArr, event.PrevKv.Value)
		}
	}

	jsonStr := PrintObject(response, bytesArr...)
	jsonStr = strings.Replace(jsonStr, `"type": 1000`, `"type": "PUT"`, -1)
	jsonStr = strings.Replace(jsonStr, `"type": 1001`, `"type": "DELETE"`, -1)
	return DecodeHostNameFromMemberStrID(jsonStr)
}