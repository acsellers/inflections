package inflections

import (
	"testing"
)

func TestUnderscore(t *testing.T) {
	if Underscore("ActiveSupport::Errors") != "active_support/errors" {
		t.Error("Could not process 'ActiveSupport::Errors', found",
			Underscore("ActiveSupport::Errors"))
	}

	if Underscore("SSLError") != "ssl_error" {
		t.Error("Could not process 'SSLError', found",
			Underscore("SSLError"))
	}
}
