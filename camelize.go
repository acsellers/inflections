package inflections

import (
	"strings"
)

func Camelize(str string) string {
	return strings.Replace(strings.ToTitle(str), " ", "", -1)
}
