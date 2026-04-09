package verify

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
	ociremote "github.com/sigstore/cosign/v2/pkg/oci/remote"

	"github.com/nebari-dev/provenance-collector/internal/report"
)

// SBOMDiscoverer checks whether container images have attached SBOM attestations.
type SBOMDiscoverer interface {
	Discover(ctx context.Context, imageRef string) (*report.SBOMInfo, error)
}

// OCISBOMDiscoverer looks for SBOM attestations in OCI registries.
type OCISBOMDiscoverer struct{}

// NewSBOMDiscoverer creates an SBOMDiscoverer that checks OCI registries.
func NewSBOMDiscoverer() SBOMDiscoverer {
	return &OCISBOMDiscoverer{}
}

func (d *OCISBOMDiscoverer) Discover(ctx context.Context, imageRef string) (*report.SBOMInfo, error) {
	ref, err := name.ParseReference(imageRef)
	if err != nil {
		return &report.SBOMInfo{}, nil
	}

	se, err := ociremote.SignedEntity(ref)
	if err != nil {
		return &report.SBOMInfo{}, nil
	}

	atts, err := se.Attestations()
	if err != nil {
		return &report.SBOMInfo{}, nil
	}

	attList, err := atts.Get()
	if err != nil || len(attList) == 0 {
		return &report.SBOMInfo{HasSBOM: false}, nil
	}

	for _, att := range attList {
		payload, err := att.Payload()
		if err != nil {
			continue
		}
		format := detectSBOMFormat(payload)
		if format != "" {
			return &report.SBOMInfo{
				HasSBOM: true,
				Format:  format,
			}, nil
		}
	}

	return &report.SBOMInfo{HasSBOM: false}, nil
}

// Known in-toto predicate types for SBOM formats.
const (
	predicateSPDX      = "https://spdx.dev/Document"
	predicateCycloneDX = "https://cyclonedx.org/bom"
)

// inTotoStatement represents the minimal structure of an in-toto attestation
// needed to extract the predicate type.
type inTotoStatement struct {
	PredicateType string          `json:"predicateType"`
	Predicate     json.RawMessage `json:"predicate"`
}

// detectSBOMFormat identifies the SBOM format from an attestation payload.
// It first tries to parse the JSON structure and check the in-toto predicateType,
// then falls back to content-based detection.
func detectSBOMFormat(payload []byte) string {
	var stmt inTotoStatement
	if err := json.Unmarshal(payload, &stmt); err == nil && stmt.PredicateType != "" {
		switch {
		case strings.HasPrefix(stmt.PredicateType, predicateSPDX):
			return "spdx"
		case strings.HasPrefix(stmt.PredicateType, predicateCycloneDX):
			return "cyclonedx"
		}

		// Check the predicate body for known SBOM markers.
		if len(stmt.Predicate) > 0 {
			if f := detectFromContent(string(stmt.Predicate)); f != "" {
				return f
			}
		}
	}

	// Fallback: scan the raw payload for known markers.
	return detectFromContent(string(payload))
}

func detectFromContent(s string) string {
	switch {
	case strings.Contains(s, "https://spdx.dev/Document") || strings.Contains(s, "SPDXRef-"):
		return "spdx"
	case strings.Contains(s, "CycloneDX") || strings.Contains(s, "cyclonedx"):
		return "cyclonedx"
	default:
		return ""
	}
}
