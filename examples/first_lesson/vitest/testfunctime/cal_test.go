package main

import (
	"testing"
)

// ＴestXcc　命名　后面参数固定不变
func TestCal(t *testing.T) {
	res := testslice()
	if len(res) == 0 {
		t.Fatal("no result", res)
	}

}

func TestUnmarshalSlice(t *testing.T) {
	res2, err := unmarshalslice()
	if err != nil {
		t.Fatal("unmarshal failed:", err)
	}
	if len(res2) == 0 {
		t.Fatal("no result", res2)
	}

}
