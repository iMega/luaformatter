package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseIf(t *testing.T) {
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
			name: "condition statement one var",
			args: args{
				code: []byte(`
if a ~= 1 then
    return 1
end
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							Assignment: &assignmentStatement{
								LastTokenType: nNumber,
								Namelist: []*element{
									{
										Token: &lexmachine.Token{
											Type:        nID,
											Value:       "myvar",
											Lexeme:      []byte("myvar"),
											TC:          1,
											StartLine:   2,
											StartColumn: 1,
											EndLine:     2,
											EndColumn:   5,
										},
									},
								},
								EqPart: &element{
									Token: &lexmachine.Token{
										Type:        nEq,
										Value:       keywords[nEq],
										Lexeme:      []byte(keywords[nEq]),
										TC:          7,
										StartLine:   2,
										StartColumn: 7,
										EndLine:     2,
										EndColumn:   7,
									},
								},
								Explist: []*exp{
									{
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nNumber,
												Value:       "1",
												Lexeme:      []byte("1"),
												TC:          9,
												StartLine:   2,
												StartColumn: 9,
												EndLine:     2,
												EndColumn:   9,
											},
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
			got, err := Parse(tt.args.code)
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
