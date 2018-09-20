package util

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	nameMap := map[uint64]string{
		0x8211f1d0f64f3269: "etcd1",
		0x91bc3c398fb3c146: "etcd2",
		0xfd422379fda50e48: "etcd3",
	}

	for key, value := range nameMap {
		fmt.Printf("%v:%v\n",key,value)
	}
}

func TestClusterID(t *testing.T) {
	id := uint64(17237436991929493444)
	fmt.Printf("%x\n",id)
}
