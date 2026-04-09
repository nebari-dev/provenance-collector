package verify

import "testing"

func TestIsSLSAPredicate(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"https://slsa.dev/provenance/v1", true},
		{"https://slsa.dev/provenance/v0.2", true},
		{"https://in-toto.io/provenance/v1", true},
		{"https://spdx.dev/Document", false},
		{"https://cyclonedx.org/bom", false},
		{"", false},
		{"random string", false},
	}
	for _, tc := range tests {
		got := isSLSAPredicate(tc.input)
		if got != tc.want {
			t.Errorf("isSLSAPredicate(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestNewProvenanceChecker(t *testing.T) {
	c := NewProvenanceChecker()
	if c == nil {
		t.Fatal("expected non-nil checker")
	}
	if _, ok := c.(*SLSAProvenanceChecker); !ok {
		t.Error("expected *SLSAProvenanceChecker")
	}
}
