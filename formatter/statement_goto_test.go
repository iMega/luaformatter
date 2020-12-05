package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseGoto(t *testing.T) {
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
			name: "goto statement",
			args: args{
				code: []byte(`
goto label
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]block{
						0: {
							Statement: statement{
								Goto: &gotoStatement{
									Element: &element{
										Token: &lexmachine.Token{
											Type:        nID,
											Value:       "label",
											Lexeme:      []byte("label"),
											TC:          6,
											StartLine:   2,
											StartColumn: 6,
											EndLine:     2,
											EndColumn:   10,
										},
									},
								},
							},
						},
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
