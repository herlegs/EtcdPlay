package util

import "fmt"

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