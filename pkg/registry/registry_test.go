package registry

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRegistry(t *testing.T) {
	RegistryTest(t, NewRegistry())
}

func RegistryTest(t *testing.T, r *Registry) {
	assert.NotNil(t, r)

	registryInformation, err := r.GetRegistryInformation(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, registryInformation)

	packageInformation, err := r.GetPackageInformation(context.Background(), "axios")
	assert.Nil(t, err)
	assert.NotNil(t, packageInformation)

}
