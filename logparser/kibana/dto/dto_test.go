package dto

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDto(t *testing.T) {
	msg := "7a373a7b75484a5 [logterm: 13, index: 10680329, vote: d9d6f3d5bf966f2] ignored MsgVote from b103b543bba249a2 [logterm: 13, index: 10680329] at term 13: lease is not expired (remaining ticks: 5)"

	reg := regexp.MustCompile(etcdNodeIDRegex)
	result := reg.FindAllString(msg, -1)
	fmt.Printf("%v\n", result)
}
