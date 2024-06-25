package leetcode

import (
	"strconv"
	"testing"
)

func TestBSTree0(t *testing.T) {
	bst := NewBSTree()
	bst.Set(5, "5")
	bst.Set(3, "3")
	bst.Set(2, "2")
	bst.Set(1, "1")
	bst.Set(4, "4")
	bst.Set(8, "8")
	bst.Set(9, "9")
	bst.Set(6, "6")
	bst.Set(7, "7")
	bst.Set(0, "0")
	if bst.Size() != 10 {
		t.Error()
	}
	for i := range 10 {
		if bst.Find(i) != strconv.Itoa(i) {
			t.Errorf("bst.Find(%v) is not '%v'!", i, i)
		}
		bst.Remove(i)
		if bst.Find(i) != nil {
			t.Errorf("%v was removed!", i)
		}
	}
	if bst.Size() != 0 {
		t.Error()
	}
}

func TestBSTree1(t *testing.T) {
	bst := NewBSTree()
	res := bst.Find(12)
	if res != nil {
		t.Errorf("Case 1 failed.")
	}
	bst.Set(12, "12")
	res = bst.Find(12)
	if res != "12" {
		t.Errorf("Case 2 failed.")
	}
	res = bst.Find(13)
	if res != nil {
		t.Errorf("Case 3 failed.")
	}
	bst.Set(13, "13")
	bst.Set(14, "14")
	bst.Set(10, "10")
	bst.Set(8, "8")
	bst.Set(9, "9")
	bst.Set(7, "7")
	err := bst.Remove(8)
	if err != nil {
		t.Errorf("Case 4 failed.")
	}
	if bst.Find(8) != nil {
		t.Errorf("Case 4.1 failed.")
	}

	bst.Set(5, "5")
	bst.Set(2, "2")
	bst.Set(4, "4")
	bst.Set(4, "40")
	bst.Set(4, "41")
	res = bst.Find(4)
	if res != "41" {
		t.Errorf("Case 5 failed.")
	}
}
