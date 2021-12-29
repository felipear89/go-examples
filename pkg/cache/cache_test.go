package cache

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

func TestWithCache(t *testing.T) {
	type Mock struct {
	}
	type args struct {
		cache *gocache.Cache
		k     string
		f     func() (*Mock, error)
	}
	tests := []struct {
		name    string
		args    args
		want    *Mock
		wantErr bool
	}{
		{
			name: "return error",
			args: args{
				cache: gocache.New(5*time.Minute, 10*time.Minute),
				k:     "foo",
				f: func() (*Mock, error) {
					return nil, errors.New("error")
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "set cache",
			args: args{
				cache: gocache.New(5*time.Minute, 10*time.Minute),
				k:     "foo",
				f: func() (*Mock, error) {
					return &Mock{}, nil
				},
			},
			want:    &Mock{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := WithCache(tt.args.cache, tt.args.k, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, result)
		})
	}
}
