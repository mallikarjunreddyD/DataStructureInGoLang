package stack

import (
	"errors"
)

func Push(a *[]int, item int) {
	*a = append(*a, item)
}
func Pop(a *[]int) (int, error) {
	if len(*a) == 0 {
		return 0, errors.New("Underflow")
	}
	item := (*a)[len(*a)-1]
	*a = append((*a)[:len(*a)-1])
	return item, nil
}
