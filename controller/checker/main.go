package main

import (
	"bufio"
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

func executeInstruction(a, b *Stack, instruction string) error {
	switch instruction {
	case "sa":
		a.Swap()
	case "sb":
		b.Swap()
	case "ss":
		a.Swap()
		b.Swap()
	case "pa":
		if v, ok := b.Pop(); ok {
			a.Push(v)
		}
	case "pb":
		if v, ok := a.Pop(); ok {
			b.Push(v)
		}
	case "ra":
		a.Rotate()
	case "rb":
		b.Rotate()
	case "rr":
		a.Rotate()
		b.Rotate()
	case "rra":
		a.ReverseRotate()
	case "rrb":
		b.ReverseRotate()
	case "rrr":
		a.ReverseRotate()
		b.ReverseRotate()
	default:
		return fmt.Errorf("Error")
	}
	return nil
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

func isSorted(s *Stack) bool {
	for i := 1; i < len(*s); i++ {
		if (*s)[i-1] > (*s)[i] {
			return false
		}
	}
	return true
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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := strings.TrimSpace(scanner.Text())
		if instruction == "" {
			continue
		}
		err := executeInstruction(&a, &b, instruction)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if isSorted(&a) && len(b) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
