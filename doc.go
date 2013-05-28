/*
Inflections is a set of functions that duplicate
the functions in the ActiveSupport::Inflector system
from Rails. Some functions are wrappers around go
stdlib functions, unless I realize the go stdlib
has differing behavior than the ActiveSupport func,
in which case the function here will try to mirror the
AS version.

Note: Some functions are not mappable to Go functions,
the current known functions are: constantize. Also the
strange optional arguments available in the Rails
functions are not available, if you actually need the
optional arguments, please tell me your use case.
*/
package inflections

import "regexp"

var (
	underscore = regexp.MustCompile("_")
)
