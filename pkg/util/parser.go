package util

import (
	"fmt"
	"strings"
)

func FormattedCurrency(value string, separator string) string {
	stringArray := strings.Split(strings.Replace(value, ".", ",", -1), separator)
	formatted := ""
	for index, str := range stringArray {

		if index == (len(stringArray) - 1) {
			formatted = fmt.Sprintf("%s.%s", formatted, str)
		} else {
			formatted = formatted + str
		}
	}

	return formatted
}
