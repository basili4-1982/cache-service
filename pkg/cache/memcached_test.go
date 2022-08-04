package cache

import (
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/mock/gomock"
	"reflect"
	"service/pkg/mock_keyvalue"
	"testing"
)

var errFailCase = errors.New("ошибка")

func TestMemcached_Delete(t *testing.T) {
	type fields struct {
		client MemcachedClient
	}
	type args struct {
		key string
	}

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_cache.NewMockMemcachedClient(ctrl)

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   error
		clientErr error
	}{
		{
			name: "Delete success",
			fields: fields{
				client: m,
			},
			args: args{
				key: "exist",
			},
			wantErr:   nil,
			clientErr: nil,
		},
		{
			name: "Delete fail",
			fields: fields{
				client: m,
			},
			args: args{
				key: "exist",
			},
			wantErr:   errFailCase,
			clientErr: errFailCase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := tt.fields.client.(*mock_cache.MockMemcachedClient)
			cli.EXPECT().Delete(tt.args.key).Return(tt.clientErr)
			m := Memcached{
				client: tt.fields.client,
			}
			if err := m.Delete(tt.args.key); err != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemcached_Get(t *testing.T) {
	type fields struct {
		client MemcachedClient
	}
	type args struct {
		key string
	}
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_cache.NewMockMemcachedClient(ctrl)

	tests := []struct {
		name      string
		fields    fields
		args      args
		want      []byte
		wantErr   error
		clientErr error
	}{
		{
			name: "Get n",
			fields: fields{
				client: m,
			},
			args: args{
				key: "exist",
			},
			want:      []byte{10},
			wantErr:   nil,
			clientErr: nil,
		},
		{
			name: "Get not exist",
			fields: fields{
				client: m,
			},
			args: args{
				key: "not exist",
			},
			want:      nil,
			wantErr:   ErrKeyNotExist,
			clientErr: memcache.ErrMalformedKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := tt.fields.client.(*mock_cache.MockMemcachedClient)
			if tt.wantErr == nil {
				cli.EXPECT().Get(tt.args.key).Return(&memcache.Item{
					Key:        tt.args.key,
					Value:      tt.want,
					Flags:      0,
					Expiration: 0,
				}, nil)
			} else {
				cli.EXPECT().Get(tt.args.key).Return(nil, tt.clientErr)
			}
			m, _ := NewMemcached(cli)

			got, err := m.Get(tt.args.key)
			if err != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemcached_Set(t *testing.T) {
	type fields struct {
		client MemcachedClient
	}

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_cache.NewMockMemcachedClient(ctrl)

	type args struct {
		key   string
		value []byte
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   error
		clientErr error
	}{
		{
			name: "Set success",
			fields: fields{
				client: m,
			},
			args: args{
				key:   "new key",
				value: []byte("ddd"),
			},
			wantErr:   nil,
			clientErr: nil,
		},
		{
			name: "Set fail",
			fields: fields{
				client: m,
			},
			args: args{
				key:   "new key",
				value: []byte("ddd"),
			},
			wantErr:   errFailCase,
			clientErr: errFailCase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := tt.fields.client.(*mock_cache.MockMemcachedClient)
			cli.EXPECT().Set(&memcache.Item{
				Key:        tt.args.key,
				Value:      tt.args.value,
				Flags:      0,
				Expiration: 0,
			}).Return(tt.clientErr)

			m := Memcached{
				client: cli,
			}
			if err := m.Set(tt.args.key, tt.args.value); err != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewMemcached(t *testing.T) {
	type args struct {
		client MemcachedClient
	}

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_cache.NewMockMemcachedClient(ctrl)

	want, _ := NewMemcached(m)

	tests := []struct {
		name    string
		args    args
		want    *Memcached
		wantErr error
	}{
		{
			name: "New success",
			args: args{
				client: m,
			},
			want:    want,
			wantErr: nil,
		},
		{
			name: "New Fail",
			args: args{
				client: nil,
			},
			want:    nil,
			wantErr: ErrReqCli,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMemcached(tt.args.client)
			if err != tt.wantErr {
				t.Errorf("NewMemcached() = err %v,want  err %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemcached() = %v, want %v", got, tt.want)
			}
		})
	}
}
