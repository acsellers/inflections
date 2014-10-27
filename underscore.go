package inflections

import (
	"regexp"
	"strings"
)

var (
	doubleColon       = regexp.MustCompile("::")
	dash              = regexp.MustCompile("-")
	uppersOrNumsLower = regexp.MustCompile("([A-Z0-9]+)([A-Z][a-z])")
	lowerUpper        = regexp.MustCompile("([a-z])([A-Z0-9])")
)

// This function will change a string from a camelcased
// form to a string with underscores. Will change "::" to
// "/" to maintain compatibility with Rails's underscore
func Underscore(str string) string {
	output := doubleColon.ReplaceAllString(str, "/")

	// Rails uses underscores while I use dashes in this function
	// Go's regexp doesn't like $1_$2, so we'll use a dash instead
	// since it will get fixed in a later replacement
	output = uppersOrNumsLower.ReplaceAllString(output, "$1-$2")
	output = lowerUpper.ReplaceAllString(output, "$1-$2")

	output = strings.ToLower(output)
	output = dash.ReplaceAllString(output, "_")

	return output
}
