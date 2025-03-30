package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOptions(t *testing.T) {
	// 测试默认选项创建
	options := NewOptions()
	assert.NotNil(t, options)
	assert.Equal(t, DefaultRegistryURL, options.RegistryURL)
	assert.Empty(t, options.Proxy)
}

func TestSetRegistryURL(t *testing.T) {
	// 测试设置 Registry URL
	options := NewOptions()

	// 测试链式调用返回值
	result := options.SetRegistryURL("https://test-registry.example.com")
	assert.Equal(t, options, result, "应该返回自身以支持链式调用")

	// 测试实际设置的值
	assert.Equal(t, "https://test-registry.example.com", options.RegistryURL)

	// 测试设置为空字符串
	options.SetRegistryURL("")
	assert.Empty(t, options.RegistryURL)

	// 测试设置为非标准 URL
	options.SetRegistryURL("http://localhost:8080")
	assert.Equal(t, "http://localhost:8080", options.RegistryURL)
}

func TestSetProxy(t *testing.T) {
	// 测试设置代理
	options := NewOptions()

	// 测试链式调用返回值
	result := options.SetProxy("http://proxy.example.com:3128")
	assert.Equal(t, options, result, "应该返回自身以支持链式调用")

	// 测试实际设置的值
	assert.Equal(t, "http://proxy.example.com:3128", options.Proxy)

	// 测试设置为空字符串
	options.SetProxy("")
	assert.Empty(t, options.Proxy)

	// 测试设置为 socks5 代理
	options.SetProxy("socks5://127.0.0.1:1080")
	assert.Equal(t, "socks5://127.0.0.1:1080", options.Proxy)
}

func TestOptionsChaining(t *testing.T) {
	// 测试选项链式调用
	options := NewOptions().
		SetRegistryURL("https://custom-registry.org").
		SetProxy("http://proxy.example.org:8888")

	assert.NotNil(t, options)
	assert.Equal(t, "https://custom-registry.org", options.RegistryURL)
	assert.Equal(t, "http://proxy.example.org:8888", options.Proxy)

	// 测试链式调用中的顺序
	options = NewOptions().
		SetRegistryURL("https://first-registry.com").
		SetProxy("http://first-proxy.com").
		SetRegistryURL("https://second-registry.com").
		SetProxy("http://second-proxy.com")

	assert.Equal(t, "https://second-registry.com", options.RegistryURL)
	assert.Equal(t, "http://second-proxy.com", options.Proxy)
}

func TestOptionsUsage(t *testing.T) {
	// 测试在 Registry 中使用选项
	options := NewOptions().
		SetRegistryURL("https://test-usage.example.com").
		SetProxy("http://test-proxy.example.com")

	registry := NewRegistry(options)
	retrievedOptions := registry.GetOptions()

	assert.Equal(t, options, retrievedOptions)
	assert.Equal(t, "https://test-usage.example.com", retrievedOptions.RegistryURL)
	assert.Equal(t, "http://test-proxy.example.com", retrievedOptions.Proxy)
}
