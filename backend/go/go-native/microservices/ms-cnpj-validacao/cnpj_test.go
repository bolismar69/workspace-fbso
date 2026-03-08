package main

import "testing"

func TestIsValidCNPJ_KnownValid(t *testing.T) {
	v1 := "04.252.011/0001-10"
	v2 := "40.688.134/0001-61"
	if !isValidCNPJ(&v1) {
		t.Fatalf("expected valid for %q", v1)
	}
	if !isValidCNPJ(&v2) {
		t.Fatalf("expected valid for %q", v2)
	}
}

func TestIsValidCNPJ_Invalid(t *testing.T) {
	cases := []string{
		"04.252.011/0001-11",
		"00000000000000",
		"11111111111111",
		"",
	}
	for _, c := range cases {
		c := c
		if isValidCNPJ(&c) {
			t.Fatalf("expected invalid for %q", c)
		}
	}
	if isValidCNPJ(nil) {
		t.Fatalf("expected invalid for nil")
	}
}

func TestNormalizeCNPJ(t *testing.T) {
	in1 := "04.252.011/0001-10"
	n1 := normalizeCNPJ(&in1)
	if n1 == nil || *n1 != "04252011000110" {
		t.Fatalf("expected normalized digits, got %v", n1)
	}

	in2 := "  1-2-3  "
	n2 := normalizeCNPJ(&in2)
	if n2 == nil || *n2 != "123" {
		t.Fatalf("expected normalized digits, got %v", n2)
	}

	empty := ""
	if normalizeCNPJ(&empty) != nil {
		t.Fatalf("expected nil for empty string")
	}

	spaces := "   "
	if normalizeCNPJ(&spaces) != nil {
		t.Fatalf("expected nil for blank string")
	}

	if normalizeCNPJ(nil) != nil {
		t.Fatalf("expected nil for nil input")
	}
}
