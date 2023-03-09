package nocnoc

import "context"

type NocNocClient interface {
	SettingsCategories(ctx context.Context, isHighlight, isRecommend *bool, includes *string) (SettingsCategoriesList, error)
	AddSettingsCategories(ctx context.Context, categoryId string, isHighlight, isRecommend bool) (AddSettingsCategoriesResponse, error)
	DeleteSettingsCategories(ctx context.Context, categoryId string) error
}

type nocNocClient struct {
	config NocNocClientConfig
}

type nocNocClientOption func(*nocNocClient)

func Options(options ...nocNocClientOption) nocNocClientOption {
	return func(cc *nocNocClient) {
		for _, option := range options {
			option(cc)
		}
	}
}

func WithConfig(c NocNocClientConfig) nocNocClientOption {
	return func(cc *nocNocClient) {
		cc.config = c
	}
}

func WithDefaultOptions(c NocNocClientConfig) nocNocClientOption {
	return Options(WithConfig(c))
}

func NewNocNocClientWithOptions(options ...nocNocClientOption) NocNocClient {
	cc := nocNocClient{}

	for _, option := range options {
		option(&cc)
	}

	return &cc
}

func NewNocNocClient(cfg NocNocClientConfig) NocNocClient {
	return NewNocNocClientWithOptions(WithDefaultOptions(cfg))
}
