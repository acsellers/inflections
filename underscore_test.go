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
	if Underscore("Address1") != "address_1" {
		t.Error("Could not process 'Address1', found",
			Underscore("Address1"))
	}
}
