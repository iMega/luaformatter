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
			skip: true,
			name: "condition statement one var",
			args: args{
				code: []byte(`
if a ~= 1 then
    return 22
end
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							If: &ifStatement{
								Exp: &exp{
									Element: &element{
										Token: &lexmachine.Token{
											Type:        nID,
											Value:       "a",
											Lexeme:      []byte("a"),
											TC:          4,
											StartLine:   2,
											StartColumn: 4,
											EndLine:     2,
											EndColumn:   4,
										},
									},
									Binop: &element{
										Token: &lexmachine.Token{
											Type:        nNegEq,
											Value:       keywords[nNegEq],
											Lexeme:      []byte(keywords[nNegEq]),
											TC:          6,
											StartLine:   2,
											StartColumn: 6,
											EndLine:     2,
											EndColumn:   7,
										},
									},
									Exp: &exp{
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
								Body: []Block{
									{
										Return: &returnStatement{
											Explist: &explist{
												List: []*exp{
													{
														Element: &element{
															Token: &lexmachine.Token{
																Type:        nNumber,
																Value:       "22",
																Lexeme:      []byte("22"),
																TC:          27,
																StartLine:   3,
																StartColumn: 12,
																EndLine:     3,
																EndColumn:   13,
															},
														},
													},
												},
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
		{
			skip: false,
			name: "condition statement with elseif",
			args: args{
				code: []byte(`
if a == 1 then
    print:qq "1"
elseif a == 3 then
    print "3"
end
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							If: &ifStatement{
								Exp: &exp{
									Element: &element{
										Token: &lexmachine.Token{
											Type:        nID,
											Value:       "a",
											Lexeme:      []byte("a"),
											TC:          4,
											StartLine:   2,
											StartColumn: 4,
											EndLine:     2,
											EndColumn:   4,
										},
									},
									Binop: &element{
										Token: &lexmachine.Token{
											Type:        nNegEq,
											Value:       keywords[nNegEq],
											Lexeme:      []byte(keywords[nNegEq]),
											TC:          6,
											StartLine:   2,
											StartColumn: 6,
											EndLine:     2,
											EndColumn:   7,
										},
									},
									Exp: &exp{
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
								Body: []Block{
									{
										Return: &returnStatement{
											Explist: &explist{
												List: []*exp{
													{
														Element: &element{
															Token: &lexmachine.Token{
																Type:        nNumber,
																Value:       "22",
																Lexeme:      []byte("22"),
																TC:          27,
																StartLine:   3,
																StartColumn: 12,
																EndLine:     3,
																EndColumn:   13,
															},
														},
													},
												},
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
