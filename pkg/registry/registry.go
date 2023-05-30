package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/crawler-go-go-go/go-requests"
	"github.com/scagogogo/npm-crawler/pkg/model"
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
	bytes, err := requests.GetBytes(ctx, targetUrl)
	if err != nil {
		return nil, err
	}
	r := &models.RegistryInformation{}
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (x *Registry) GetPackageInformation(ctx context.Context, packageName string) (*models.Package, error) {
	targetUrl := fmt.Sprintf("%s/%s", x.options.RegistryURL, packageName)
	bytes, err := requests.GetBytes(ctx, targetUrl)
	if err != nil {
		return nil, err
	}
	r := &models.Package{}
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
