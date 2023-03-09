package frasers

import "context"

type FrasersClient interface {
	ProjectList(context.Context, ProjectListQuery) (ProjectList, error)
	ZoneList(context.Context) (ZoneList, error)
	Project(ctx context.Context) (Project, error)
	Plot(ctx context.Context) (Plot, error)
	PlotWarranty(ctx context.Context) (PlotWarranty, error)
}

type frasersClient struct {
	config FrasersClientConfig
}

type frasersClientOption func(*frasersClient)

func Options(options ...frasersClientOption) frasersClientOption {
	return func(cc *frasersClient) {
		for _, option := range options {
			option(cc)
		}
	}
}

func WithConfig(c FrasersClientConfig) frasersClientOption {
	return func(cc *frasersClient) {
		cc.config = c
	}
}

func WithDefaultOptions(c FrasersClientConfig) frasersClientOption {
	return Options(WithConfig(c))
}

func NewFrasersClientWithOptions(options ...frasersClientOption) FrasersClient {
	cc := frasersClient{}

	for _, option := range options {
		option(&cc)
	}

	return &cc
}

func NewFrasersClient(cfg FrasersClientConfig) FrasersClient {
	return NewFrasersClientWithOptions(WithDefaultOptions(cfg))
}
