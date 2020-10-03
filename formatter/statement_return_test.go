package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseReturn(t *testing.T) {
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
			name: "return statement two var",
			args: args{
				code: []byte(`
return 1+2, b
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Return: &returnStatement{
							Explist: &explist{
								List: []*exp{
									{
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nNumber,
												Value:       "1",
												Lexeme:      []byte("1"),
												TC:          8,
												StartLine:   2,
												StartColumn: 8,
												EndLine:     2,
												EndColumn:   8,
											},
										},
										Binop: &element{
											Token: &lexmachine.Token{
												Type:        nAddition,
												Value:       "+",
												Lexeme:      []byte("+"),
												TC:          9,
												StartLine:   2,
												StartColumn: 9,
												EndLine:     2,
												EndColumn:   9,
											},
										},
										Exp: &exp{
											Element: &element{
												Token: &lexmachine.Token{
													Type:        nNumber,
													Value:       "2",
													Lexeme:      []byte("2"),
													TC:          10,
													StartLine:   2,
													StartColumn: 10,
													EndLine:     2,
													EndColumn:   10,
												},
											},
										},
									},
									{
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nID,
												Value:       "b",
												Lexeme:      []byte("b"),
												TC:          13,
												StartLine:   2,
												StartColumn: 13,
												EndLine:     2,
												EndColumn:   13,
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
