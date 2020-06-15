package array

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "test2",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
				},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "test3",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "test4",
			args: args{
				matrix: [][]int{
					{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "test5",
			args: args{
				matrix: [][]int{},
			},
			want: []int{},
		},
		{
			name: "test2",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
			},
			want: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
