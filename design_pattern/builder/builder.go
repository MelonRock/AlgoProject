package builder

import "fmt"

type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigOptFun func(option *ResourcePoolConfigOption)

func NewResourcePoolConfig(name string, opts ...ResourcePoolConfigOptFun) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}
	option := &ResourcePoolConfigOption{
		maxTotal: 10,
		maxIdle:  9,
		minIdle:  1,
	}
	for _, opt := range opts {
		opt(option)
	}
	if option.maxTotal < 0 || option.maxIdle < 0 || option.minIdle < 0 {
		return nil, fmt.Errorf("args err, option: %v", option)
	}
	if option.maxIdle < option.maxIdle || option.minIdle > option.maxIdle {
		return nil, fmt.Errorf("args err, option: %v", option)
	}
	return &ResourcePoolConfig{
		name:     name,
		maxTotal: option.maxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil
}
