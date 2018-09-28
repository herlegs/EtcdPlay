package dto

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	etcdNodeIDRegex = "[0-9a-f]{15,16}"

	defaultClusterSize = 7

	ExampleFormat = "{time} {host} logs: {message}"
)

var (
	//default only 7 alias needed for cluster
	SeriousAliases = []string{"N1", "N2", "N3", "N4", "N5", "N6", "N7"}
	FunnyAliases   = []string{"Zezhou", "Weilun", "Troyzz", "Jimzzz", "Zhiyuan", "Karanz", "Yuguang"}
)

type KibanaRawLog struct {
	Hits *struct {
		Hits []*struct {
			Source *KibanaLogEntry `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type KibanaLogEntry struct {
	LogLevel     string `json:"log_level"`
	Hostname     string `json:"hostname"`
	TimestampStr string `json:"@timestamp"`
	Message      string `json:"message"`
	Time         time.Time
}

// SortableLogs is to define sorting logic for log entries
type SortableLogs []*KibanaLogEntry

func (logs SortableLogs) Len() int {
	return len(logs)
}

func (logs SortableLogs) Swap(i, j int) {
	logs[i], logs[j] = logs[j], logs[i]
}

func (logs SortableLogs) Less(i, j int) bool {
	return logs[i].Time.Before(logs[j].Time)
}

// CombinedLogs is combined logs of all hosts
type CombinedLogs []*KibanaLogEntry

func (logs CombinedLogs) getAllHosts() map[string]bool {
	hosts := map[string]bool{}
	for _, log := range logs {
		if !hosts[log.Hostname] {
			hosts[log.Hostname] = true
		}
	}
	return hosts
}

func (logs CombinedLogs) getAllEtcdNodeIDs() map[string]bool {
	nodeIDs := map[string]bool{}
	reg := regexp.MustCompile(etcdNodeIDRegex)
	for _, log := range logs {
		nodes := reg.FindAllString(log.Message, -1)
		for _, node := range nodes {
			if !nodeIDs[node] {
				nodeIDs[node] = true
			}
		}
	}
	return nodeIDs
}

//return hostname -> etcdnode mapping
func (logs CombinedLogs) GuessHostEtcdNodeMapping() (map[string]string, error) {
	mapping := map[string]string{}
	allEtcdNodes := logs.getAllEtcdNodeIDs()

	for _, log := range logs {
		for node := range allEtcdNodes {
			if strings.HasPrefix(log.Message, node) {
				mapping[log.Hostname] = node
			}
		}
	}

	hosts := logs.getAllHosts()
	if len(hosts) != defaultClusterSize || len(mapping) != defaultClusterSize {
		return nil, fmt.Errorf("error getting host and etcd nodes info. Hosts[%v] HostNodeMapping[%v]", len(hosts), len(mapping))
	}

	return mapping, nil
}

func (logs CombinedLogs) GroupByHost() map[string][]*KibanaLogEntry {
	maps := map[string][]*KibanaLogEntry{}
	for _, log := range logs {
		maps[log.Hostname] = append(maps[log.Hostname], log)
	}
	return maps
}

func (log *KibanaLogEntry) FormatString(format string) string {
	if format == "" {
		format = ExampleFormat
	}

	formatted := format
	formatted = strings.Replace(formatted, "{time}", log.TimestampStr, -1)
	formatted = strings.Replace(formatted, "{host}", log.Hostname, -1)
	formatted = strings.Replace(formatted, "{message}", log.Message, -1)

	return formatted
}
