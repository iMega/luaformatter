package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseFor(t *testing.T) {
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
			name: "for generic statement",
			args: args{
				code: []byte(`
for i, pkg in ipairs(packages) do
    for name, version in pairs(pkg) do
        print(version)
    end
end
`,
				),
			},
			want: &document{
				Body: make(map[uint64]Block),
				Bod: &body{
					Blocks: map[uint64]block{
						0: {
							Statement: statement{
								For: &forStatement{
									FieldList: &fieldlist{
										List: []*field{
											{
												Key: &exp{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nID,
															Value:       "i",
															Lexeme:      []byte("i"),
															TC:          5,
															StartLine:   2,
															StartColumn: 5,
															EndLine:     2,
															EndColumn:   5,
														},
													},
												},
											},
											{
												Key: &exp{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nID,
															Value:       "pkg",
															Lexeme:      []byte("pkg"),
															TC:          8,
															StartLine:   2,
															StartColumn: 8,
															EndLine:     2,
															EndColumn:   10,
														},
													},
												},
											},
										},
									},
									Explist: &explist{
										List: []*exp{
											{
												Prefixexp: &prefixexpStatement{
													FuncCall: &funcCallStatement{
														Prefixexp: &prefixexpStatement{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "ipairs",
																	Lexeme:      []byte("ipairs"),
																	TC:          15,
																	StartLine:   2,
																	StartColumn: 15,
																	EndLine:     2,
																	EndColumn:   20,
																},
															},
														},
														Explist: &explist{
															List: []*exp{
																{
																	Element: &element{
																		Token: &lexmachine.Token{
																			Type:        nID,
																			Value:       "packages",
																			Lexeme:      []byte("packages"),
																			TC:          22,
																			StartLine:   2,
																			StartColumn: 22,
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
										},
									},
									Body: &body{
										Qty: 1,
										Blocks: map[uint64]block{
											0: {
												Statement: statement{
													Do: &doStatement{
														Body: &body{
															Qty: 1,
															Blocks: map[uint64]block{
																0: {
																	Statement: statement{
																		For: &forStatement{
																			FieldList: &fieldlist{
																				List: []*field{
																					{
																						Key: &exp{
																							Element: &element{
																								Token: &lexmachine.Token{
																									Type:        nID,
																									Value:       "name",
																									Lexeme:      []byte("name"),
																									TC:          43,
																									StartLine:   3,
																									StartColumn: 9,
																									EndLine:     3,
																									EndColumn:   12,
																								},
																							},
																						},
																					},
																					{
																						Key: &exp{
																							Element: &element{
																								Token: &lexmachine.Token{
																									Type:        nID,
																									Value:       "version",
																									Lexeme:      []byte("version"),
																									TC:          49,
																									StartLine:   3,
																									StartColumn: 15,
																									EndLine:     3,
																									EndColumn:   21,
																								},
																							},
																						},
																					},
																				},
																			},
																			Explist: &explist{
																				List: []*exp{
																					{
																						Prefixexp: &prefixexpStatement{
																							FuncCall: &funcCallStatement{
																								Prefixexp: &prefixexpStatement{
																									Element: &element{
																										Token: &lexmachine.Token{
																											Type:        nID,
																											Value:       "pairs",
																											Lexeme:      []byte("pairs"),
																											TC:          60,
																											StartLine:   3,
																											StartColumn: 26,
																											EndLine:     3,
																											EndColumn:   30,
																										},
																									},
																								},
																								Explist: &explist{
																									List: []*exp{
																										{
																											Element: &element{
																												Token: &lexmachine.Token{
																													Type:        nID,
																													Value:       "pkg",
																													Lexeme:      []byte("pkg"),
																													TC:          66,
																													StartLine:   3,
																													StartColumn: 32,
																													EndLine:     3,
																													EndColumn:   34,
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
																			Body: &body{
																				Qty: 1,
																				Blocks: map[uint64]block{
																					0: {
																						Statement: statement{
																							Do: &doStatement{
																								Body: &body{
																									Qty: 1,
																									Blocks: map[uint64]block{
																										0: {
																											Statement: statement{
																												FuncCall: &funcCallStatement{
																													Prefixexp: &prefixexpStatement{
																														Element: &element{
																															Token: &lexmachine.Token{
																																Type:        nID,
																																Value:       "print",
																																Lexeme:      []byte("print"),
																																TC:          82,
																																StartLine:   4,
																																StartColumn: 9,
																																EndLine:     4,
																																EndColumn:   13,
																															},
																														},
																													},
																													Explist: &explist{
																														List: []*exp{
																															{
																																Element: &element{
																																	Token: &lexmachine.Token{
																																		Type:        nID,
																																		Value:       "version",
																																		Lexeme:      []byte("version"),
																																		TC:          88,
																																		StartLine:   4,
																																		StartColumn: 15,
																																		EndLine:     4,
																																		EndColumn:   21,
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
			name: "for numerical statement",
			args: args{
				code: []byte(`
for i=10,1,-1 do print(i) end
`,
				),
			},
			want: &document{
				Body: make(map[uint64]Block),
				Bod: &body{
					Blocks: map[uint64]block{
						0: {
							Statement: statement{
								For: &forStatement{
									FieldList: &fieldlist{
										List: []*field{
											{
												Key: &exp{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nID,
															Value:       "i",
															Lexeme:      []byte("i"),
															TC:          5,
															StartLine:   2,
															StartColumn: 5,
															EndLine:     2,
															EndColumn:   5,
														},
													},
												},
												Val: &exp{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nNumber,
															Value:       "10",
															Lexeme:      []byte("10"),
															TC:          7,
															StartLine:   2,
															StartColumn: 7,
															EndLine:     2,
															EndColumn:   8,
														},
													},
												},
											},
											{
												Key: &exp{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nNumber,
															Value:       "1",
															Lexeme:      []byte("1"),
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
												Key: &exp{
													Exp: &exp{
														Element: &element{
															Token: &lexmachine.Token{
																Type:        nNumber,
																Value:       "1",
																Lexeme:      []byte("1"),
																TC:          13,
																StartLine:   2,
																StartColumn: 13,
																EndLine:     2,
																EndColumn:   13,
															},
														},
													},
													Unop: &element{
														Token: &lexmachine.Token{
															Type:        nSubtraction,
															Value:       "-",
															Lexeme:      []byte("-"),
															TC:          12,
															StartLine:   2,
															StartColumn: 12,
															EndLine:     2,
															EndColumn:   12,
														},
													},
												},
											},
										},
									},
									Body: &body{
										Qty: 1,
										Blocks: map[uint64]block{
											0: {
												Statement: statement{
													Do: &doStatement{
														&body{
															Qty: 1,
															Blocks: map[uint64]block{
																0: {
																	Statement: statement{
																		FuncCall: &funcCallStatement{
																			Prefixexp: &prefixexpStatement{
																				Element: &element{
																					Token: &lexmachine.Token{
																						Type:        nID,
																						Value:       "print",
																						Lexeme:      []byte("print"),
																						TC:          18,
																						StartLine:   2,
																						StartColumn: 18,
																						EndLine:     2,
																						EndColumn:   22,
																					},
																				},
																			},
																			Explist: &explist{
																				List: []*exp{
																					{
																						Element: &element{
																							Token: &lexmachine.Token{
																								Type:        nID,
																								Value:       "i",
																								Lexeme:      []byte("i"),
																								TC:          24,
																								StartLine:   2,
																								StartColumn: 24,
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
					Qty: 1,
				},
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "for generic statement",
			args: args{
				code: []byte(`
for i, pkg in ipairs(packages) do print(i) end
`,
				),
			},
			want: &document{
				Body: make(map[uint64]Block),
				Bod: &body{
					Blocks: map[uint64]block{
						0: {
							Statement: statement{
								For: &forStatement{
									FieldList: &fieldlist{
										List: []*field{
											{
												Key: &exp{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nID,
															Value:       "i",
															Lexeme:      []byte("i"),
															TC:          5,
															StartLine:   2,
															StartColumn: 5,
															EndLine:     2,
															EndColumn:   5,
														},
													},
												},
											},
											{
												Key: &exp{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nID,
															Value:       "pkg",
															Lexeme:      []byte("pkg"),
															TC:          8,
															StartLine:   2,
															StartColumn: 8,
															EndLine:     2,
															EndColumn:   10,
														},
													},
												},
											},
										},
									},
									Explist: &explist{
										List: []*exp{
											{
												Prefixexp: &prefixexpStatement{
													FuncCall: &funcCallStatement{
														Prefixexp: &prefixexpStatement{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "ipairs",
																	Lexeme:      []byte("ipairs"),
																	TC:          15,
																	StartLine:   2,
																	StartColumn: 15,
																	EndLine:     2,
																	EndColumn:   20,
																},
															},
														},
														Explist: &explist{
															List: []*exp{
																{
																	Element: &element{
																		Token: &lexmachine.Token{
																			Type:        nID,
																			Value:       "packages",
																			Lexeme:      []byte("packages"),
																			TC:          22,
																			StartLine:   2,
																			StartColumn: 22,
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
										},
									},
									Body: &body{
										Qty: 1,
										Blocks: map[uint64]block{
											0: {
												Statement: statement{
													Do: &doStatement{
														&body{
															Qty: 1,
															Blocks: map[uint64]block{
																0: {
																	Statement: statement{
																		FuncCall: &funcCallStatement{
																			Prefixexp: &prefixexpStatement{
																				Element: &element{
																					Token: &lexmachine.Token{
																						Type:        nID,
																						Value:       "print",
																						Lexeme:      []byte("print"),
																						TC:          35,
																						StartLine:   2,
																						StartColumn: 35,
																						EndLine:     2,
																						EndColumn:   39,
																					},
																				},
																			},
																			Explist: &explist{
																				List: []*exp{
																					{
																						Element: &element{
																							Token: &lexmachine.Token{
																								Type:        nID,
																								Value:       "i",
																								Lexeme:      []byte("i"),
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
