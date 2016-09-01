package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsHexWord(word string) bool {
	for _, c := range word {
		lower_c := strings.ToLower(string(c))
		if lower_c < "a" || lower_c > "f" {
			return false
		}
	}
	return true
}

func main() {
	f, _ := os.Open("/usr/share/dict/words")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		currentWord := scanner.Text()
		if len(currentWord) > 1 && IsHexWord(currentWord) {
			fmt.Println(currentWord)
		}
	}
	if err := scanner.Err(); err != nil{
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
