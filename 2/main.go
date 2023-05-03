package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func stringUnpack(s string) (string, error) {
	var newString strings.Builder
	if len(s) == 0 {
		return "", nil
	}
	runeArray := []rune(s)
	for i := 0; i < len(runeArray)-1; i++ {
		if unicode.IsDigit(runeArray[i]) {
			if i == 0 {
				return "", errors.New("Unexpected digit")
			}
			num, err := strconv.Atoi(string(runeArray[i]))
			if err != nil {
				return "", errors.New("Atoi error")
			}
			for j := 1; j < num; j++ {
				newString.WriteRune(runeArray[i-1])
			}
		} else {
			newString.WriteRune(runeArray[i])
		}
	}
	if !unicode.IsDigit(runeArray[len(runeArray)-1]) {
		newString.WriteRune(runeArray[len(runeArray)-1])
	}
	return newString.String(), nil
}
func main() {
	//s, _ := stringUnpack("a4bc2d5e")
	s, _ := stringUnpack("")
	fmt.Println(s)

}
