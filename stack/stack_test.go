package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	a := []int{}
	Push(&a, 10)
	Push(&a, 11)
	Push(&a, 12)
	if a[0] != 10 || a[1] != 11 || a[2] != 12 {
		t.Error("Expecting a[0]=10, a[1]=11 and a[2]=12")
	}
}
func TestPop(t *testing.T) {
	a := []int{}
	Push(&a, 10)
	Push(&a, 11)
	Push(&a, 12)
	item, err := Pop(&a)
	if err != nil || item != 12 {
		t.Error("Expecting error as nil and item as 12")
	}
	item, err = Pop(&a)
	if err != nil || item != 11 {
		t.Error("Expecting error as nil and item as 11")
	}
	item, err = Pop(&a)
	if err != nil || item != 10 {
		t.Error("Expecting error as nil and item as 10")
	}
	_, err = Pop(&a)
	if err == nil {
		t.Error("Expecting error")
	}
}
