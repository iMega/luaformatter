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
		skip bool
		name string
		args args
		want statementIntf
	}{
		{
			skip: false,
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
			skip: false,
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
			skip: false,
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
			skip: false,
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
			skip: false,
			name: "function a() [end function] b() end",
			args: args{
				prev: &element{
					Token: &lexmachine.Token{
						Type: nEnd,
					},
				},
				cur: &element{
					Token: &lexmachine.Token{
						Type: nFunction,
					},
				},
			},
			want: &functionStatement{},
		},
		{
			skip: false,
			name: "return 1+2[, b]",
			args: args{
				prev: &element{
					Token: &lexmachine.Token{
						Type: nComma,
					},
				},
				cur: &element{
					Token: &lexmachine.Token{
						Type: nID,
					},
				},
			},
			want: &exp{},
		},
	}
	for _, tt := range tests {
		if tt.skip == true {
			continue
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := getStatement(tt.args.prev, tt.args.cur); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}
