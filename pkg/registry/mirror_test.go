package registry

import (
	"testing"
)

func TestNewYarnRegistry(t *testing.T) {
	RegistryTest(t, NewYarnRegistry())
}

func TestNewCnpmRegistry(t *testing.T) {
	RegistryTest(t, NewCnpmRegistry())
}

func TestNewHuaWeiCloudRegistry(t *testing.T) {
	RegistryTest(t, NewHuaWeiCloudRegistry())
}

func TestNewNpmMirrorRegistry(t *testing.T) {
	RegistryTest(t, NewNpmMirrorRegistry())
}

func TestNewNpmjsComRegistry(t *testing.T) {
	RegistryTest(t, NewNpmjsComRegistry())
}

func TestNewTaoBaoRegistry(t *testing.T) {
	RegistryTest(t, NewTaoBaoRegistry())
}

func TestNewTencentRegistry(t *testing.T) {
	RegistryTest(t, NewTencentRegistry())
}
