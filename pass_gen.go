package main

import "fmt"

var perm_count uint64

func swap(per_text string, init, final int) string {
	swap_text := []rune(per_text)
	swap_text[init], swap_text[final] = swap_text[final], swap_text[init]
	per_text = string(swap_text)
	return per_text
}

func permute(per_text string, init, final int) {
	if init == final {
		fmt.Println(per_text)
		perm_count += 1
		return
	}
	for i := init; i <= final; i++ {
		per_text = swap(per_text, init, i)
		permute(per_text, init+1, final)
	}
}

func main() {
	var text string
	text = "enter the text to permute"
	final := len([]rune(text)) - 1
	permute(text, 0, final)
	fmt.Printf("total permutation count :: >> %d ", perm_count)
}
