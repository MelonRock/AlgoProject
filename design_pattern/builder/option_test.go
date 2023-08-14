package builder

import (
	"fmt"
	"testing"
)

func TestNewResourcePoolConfig(t *testing.T) {
	type args struct {
		name string
		opts []ResourcePoolConfigOptFun
	}
	tests := []struct {
		name    string
		args    args
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "sucess",
			args: args{
				name: "test",
				opts: []ResourcePoolConfigOptFun{
					func(option *ResourcePoolConfigOption) {
						option.minIdle = 2
					},
					func(option *ResourcePoolConfigOption) {
						option.maxIdle = 5
					},
				},
			},
			want: &ResourcePoolConfig{
				name:     "test",
				maxTotal: 10,
				maxIdle:  5,
				minIdle:  2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewResourcePoolConfig(tt.args.name, tt.args.opts...)
			if err != nil {
				fmt.Errorf("test error%v", err)
			}
			fmt.Printf("args %v", got)
		})
	}

}
