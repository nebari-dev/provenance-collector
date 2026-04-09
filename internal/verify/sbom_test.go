package verify

import (
	"testing"
)

func TestDetectSBOMFormat_InTotoSPDX(t *testing.T) {
	payload := `{
		"_type": "https://in-toto.io/Statement/v0.1",
		"predicateType": "https://spdx.dev/Document",
		"predicate": {"spdxVersion": "SPDX-2.3"}
	}`
	got := detectSBOMFormat([]byte(payload))
	if got != "spdx" {
		t.Errorf("expected spdx, got %q", got)
	}
}

func TestDetectSBOMFormat_InTotoCycloneDX(t *testing.T) {
	payload := `{
		"_type": "https://in-toto.io/Statement/v0.1",
		"predicateType": "https://cyclonedx.org/bom/v1.4",
		"predicate": {"bomFormat": "CycloneDX"}
	}`
	got := detectSBOMFormat([]byte(payload))
	if got != "cyclonedx" {
		t.Errorf("expected cyclonedx, got %q", got)
	}
}

func TestDetectSBOMFormat_InTotoUnknownPredicateWithSPDXContent(t *testing.T) {
	payload := `{
		"_type": "https://in-toto.io/Statement/v0.1",
		"predicateType": "https://example.com/custom",
		"predicate": {"spdxVersion": "SPDX-2.3", "SPDXID": "SPDXRef-DOCUMENT"}
	}`
	got := detectSBOMFormat([]byte(payload))
	if got != "spdx" {
		t.Errorf("expected spdx from predicate content, got %q", got)
	}
}

func TestDetectSBOMFormat_RawSPDXContent(t *testing.T) {
	payload := `{"spdxVersion": "SPDX-2.3", "SPDXID": "SPDXRef-DOCUMENT"}`
	got := detectSBOMFormat([]byte(payload))
	if got != "spdx" {
		t.Errorf("expected spdx from raw content, got %q", got)
	}
}

func TestDetectSBOMFormat_RawCycloneDXContent(t *testing.T) {
	payload := `{"bomFormat": "CycloneDX", "specVersion": "1.4"}`
	got := detectSBOMFormat([]byte(payload))
	if got != "cyclonedx" {
		t.Errorf("expected cyclonedx from raw content, got %q", got)
	}
}

func TestDetectSBOMFormat_Empty(t *testing.T) {
	got := detectSBOMFormat([]byte("{}"))
	if got != "" {
		t.Errorf("expected empty, got %q", got)
	}
}

func TestDetectSBOMFormat_InvalidJSON(t *testing.T) {
	// Falls back to content-based detection
	got := detectSBOMFormat([]byte("not json but has SPDXRef- in it"))
	if got != "spdx" {
		t.Errorf("expected spdx from fallback, got %q", got)
	}
}

func TestDetectSBOMFormat_NoMatch(t *testing.T) {
	got := detectSBOMFormat([]byte("just some random text"))
	if got != "" {
		t.Errorf("expected empty, got %q", got)
	}
}

func TestNewSBOMDiscoverer(t *testing.T) {
	d := NewSBOMDiscoverer()
	if d == nil {
		t.Fatal("expected non-nil discoverer")
	}
	if _, ok := d.(*OCISBOMDiscoverer); !ok {
		t.Error("expected *OCISBOMDiscoverer")
	}
}

func TestDetectFromContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"spdx url", `"documentNamespace": "https://spdx.dev/Document/test"`, "spdx"},
		{"spdx ref", `"SPDXID": "SPDXRef-Package"`, "spdx"},
		{"cyclonedx upper", `"bomFormat": "CycloneDX"`, "cyclonedx"},
		{"cyclonedx lower", `"specVersion": "cyclonedx/1.4"`, "cyclonedx"},
		{"no match", `{"name": "test"}`, ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := detectFromContent(tc.input)
			if got != tc.expected {
				t.Errorf("detectFromContent(%q) = %q, want %q", tc.input, got, tc.expected)
			}
		})
	}
}
