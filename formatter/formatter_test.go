package formatter

import (
	"reflect"
	"testing"

	"github.com/timtadh/lexmachine"
)

func Test_getStatement(t *testing.T) {
	type args struct {
		prev *element
		cur  *element
	}
	tests := []struct {
		name string
		args args
		want statementIntf
	}{
		{
			name: "[function] a() end",
			args: args{
				cur: &element{
					Token: &lexmachine.Token{
						Type: nFunction,
					},
				},
			},
			want: &functionStatement{},
		},
		{
			name: "[return]",
			args: args{
				cur: &element{
					Token: &lexmachine.Token{
						Type: nReturn,
					},
				},
			},
			want: &returnStatement{},
		},
		{
			name: "[return function] a() end",
			args: args{
				prev: &element{
					Token: &lexmachine.Token{
						Type: nReturn,
					},
				},
				cur: &element{
					Token: &lexmachine.Token{
						Type: nFunction,
					},
				},
			},
			want: &explist{},
		},
		{
			name: "[return 1]",
			args: args{
				prev: &element{
					Token: &lexmachine.Token{
						Type: nReturn,
					},
				},
				cur: &element{
					Token: &lexmachine.Token{
						Type: nNumber,
					},
				},
			},
			want: &explist{},
		},
		{
			name: "function a() [end function] b() end",
			args: args{
				prev: &element{
					Token: &lexmachine.Token{
						Type: nReturn,
					},
				},
				cur: &element{
					Token: &lexmachine.Token{
						Type: nNumber,
					},
				},
			},
			want: &explist{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStatement(tt.args.prev, tt.args.cur); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}
