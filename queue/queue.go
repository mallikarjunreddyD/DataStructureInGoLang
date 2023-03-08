package queue

import "errors"

func Enqueue(a *[]int, item int) {
	*a = append(*a, item)
}
func Dequeue(a *[]int) (int, error) {
	if len(*a) == 0 {
		return 0, errors.New("queue is Empty")
	}
	item := (*a)[0]
	*a = append((*a)[1:len(*a)])
	return item, nil
}
