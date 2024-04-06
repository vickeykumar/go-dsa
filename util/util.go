package util

func Max(i,j int) int {
	if i>j {
		return i
	}
	return j
}

func Min(i,j int) int {
	if i<j {
		return i 
	}
	return j
}

func ReverseString(str string) string {
	runes := []rune(str) 
	i:=0;
	j:=len(str)-1
	for i<j {
		runes[i],runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

func Pow(base, exponent int) int {
    result := 1
    for exponent > 0 {
        if exponent%2 == 1 {
            result *= base
        }
        base *= base
        exponent /= 2
    }
    return result
}
