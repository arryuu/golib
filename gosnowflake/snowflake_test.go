package gosnowflake

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateID(t *testing.T) {
	node, err := NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i, node.Generate())
		}(i)
	}

	time.Sleep(2e9)
}
