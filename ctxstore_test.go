package ctxstore

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateRootContext(t *testing.T) {
	tests := []struct {
		name  string
		want  context.Context
		want1 context.CancelFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GenerateRootContext(context.Background(), Options{})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRootContext() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GenerateRootContext() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
	}
	type testCase[T any] struct {
		name  string
		args  args
		want  T
		want1 bool
	}
	ctx, _ := GenerateRootContext(context.Background(), Options{})
	tests := []testCase[string]{
		{
			name:  "a",
			args:  args{ctx, "blah"},
			want:  "",
			want1: false,
		},
		{
			name:  "b",
			args:  args{ctx, "b"},
			want:  "good data",
			want1: true,
		},
	}
	Put[string](ctx, "b", "good data")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Get[string](tt.args.ctx, tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
	testsB := []testCase[*bytes.Buffer]{
		{
			name:  "c",
			args:  args{ctx, "c"},
			want:  nil,
			want1: false,
		},
		{
			name:  "d",
			args:  args{ctx, "d"},
			want:  bytes.NewBufferString("good data"),
			want1: true,
		},
	}
	Put[*bytes.Buffer](ctx, "d", bytes.NewBufferString("good data"))
	for _, tt := range testsB {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Get[*bytes.Buffer](tt.args.ctx, tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPut(t *testing.T) {
	type args[T any] struct {
		ctx context.Context
		key string
		val T
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	ctx, _ := GenerateRootContext(context.Background(), Options{})
	tests := []testCase[string]{
		{
			name: "a",
			args: args[string]{
				ctx: ctx,
				key: "a",
				val: "this is data",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Put(tt.args.ctx, tt.args.key, tt.args.val)
		})
	}
}

func TestPutCollision(t *testing.T) {
	ctx, _ := GenerateRootContext(context.Background(), Options{})
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Successfully tested panic upon type collision\n")
		} else {
			t.Errorf("Failed to panic upon type collision")
		}
	}()
	Put[int](ctx, "samekey", 0)
	Put[string](ctx, "samekey", "string now")
}
