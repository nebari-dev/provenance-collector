package verify

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"

	"github.com/nebari-dev/provenance-collector/internal/report"
)

// ProvenanceChecker checks whether container images have SLSA provenance
// attestations attached via OCI referrers or cosign attestation tags.
type ProvenanceChecker interface {
	Check(ctx context.Context, imageRef string) (*report.ProvenanceInfo, error)
}

// SLSAProvenanceChecker checks for SLSA provenance using OCI referrers.
type SLSAProvenanceChecker struct{}

// NewProvenanceChecker creates a ProvenanceChecker.
func NewProvenanceChecker() ProvenanceChecker {
	return &SLSAProvenanceChecker{}
}

// Known SLSA predicate type prefixes.
var slsaPredicates = []string{
	"https://slsa.dev/provenance/",
	"https://in-toto.io/provenance/",
}

func (c *SLSAProvenanceChecker) Check(ctx context.Context, imageRef string) (*report.ProvenanceInfo, error) {
	ref, err := name.ParseReference(imageRef)
	if err != nil {
		return &report.ProvenanceInfo{}, nil
	}

	desc, err := remote.Get(ref, remote.WithContext(ctx))
	if err != nil {
		return &report.ProvenanceInfo{}, nil
	}

	// Check the OCI referrers tag fallback: <algo>-<hex>
	referrersTag := strings.Replace(desc.Digest.String(), ":", "-", 1)
	referrersRef, err := name.ParseReference(fmt.Sprintf("%s:%s", ref.Context().String(), referrersTag))
	if err != nil {
		return &report.ProvenanceInfo{}, nil
	}

	idx, err := remote.Index(referrersRef, remote.WithContext(ctx))
	if err != nil {
		return &report.ProvenanceInfo{}, nil
	}

	manifest, err := idx.IndexManifest()
	if err != nil || manifest == nil {
		return &report.ProvenanceInfo{}, nil
	}

	for _, m := range manifest.Manifests {
		at := string(m.ArtifactType)
		// Check artifactType for sigstore bundles
		if at != "" {
			predType := m.Annotations["dev.sigstore.bundle.predicateType"]
			if isSLSAPredicate(predType) {
				return &report.ProvenanceInfo{
					HasProvenance: true,
					PredicateType: predType,
				}, nil
			}
		}

		// Check annotations directly for predicate types
		for _, v := range m.Annotations {
			if isSLSAPredicate(v) {
				return &report.ProvenanceInfo{
					HasProvenance: true,
					PredicateType: v,
				}, nil
			}
		}
	}

	return &report.ProvenanceInfo{HasProvenance: false}, nil
}

func isSLSAPredicate(s string) bool {
	for _, prefix := range slsaPredicates {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}
