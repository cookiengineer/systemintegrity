package structs

func isNumber(chunk string) bool {

	result := true

	for c := 0; c < len(chunk); c++ {

		character := string(chunk[c])

		if character >= "0" && character <= "9" {
			continue
		} else if character == "+" || character == "-" || character == "." {
			continue
		} else if character == "E" || character == "e" {
			continue
		} else {
			result = false
			break
		}

	}

	return result

}
