package inflections

import (
	"testing"
)

func TestDasherize(t *testing.T) {
	if Dasherize("puni_puni") != "puni-puni" {
		t.Error("Could not process 'puni_puni', found",
			Dasherize("puni_puni"))
	}
}

func TestDemodularize(t *testing.T) {
	if Demodularize("Errors") != "Errors" {
		t.Error("Could not process 'Errors', found",
			Demodularize("Errors"))
	}
	if Demodularize("ActiveSupport::Unmaintainable") != "Unmaintainable" {
		t.Error("Could not process 'ActiveSupport::Unmaintainable', found",
			Demodularize("ActiveSupport::Unmaintainable"))
	}
}
