package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseRepeat(t *testing.T) {
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
			name: "repeat statement",
			args: args{
				code: []byte(`
repeat
    break
until a
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							Repeat: &repeatStatement{
								Body: []Block{
									{
										Statement: statement{
											Break: &breakStatement{},
										},
									},
								},
								Exp: &exp{
									Element: &element{
										Token: &lexmachine.Token{
											Type:        nID,
											Value:       "a",
											Lexeme:      []byte("a"),
											TC:          24,
											StartLine:   4,
											StartColumn: 7,
											EndLine:     4,
											EndColumn:   7,
										},
									},
								},
							},
						},
					},
				},
				QtyBlocks: 1,
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
