package inflections

import "testing"

func TestUninflectedPluralize(t *testing.T) {
	if Pluralize("ox") != "oxen" {
		t.Error("Could not pluralize 'ox', found:", Pluralize("ox"))
	}
}

func TestNormalPluralize(t *testing.T) {
	if Pluralize("cat") != "cats" {
		t.Error("Could not pluralize 'cat', found:", Pluralize("cat"))
	}

	if Pluralize("user") != "users" {
		t.Error("Could not pluralize 'user', found:", Pluralize("user"))
	}
}
