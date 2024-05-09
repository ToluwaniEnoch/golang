package iteration

func Repeat(char string, numberOfTimesToRepeat int) string {
	repeatedString := ""

	for i :=0; i < numberOfTimesToRepeat; i++ {
		repeatedString += char
	}

	return repeatedString
}
