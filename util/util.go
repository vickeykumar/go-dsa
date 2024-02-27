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