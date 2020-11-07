package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParsePrefixexp(t *testing.T) {
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
			name: "prefixexp statement convert to function call",
			args: args{
				code: []byte(`
r["ff"]["dd"].name.name2["ee"]()
`,
				),
			},
			want: &document{
				Body: make(map[uint64]Block),
				Bod: &body{
					Blocks: map[uint64]block{
						0: {
							Statement: statement{
								FuncCall: &funcCallStatement{
									Prefixexp: &prefixexpStatement{
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nID,
												Value:       "r",
												Lexeme:      []byte("r"),
												TC:          1,
												StartLine:   2,
												StartColumn: 1,
												EndLine:     2,
												EndColumn:   1,
											},
										},
										Prefixexp: &prefixexpStatement{
											FieldAccessorExp: &exp{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nString,
														Value:       `"ff"`,
														Lexeme:      []byte(`"ff"`),
														TC:          3,
														StartLine:   2,
														StartColumn: 3,
														EndLine:     2,
														EndColumn:   6,
													},
												},
											},
											Prefixexp: &prefixexpStatement{
												FieldAccessorExp: &exp{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nString,
															Value:       `"dd"`,
															Lexeme:      []byte(`"dd"`),
															TC:          9,
															StartLine:   2,
															StartColumn: 9,
															EndLine:     2,
															EndColumn:   12,
														},
													},
												},
												Prefixexp: &prefixexpStatement{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nID,
															Value:       ".name.name2",
															Lexeme:      []byte(".name.name2"),
															TC:          14,
															StartLine:   2,
															StartColumn: 14,
															EndLine:     2,
															EndColumn:   24,
														},
													},
													Prefixexp: &prefixexpStatement{
														FieldAccessorExp: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"ee"`,
																	Lexeme:      []byte(`"ee"`),
																	TC:          26,
																	StartLine:   2,
																	StartColumn: 26,
																	EndLine:     2,
																	EndColumn:   29,
																},
															},
														},
													},
												},
											},
										},
									},
									Explist: &explist{
										List: []*exp{
											{},
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
		{
			skip: false,
			name: "prefixexp statement convert to assignment statement",
			args: args{
				code: []byte(`
a["bb"]["cc"].dd.ee["ff"], g["hh"] = 1, 2
`,
				),
			},
			want: &document{
				Body: make(map[uint64]Block),
				Bod: &body{
					Blocks: map[uint64]block{
						0: {
							Statement: statement{
								Assignment: &assignmentStatement{
									VarList: &explist{
										List: []*exp{
											{
												Prefixexp: &prefixexpStatement{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nID,
															Value:       "a",
															Lexeme:      []byte("a"),
															TC:          1,
															StartLine:   2,
															StartColumn: 1,
															EndLine:     2,
															EndColumn:   1,
														},
													},
													Prefixexp: &prefixexpStatement{
														FieldAccessorExp: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"bb"`,
																	Lexeme:      []byte(`"bb"`),
																	TC:          3,
																	StartLine:   2,
																	StartColumn: 3,
																	EndLine:     2,
																	EndColumn:   6,
																},
															},
														},
														Prefixexp: &prefixexpStatement{
															FieldAccessorExp: &exp{
																Element: &element{
																	Token: &lexmachine.Token{
																		Type:        nString,
																		Value:       `"cc"`,
																		Lexeme:      []byte(`"cc"`),
																		TC:          9,
																		StartLine:   2,
																		StartColumn: 9,
																		EndLine:     2,
																		EndColumn:   12,
																	},
																},
															},
															Prefixexp: &prefixexpStatement{
																Element: &element{
																	Token: &lexmachine.Token{
																		Type:        nID,
																		Value:       ".dd.ee",
																		Lexeme:      []byte(".dd.ee"),
																		TC:          14,
																		StartLine:   2,
																		StartColumn: 14,
																		EndLine:     2,
																		EndColumn:   19,
																	},
																},
																Prefixexp: &prefixexpStatement{
																	FieldAccessorExp: &exp{
																		Element: &element{
																			Token: &lexmachine.Token{
																				Type:        nString,
																				Value:       `"ff"`,
																				Lexeme:      []byte(`"ff"`),
																				TC:          21,
																				StartLine:   2,
																				StartColumn: 21,
																				EndLine:     2,
																				EndColumn:   24,
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
											{
												Prefixexp: &prefixexpStatement{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nID,
															Value:       "g",
															Lexeme:      []byte("g"),
															TC:          28,
															StartLine:   2,
															StartColumn: 28,
															EndLine:     2,
															EndColumn:   28,
														},
													},
													FieldAccessorExp: &exp{
														Element: &element{
															Token: &lexmachine.Token{
																Type:        nString,
																Value:       `"hh"`,
																Lexeme:      []byte(`"hh"`),
																TC:          30,
																StartLine:   2,
																StartColumn: 30,
																EndLine:     2,
																EndColumn:   33,
															},
														},
													},
												},
											},
										},
									},
									HasEqPart: true,
									Explist: &explist{
										List: []*exp{
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nNumber,
														Value:       "1",
														Lexeme:      []byte("1"),
														TC:          38,
														StartLine:   2,
														StartColumn: 38,
														EndLine:     2,
														EndColumn:   38,
													},
												},
											},
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nNumber,
														Value:       "2",
														Lexeme:      []byte("2"),
														TC:          41,
														StartLine:   2,
														StartColumn: 41,
														EndLine:     2,
														EndColumn:   41,
													},
												},
											},
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
