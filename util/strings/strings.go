package strings

import "strconv"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPalindromo(s string) bool {
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

func MakePalindromo(i int) (result int) {
	a := strconv.Itoa(i)
	b := Reverse(a)
	result, _ = strconv.Atoi(a + b)
	return
}
