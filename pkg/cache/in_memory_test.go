package cache

import (
	"reflect"
	"testing"
)

func TestInMemory_Delete(t *testing.T) {
	type fields struct {
		storage map[string][]byte
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Delete success",
			fields: fields{
				storage: map[string][]byte{
					"exist": {10},
				},
			},
			args: args{
				key: "exist",
			},
			wantErr: false,
		},
		{
			name: "Delete fail",
			fields: fields{
				storage: map[string][]byte{
					"exist": {10},
				},
			},
			args: args{
				key: "not exist",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InMemory{
				storage: tt.fields.storage,
			}
			if err := i.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemory_Get(t *testing.T) {
	type fields struct {
		storage map[string][]byte
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Get exist",
			fields: fields{
				storage: map[string][]byte{
					"exist": {10},
				},
			},
			args: args{
				key: "exist",
			},
			want:    []byte{10},
			wantErr: false,
		},
		{
			name: "Get not exist",
			fields: fields{
				storage: map[string][]byte{
					"exist": {10},
				},
			},
			args: args{
				key: "not exist",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemory{
				storage: tt.fields.storage,
			}
			got, err := i.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemory_Set(t *testing.T) {
	type fields struct {
		storage map[string][]byte
	}
	type args struct {
		key   string
		value []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Set",
			fields: fields{
				storage: map[string][]byte{},
			},
			args: args{
				key:   "exist",
				value: []byte{10},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemory{
				storage: tt.fields.storage,
			}
			if err := i.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewInMemory(t *testing.T) {
	tests := []struct {
		name string
		want *InMemory
	}{
		{
			name: "New Success",
			want: &InMemory{
				storage: map[string][]byte{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}
