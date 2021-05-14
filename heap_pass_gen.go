package main

import "fmt"

var perm_count, iter_count uint64

func swap(per_text string, init, final int) string {
	swap_text := []rune(per_text)
	swap_text[init], swap_text[final] = swap_text[final], swap_text[init]
	per_text = string(swap_text)
	return per_text
}

func permute(per_text string, final int) {
	iter_count += 1
	if final == 1 {
		fmt.Println(per_text)
		perm_count += 1
		return
	}
	for i := 0; i < final-1; i++ {
		permute(per_text, final-1)
		if final%2 == 0 {
			swap(per_text, i, final-1)
		} else {
			swap(per_text, 0, final-1)
		}

	}
}

func main() {
	var text string
	text = "ABCD"
	final := len([]rune(text))
	permute(text, final)
	fmt.Println("the total no of permuation is :>> ", perm_count)
	fmt.Println("the total no of iterations are : >> ", iter_count)
}
