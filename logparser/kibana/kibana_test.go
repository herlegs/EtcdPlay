package kibana

import (
	"fmt"
	"testing"

	"github.com/herlegs/EtcdPlay/logparser/kibana/dto"

	"github.com/herlegs/EtcdPlay/util"
)

const (
	logFolderPrefix = "../logs/"
)

func TestParseKibanaLogEntriesFromFile(t *testing.T) {
	logs := ParseKibanaLogEntriesFromFile(logFolderPrefix + "kibana_9_28")
	fmt.Printf("%v\n", util.PrintObject(logs))
}

func TestJob_9_28(t *testing.T) {
	outFolder := logFolderPrefix + "2018_9_28/"
	logs := ParseKibanaLogEntriesFromFile(logFolderPrefix + "kibana_9_28")
	aliasMap := GetAliasMap(logs, dto.SeriousAliases)

	SaveLogsToFile(outFolder+"all", logs, aliasMap)

	hostLogs := SeparateLogsByHost(logs)
	for _, hostLog := range hostLogs {
		hostName := hostLog[0].Hostname
		SaveLogsToFile(outFolder+hostName, hostLog, aliasMap)
	}
}
