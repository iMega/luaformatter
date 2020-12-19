package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBreak(t *testing.T) {
	type args struct {
		code []byte
	}
	tests := []struct {
		skip    bool
		name    string
		args    args
		want    *document
		wantErr bool
	}{
		{
			skip: false,
			name: "break statement",
			args: args{
				code: []byte(`
break
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]statement{
						0: &breakStatement{},
					},
					Qty: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if tt.skip == true {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !assert.Equal(t, got, tt.want) {
				t.Errorf("Parse() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestSynteticBreak(t *testing.T) {
	b := &breakStatement{}

	b.AppendStatement(nil)
	b.GetBody(nil, nil)
	b.GetStatement(nil, nil)
}
