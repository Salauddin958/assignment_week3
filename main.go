package main

import "strings"

func main() {
	// tested in hackerrank environment
}

//https://www.hackerrank.com/challenges/camelcase/problem
// solution
func camelcase(s string) int32 {
	if len(s) == 0 {
		return 0
	}
	ans := 0
	for _, value := range s {
		if value >= 65 && value <= 90 {
			ans++
		}
	}
	return int32(ans + 1)
}

//https://www.hackerrank.com/challenges/caesar-cipher-1/problem
//solution
func caesarCipher(s string, k, n int32) string {
	if k == 0 || n == 0 {
		return s
	}
	k = k % 26
	str := []byte(s)
	sl := []byte{}
	for _, v := range str {
		if v >= 65 && v <= 90 {
			if byte(int32(v)+k) > 90 {
				sl = append(sl, byte(int32(v)+k-26))
			} else {
				sl = append(sl, byte(int32(v)+k))
			}
		} else if v >= 97 && v <= 122 {
			if byte(int32(v)+k) > 122 {
				sl = append(sl, byte(int32(v)+k-26))
			} else {
				sl = append(sl, byte(int32(v)+k))
			}
		} else {
			sl = append(sl, v)
		}
	}
	return string(sl)
}

//https://www.hackerrank.com/challenges/pangrams/problem
//solution
func pangrams(s string) string {
	// Write your code here
	if len(s) == 0 {
		return s
	}
	res := strings.ToUpper(s)
	arr := [26]int{}
	for _, v := range []byte(res) {
		if v >= 65 && v <= 90 {
			arr[int(v)-65]++
		}
	}
	for _, v := range arr {
		if v == 0 {
			return "not pangram"
		}
	}
	return "pangram"
}

//https://www.hackerrank.com/challenges/string-construction/problem
// solution
func stringConstruction(s string) int32 {
	if len(s) == 0 {
		return 0
	}
	var ans int32 = 0
	arr := [26]int{}
	for _, v := range []byte(s) {
		if v >= 97 && v <= 122 {
			arr[int(v)-97]++
		}
	}
	for _, v := range arr {
		if v >= 1 {
			ans++
		}
	}
	return ans
}
