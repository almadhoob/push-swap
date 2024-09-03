package swapper

import "fmt"

type Stack []int

func (s *Stack) Push(v int) {
	*s = append([]int{v}, *s...)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	v := (*s)[0]
	*s = (*s)[1:]
	return v, true
}

func (s *Stack) Swap() {
	if len(*s) < 2 {
		return
	}
	(*s)[0], (*s)[1] = (*s)[1], (*s)[0]
}

func (s *Stack) Rotate() {
	if len(*s) < 2 {
		return
	}
	v := (*s)[0]
	*s = append((*s)[1:], v)
}

func (s *Stack) ReverseRotate() {
	if len(*s) < 2 {
		return
	}
	v := (*s)[len(*s)-1]
	*s = append([]int{v}, (*s)[:len(*s)-1]...)
}

func IsSorted(s *Stack) bool {
	for i := 1; i < len(*s); i++ {
		if (*s)[i-1] > (*s)[i] {
			return false
		}
	}
	return true
}

func (s *Stack) Print() {
	for _, v := range *s {
		fmt.Println(v)
	}
	fmt.Println("=")
}
