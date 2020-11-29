package formatter

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		s []int
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found 1",
			args: args{s: []int{1, 2, 3}, n: 1},
			want: true,
		},
		{
			name: "found 2",
			args: args{s: []int{1, 2, 3}, n: 2},
			want: true,
		},
		{
			name: "found 3",
			args: args{s: []int{1, 2, 3}, n: 3},
			want: true,
		},
		{
			name: "not found 0",
			args: args{s: []int{1, 2, 3}, n: 0},
			want: false,
		},
		{
			name: "not found 4",
			args: args{s: []int{1, 2, 3}, n: 4},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
