package certificates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCaCertificateAsn10id(t *testing.T) {
	assert.Equal(t, []int{1, 3, 6, 1, 4, 1, 59193, 100, 100, 100}, GetCaCertificateAsn10id())
}

func TestGetTechCertificateAsn10id(t *testing.T) {
	assert.Equal(t, []int{1, 3, 6, 1, 4, 1, 59193, 100, 100, 200}, GetTechCertificateAsn10id())
}

func TestGetDeviceCertificateAsn10id(t *testing.T) {
	assert.Equal(t, []int{1, 3, 6, 1, 4, 1, 59193, 100, 100, 300}, GetDeviceCertificateAsn10id())
}
