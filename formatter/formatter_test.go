package formatter

import (
	"reflect"
	"testing"

	"github.com/timtadh/lexmachine"
)

func Test_getStatement(t *testing.T) {
	type args struct {
		last    *element
		current *element
	}
	tests := []struct {
		name string
		args args
		want statementIntf
	}{
		// {
		// 	name: ",myvar",
		// 	args: args{
		// 		last: &element{
		// 			Token: &lexmachine.Token{
		// 				Type:  nComma,
		// 				Value: ",",
		// 			},
		// 		},
		// 		current: &element{
		// 			Token: &lexmachine.Token{
		// 				Type:  nID,
		// 				Value: "myvar",
		// 			},
		// 		},
		// 	},
		// 	want: nil,
		// },
		{
			name: "= function",
			args: args{
				last: &element{
					Token: &lexmachine.Token{
						Type:  nEq,
						Value: "=",
					},
				},
				current: &element{
					Token: &lexmachine.Token{
						Type:  nFunction,
						Value: "function",
					},
				},
			},
			want: &functionStatement{},
		},
		{
			name: "function",
			args: args{
				last: nil,
				current: &element{
					Token: &lexmachine.Token{
						Type:  nFunction,
						Value: "function",
					},
				},
			},
			want: &functionStatement{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStatement(tt.args.last, tt.args.current); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}
