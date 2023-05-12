package main

import (
	"reflect"
	"testing"

	"github.com/tkanata/go-training-neko/model"
)

// mainパッケージだとテストが動かない...？
func Test_generateNeko(t *testing.T) {
	type args struct {
		n     model.Neko
		color string
	}
	tests := []struct {
		name string
		args args
		want model.Neko
	}{
		{name: "しろ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateNeko(tt.args.n, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateNeko() = %v, want %v", got, tt.want)
			}
		})
	}
}
