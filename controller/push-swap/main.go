package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func pa(a, b *Stack) {
	if v, ok := b.Pop(); ok {
		a.Push(v)
		fmt.Println("pa")
	}
}

func pb(a, b *Stack) {
	if v, ok := a.Pop(); ok {
		b.Push(v)
		fmt.Println("pb")
	}
}

func sa(a *Stack) {
	a.Swap()
	fmt.Println("sa")
}

func ra(a *Stack) {
	a.Rotate()
	fmt.Println("ra")
}

func rra(a *Stack) {
	a.ReverseRotate()
	fmt.Println("rra")
}

func sortStack(a, b *Stack) {
	if len(*a) <= 3 {
		sortThree(a)
		return
	}

	mid := (*a)[len(*a)/2]
	for len(*a) > 3 {
		if (*a)[0] < mid {
			pb(a, b)
		} else {
			ra(a)
		}
	}

	sortThree(a)

	for len(*b) > 0 {
		pa(a, b)
		if (*a)[0] > (*a)[1] {
			sa(a)
		}
	}
}

func sortThree(a *Stack) {
	if len(*a) == 2 {
		if (*a)[0] > (*a)[1] {
			sa(a)
		}
		return
	}
	if (*a)[0] > (*a)[1] && (*a)[1] > (*a)[2] {
		sa(a)
		rra(a)
	} else if (*a)[0] > (*a)[1] && (*a)[1] < (*a)[2] && (*a)[0] > (*a)[2] {
		ra(a)
	} else if (*a)[0] > (*a)[1] && (*a)[1] < (*a)[2] && (*a)[0] < (*a)[2] {
		sa(a)
	} else if (*a)[0] < (*a)[1] && (*a)[1] > (*a)[2] && (*a)[0] < (*a)[2] {
		sa(a)
		ra(a)
	} else if (*a)[0] < (*a)[1] && (*a)[1] > (*a)[2] && (*a)[0] > (*a)[2] {
		rra(a)
	}
}

func parseArgs(args []string) ([]int, error) {
	if len(args) == 0 {
		return nil, nil
	}

	// Join all arguments and split by spaces
	allArgs := strings.Join(args, " ")
	strNumbers := strings.Fields(allArgs)

	var numbers []int
	seen := make(map[int]bool)

	for _, strNum := range strNumbers {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return nil, fmt.Errorf("Error")
		}
		if seen[num] {
			return nil, fmt.Errorf("Error")
		}
		seen[num] = true
		numbers = append(numbers, num)
	}

	return numbers, nil
}

func main() {
	numbers, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if len(numbers) == 0 {
		return
	}

	a := Stack(numbers)
	var b Stack

	sortStack(&a, &b)
}
