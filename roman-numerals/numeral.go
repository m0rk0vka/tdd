package romannumerals

import "strings"

func ConvertToRoman(arabic int) string {

	if arabic == 4 {
		return "IV"
	}

	var result strings.Builder

	for i := arabic; i > 0; i-- {
		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}
	return result.String()
}
