package registry

// ------------------------------------------------- --------------------------------------------------------------------

const RegistryUrlYarn = "https://registry.yarnpkg.com"

// NewYarnRegistry Yarn
func NewYarnRegistry() *Registry {
	return NewRegistry(NewOptions().SetRegistryURL(RegistryUrlYarn))
}

// ------------------------------------------------- --------------------------------------------------------------------

const RegistryUrlCnpm = "http://r.cnpmjs.org"

// NewCnpmRegistry cnpm
func NewCnpmRegistry() *Registry {
	return NewRegistry(NewOptions().SetRegistryURL(RegistryUrlCnpm))
}

// ------------------------------------------------- --------------------------------------------------------------------

const RegistryUrlTaoBao = "https://registry.npm.taobao.org"

// NewTaoBaoRegistry 淘宝
func NewTaoBaoRegistry() *Registry {
	return NewRegistry(NewOptions().SetRegistryURL(RegistryUrlTaoBao))
}

// ------------------------------------------------- --------------------------------------------------------------------

const RegistryUrlNpmjsCom = "https://skimdb.npmjs.com"

// NewNpmjsComRegistry npmMirror
func NewNpmjsComRegistry() *Registry {
	return NewRegistry(NewOptions().SetRegistryURL(RegistryUrlNpmjsCom))
}

// ------------------------------------------------- --------------------------------------------------------------------

const RegistryUrlTencent = "http://mirrors.cloud.tencent.com/npm"

// NewTencentRegistry 腾讯镜像
func NewTencentRegistry() *Registry {
	return NewRegistry(NewOptions().SetRegistryURL(RegistryUrlTencent))
}

// ------------------------------------------------- --------------------------------------------------------------------

const RegistryUrlNpmMirror = "https://registry.npmmirror.com"

// NewNpmMirrorRegistry npm mirror
func NewNpmMirrorRegistry() *Registry {
	return NewRegistry(NewOptions().SetRegistryURL(RegistryUrlNpmMirror))
}

// ------------------------------------------------- --------------------------------------------------------------------

const RegistryUrlHuaWeiCloud = "https://mirrors.huaweicloud.com/repository/npm"

// NewHuaWeiCloudRegistry 华为云
func NewHuaWeiCloudRegistry() *Registry {
	return NewRegistry(NewOptions().SetRegistryURL(RegistryUrlHuaWeiCloud))
}

// ------------------------------------------------- --------------------------------------------------------------------
