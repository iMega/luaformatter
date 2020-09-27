package formatter

import (
	"testing"

	"github.com/timtadh/lexmachine"
)

func Test_exp_IsEnd(t *testing.T) {
	type fields struct {
		Table *tableconstructor
		Func  *functionStatement
		Binop *binop
		Unop  *unop
	}
	type args struct {
		prev *element
		cur  *element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				prev: &element{
					Token: &lexmachine.Token{
						Type: nNumber,
					},
				},
				cur: &element{
					Token: &lexmachine.Token{
						Type: nAddition,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &exp{
				Table: tt.fields.Table,
				Func:  tt.fields.Func,
				// Binop: tt.fields.Binop,
				// Unop:  tt.fields.Unop,
			}
			if got := s.IsEnd(tt.args.prev, tt.args.cur); got != tt.want {
				t.Errorf("exp.IsEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
