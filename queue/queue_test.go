package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	a := []int{}
	Enqueue(&a, 10)
	Enqueue(&a, 11)
	Enqueue(&a, 12)
	if a[0] != 10 || a[1] != 11 || a[2] != 12 {
		t.Error("Expecting a[0]=10, a[1]=11 and a[2]=12")
	}
}
func TestDequeue(t *testing.T) {
	a := []int{}
	Enqueue(&a, 10)
	Enqueue(&a, 11)
	Enqueue(&a, 12)
	item, err := Dequeue(&a)
	if err != nil || item != 10 {
		t.Error("Expecting error as nil and item as 10")
	}
	item, err = Dequeue(&a)
	if err != nil || item != 11 {
		t.Error("Expecting error as nil and item as 11")

	}
	item, err = Dequeue(&a)
	if err != nil || item != 12 {
		t.Error("Expecting error as nil and item as 12")
	}
	_, err = Dequeue(&a)
	if err == nil {
		t.Error("Expecting error")
	}
}
