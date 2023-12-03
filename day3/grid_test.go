package main

import (
	"reflect"
	"testing"
)

func Test_grid_SubGrid(t *testing.T) {
	type args struct {
		x int
		y int
		w int
		h int
	}
	tests := []struct {
		name string
		g    Grid
		args args
		want Grid
	}{
		{
			name: "Simple 1",
			g:    LoadFile("input-test1.txt"),
			args: args{
				x: 2,
				y: 1,
				w: 5,
				h: 5,
			},
			want: LoadString(`.*...
35..6
....#
7*...
...+.`),
		},
		{
			name: "Negative 1",
			g:    LoadFile("input-test1.txt"),
			args: args{
				x: -3,
				y: -4,
				w: 5,
				h: 5,
			},
			want: LoadString(`46`),
		},
		{
			name: "Negative 2",
			g:    LoadFile("input-test1.txt"),
			args: args{
				x: -1,
				y: -1,
				w: 5,
				h: 3,
			},
			want: LoadString(`467.
...*`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.SubGrid(tt.args.x, tt.args.y, tt.args.w, tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubGrid() = %v, want %v", got, tt.want)
			} else {
				t.Log(got.String())
			}
		})
	}
}
