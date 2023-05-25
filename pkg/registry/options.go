package registry

type Options struct {

	// registry服务器的地址
	RegistryURL string

	// 请求时使用的代理IP
	Proxy string
}

const DefaultRegistryURL = "https://registry.npmjs.org"

func NewOptions() *Options {
	return &Options{
		RegistryURL: DefaultRegistryURL,
	}
}

func (x *Options) SetRegistryURL(registryURL string) *Options {
	x.RegistryURL = registryURL
	return x
}

func (x *Options) SetProxy(proxy string) *Options {
	x.Proxy = proxy
	return x
}
