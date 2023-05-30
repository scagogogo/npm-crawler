package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/crawler-go-go-go/go-requests"
	"github.com/scagogogo/npm-crawler/pkg/models"
)

type Registry struct {
	options *Options
}

func NewRegistry(options ...*Options) *Registry {
	if len(options) == 0 {
		options = append(options, NewOptions())
	}
	return &Registry{
		options: options[0],
	}
}

// GetRegistryInformation 获取registry的状态
func (x *Registry) GetRegistryInformation(ctx context.Context) (*models.RegistryInformation, error) {
	targetUrl := x.options.RegistryURL
	bytes, err := x.getBytes(ctx, targetUrl)
	if err != nil {
		return nil, err
	}
	return unmarshalJson[*models.RegistryInformation](bytes)
}

func (x *Registry) GetPackageInformation(ctx context.Context, packageName string) (*models.Package, error) {
	targetUrl := fmt.Sprintf("%s/%s", x.options.RegistryURL, packageName)
	bytes, err := x.getBytes(ctx, targetUrl)
	if err != nil {
		return nil, err
	}
	return unmarshalJson[*models.Package](bytes)
}

func unmarshalJson[T any](bytes []byte) (T, error) {
	var r T
	err := json.Unmarshal(bytes, &r)
	if err != nil {
		var zero T
		return zero, err
	}
	return r, nil
}

// 内部使用统一的方法来请求
func (x *Registry) getBytes(ctx context.Context, targetUrl string) ([]byte, error) {
	options := requests.NewOptions[any, []byte](targetUrl, requests.BytesResponseHandler())
	if x.options.Proxy != "" {
		options.AppendRequestSetting(requests.RequestSettingProxy(x.options.Proxy))
	}
	return requests.SendRequest[any, []byte](ctx, options)
}
