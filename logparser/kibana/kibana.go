package kibana

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/herlegs/EtcdPlay/logparser/kibana/dto"
	"github.com/herlegs/EtcdPlay/logparser/logutil"
)

const (
	KibanaTimeFormat = "2006-01-02T15:04:05.000Z"
)

func ParseKibanaLogEntriesFromFile(file string) []*dto.KibanaLogEntry {
	rawLog := logutil.ReadFile(file)

	rawLogDto := &dto.KibanaRawLog{}
	err := json.Unmarshal([]byte(rawLog), rawLogDto)
	if err != nil {
		fmt.Printf("error unmarshal kibana raw log: %v\n", err)
		return nil
	}

	var entries []*dto.KibanaLogEntry
	if rawLogDto.Hits != nil {
		for _, hit := range rawLogDto.Hits.Hits {
			hit.Source.Time = logutil.ParseTime(hit.Source.TimestampStr, KibanaTimeFormat)
			entries = append(entries, hit.Source)
		}
	}

	sortLogs(entries)

	return entries
}

func SeparateLogsByHost(logs dto.CombinedLogs) map[string][]*dto.KibanaLogEntry {
	return logs.GroupByHost()
}

func GetAliasMap(logs dto.CombinedLogs, aliasSource []string) map[string]string {
	hostEtcdNodeMap, err := logs.GuessHostEtcdNodeMapping()
	if err != nil {
		fmt.Printf("error getting host etcd node map: %v\n", err)
		return nil
	}
	return getHostAndNodeAlias(hostEtcdNodeMap, aliasSource)
}

func SaveLogsToFile(filename string, logs []*dto.KibanaLogEntry, aliasMap map[string]string) {
	fileContent := ""
	for i, log := range logs {
		fileContent += replaceWithAlias(log.FormatString(""), aliasMap) + "\n"
		if i > 0 && log.TimestampStr != logs[i-1].TimestampStr {
			fileContent += "\n"
		}
	}
	logutil.WriteFile(filename, fileContent)
}

func sortLogs(logs dto.SortableLogs) {
	sort.Sort(logs)
}

func getHostAndNodeAlias(hostNodeMap map[string]string, aliases []string) map[string]string {
	aliasMap := map[string]string{}
	i := 0
	for host, node := range hostNodeMap {
		aliasMap[host], aliasMap[node] = aliases[i], aliases[i]
		i++
	}
	return aliasMap
}

func replaceWithAlias(str string, aliasMap map[string]string) string {
	for key, alias := range aliasMap {
		str = strings.Replace(str, key, alias, -1)
	}
	return str
}
