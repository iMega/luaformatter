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
											Type:        nInequality,
											Value:       keywords[nInequality],
											Lexeme:      []byte(keywords[nInequality]),
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
    break
elseif b == 3 then
    break
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
											Type:        nEquality,
											Value:       keywords[nEquality],
											Lexeme:      []byte(keywords[nEquality]),
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
										Statement: statement{
											Break: &breakStatement{},
										},
									},
								},
								ElseIfPart: []*elseifStatement{
									{
										Exp: &exp{
											Element: &element{
												Token: &lexmachine.Token{
													Type:        nID,
													Value:       "b",
													Lexeme:      []byte("b"),
													TC:          33,
													StartLine:   4,
													StartColumn: 8,
													EndLine:     4,
													EndColumn:   8,
												},
											},
											Binop: &element{
												Token: &lexmachine.Token{
													Type:        nEquality,
													Value:       keywords[nEquality],
													Lexeme:      []byte(keywords[nEquality]),
													TC:          35,
													StartLine:   4,
													StartColumn: 10,
													EndLine:     4,
													EndColumn:   11,
												},
											},
											Exp: &exp{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nNumber,
														Value:       "3",
														Lexeme:      []byte("3"),
														TC:          38,
														StartLine:   4,
														StartColumn: 13,
														EndLine:     4,
														EndColumn:   13,
													},
												},
											},
										},
										Body: []Block{
											{
												Statement: statement{
													Break: &breakStatement{},
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
			name: "condition statement with elseif and else",
			args: args{
				code: []byte(`
if c == 5 then
    break
elseif d == 6 then
    break
else
    break
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
											Value:       "c",
											Lexeme:      []byte("c"),
											TC:          4,
											StartLine:   2,
											StartColumn: 4,
											EndLine:     2,
											EndColumn:   4,
										},
									},
									Binop: &element{
										Token: &lexmachine.Token{
											Type:        nEquality,
											Value:       keywords[nEquality],
											Lexeme:      []byte(keywords[nEquality]),
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
												Value:       "5",
												Lexeme:      []byte("5"),
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
										Statement: statement{
											Break: &breakStatement{},
										},
									},
								},
								ElseIfPart: []*elseifStatement{
									{
										Exp: &exp{
											Element: &element{
												Token: &lexmachine.Token{
													Type:        nID,
													Value:       "d",
													Lexeme:      []byte("d"),
													TC:          33,
													StartLine:   4,
													StartColumn: 8,
													EndLine:     4,
													EndColumn:   8,
												},
											},
											Binop: &element{
												Token: &lexmachine.Token{
													Type:        nEquality,
													Value:       keywords[nEquality],
													Lexeme:      []byte(keywords[nEquality]),
													TC:          35,
													StartLine:   4,
													StartColumn: 10,
													EndLine:     4,
													EndColumn:   11,
												},
											},
											Exp: &exp{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nNumber,
														Value:       "6",
														Lexeme:      []byte("6"),
														TC:          38,
														StartLine:   4,
														StartColumn: 13,
														EndLine:     4,
														EndColumn:   13,
													},
												},
											},
										},
										Body: []Block{
											{
												Statement: statement{
													Break: &breakStatement{},
												},
											},
										},
									},
								},
								ElsePart: &elseStatement{
									Body: []Block{
										{
											Statement: statement{
												Break: &breakStatement{},
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