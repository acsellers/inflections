package inflections

import (
	"strings"
)

func CamelCase(str string) string {
	return Camelize(str)
}

func Dasherize(str string) string {
	return underscore.ReplaceAllString(str, "-")
}

func Demodularize(str string) string {
	items := strings.Split(str, "::")
	return items[len(items)-1]
}

func Titlecase(str string) string {
	return strings.ToTitle(str)
}

func Titleize(str string) string {
	return strings.ToTitle(str)
}
