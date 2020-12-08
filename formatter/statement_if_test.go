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
			name: "condition statement with internal function",
			args: args{
				code: []byte(`
local a = 1
local b, c
if a == 1 then
    b = function()
        return 1
    end
elseif a == 2 then
    b = function()
        return 2
    end
elseif a == 3 then
    b = function()
        return 3
    end
else
    b = function()
        return 4
    end
end
c = b()
print(c)
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]block{
						0: {
							Statement: statement{
								Assignment: &assignmentStatement{
									IsLocal: true,
									VarList: &explist{
										List: []*exp{
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "a",
														Lexeme:      []byte("a"),
														TC:          7,
														StartLine:   2,
														StartColumn: 7,
														EndLine:     2,
														EndColumn:   7,
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
														TC:          11,
														StartLine:   2,
														StartColumn: 11,
														EndLine:     2,
														EndColumn:   11,
													},
												},
											},
										},
									},
								},
							},
						},
						1: {
							Statement: statement{
								Assignment: &assignmentStatement{
									IsLocal: true,
									VarList: &explist{
										List: []*exp{
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "b",
														Lexeme:      []byte("b"),
														TC:          19,
														StartLine:   3,
														StartColumn: 7,
														EndLine:     3,
														EndColumn:   7,
													},
												},
											},
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "c",
														Lexeme:      []byte("c"),
														TC:          22,
														StartLine:   3,
														StartColumn: 10,
														EndLine:     3,
														EndColumn:   10,
													},
												},
											},
										},
									},
								},
							},
						},
						2: {
							Statement: statement{
								If: &ifStatement{
									Exp: &exp{
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nID,
												Value:       "a",
												Lexeme:      []byte("a"),
												TC:          27,
												StartLine:   4,
												StartColumn: 4,
												EndLine:     4,
												EndColumn:   4,
											},
										},
										Binop: &element{
											Token: &lexmachine.Token{
												Type:        nEquality,
												Value:       string(keywords[nEquality]),
												Lexeme:      keywords[nEquality],
												TC:          29,
												StartLine:   4,
												StartColumn: 6,
												EndLine:     4,
												EndColumn:   7,
											},
										},
										Exp: &exp{
											Element: &element{
												Token: &lexmachine.Token{
													Type:        nNumber,
													Value:       "1",
													Lexeme:      []byte("1"),
													TC:          32,
													StartLine:   4,
													StartColumn: 9,
													EndLine:     4,
													EndColumn:   9,
												},
											},
										},
									},
									Body: &body{
										Qty: 1,
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
																				Value:       "b",
																				Lexeme:      []byte("b"),
																				TC:          43,
																				StartLine:   5,
																				StartColumn: 5,
																				EndLine:     5,
																				EndColumn:   5,
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
																	Func: &functionStatement{
																		IsAnonymous: true,
																		Body: &body{
																			Qty: 1,
																			Blocks: map[uint64]block{
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
																											TC:          73,
																											StartLine:   6,
																											StartColumn: 16,
																											EndLine:     6,
																											EndColumn:   16,
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
									ElseIfPart: []*elseifStatement{
										{
											Exp: &exp{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "a",
														Lexeme:      []byte("a"),
														TC:          90,
														StartLine:   8,
														StartColumn: 8,
														EndLine:     8,
														EndColumn:   8,
													},
												},
												Binop: &element{
													Token: &lexmachine.Token{
														Type:        nEquality,
														Value:       string(keywords[nEquality]),
														Lexeme:      keywords[nEquality],
														TC:          92,
														StartLine:   8,
														StartColumn: 10,
														EndLine:     8,
														EndColumn:   11,
													},
												},
												Exp: &exp{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nNumber,
															Value:       "2",
															Lexeme:      []byte("2"),
															TC:          95,
															StartLine:   8,
															StartColumn: 13,
															EndLine:     8,
															EndColumn:   13,
														},
													},
												},
											},
											Body: &body{
												Qty: 1,
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
																						Value:       "b",
																						Lexeme:      []byte("b"),
																						TC:          106,
																						StartLine:   9,
																						StartColumn: 5,
																						EndLine:     9,
																						EndColumn:   5,
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
																			Func: &functionStatement{
																				IsAnonymous: true,
																				Body: &body{
																					Qty: 1,
																					Blocks: map[uint64]block{
																						0: {
																							Return: &returnStatement{
																								Explist: &explist{
																									List: []*exp{
																										{
																											Element: &element{
																												Token: &lexmachine.Token{
																													Type:        nNumber,
																													Value:       "2",
																													Lexeme:      []byte("2"),
																													TC:          136,
																													StartLine:   10,
																													StartColumn: 16,
																													EndLine:     10,
																													EndColumn:   16,
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
										{
											Exp: &exp{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "a",
														Lexeme:      []byte("a"),
														TC:          153,
														StartLine:   12,
														StartColumn: 8,
														EndLine:     12,
														EndColumn:   8,
													},
												},
												Binop: &element{
													Token: &lexmachine.Token{
														Type:        nEquality,
														Value:       string(keywords[nEquality]),
														Lexeme:      keywords[nEquality],
														TC:          155,
														StartLine:   12,
														StartColumn: 10,
														EndLine:     12,
														EndColumn:   11,
													},
												},
												Exp: &exp{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nNumber,
															Value:       "3",
															Lexeme:      []byte("3"),
															TC:          158,
															StartLine:   12,
															StartColumn: 13,
															EndLine:     12,
															EndColumn:   13,
														},
													},
												},
											},
											Body: &body{
												Qty: 1,
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
																						Value:       "b",
																						Lexeme:      []byte("b"),
																						TC:          169,
																						StartLine:   13,
																						StartColumn: 5,
																						EndLine:     13,
																						EndColumn:   5,
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
																			Func: &functionStatement{
																				IsAnonymous: true,
																				Body: &body{
																					Qty: 1,
																					Blocks: map[uint64]block{
																						0: {
																							Return: &returnStatement{
																								Explist: &explist{
																									List: []*exp{
																										{
																											Element: &element{
																												Token: &lexmachine.Token{
																													Type:        nNumber,
																													Value:       "3",
																													Lexeme:      []byte("3"),
																													TC:          199,
																													StartLine:   14,
																													StartColumn: 16,
																													EndLine:     14,
																													EndColumn:   16,
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
									ElsePart: &elseStatement{
										Body: &body{
											Qty: 1,
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
																					Value:       "b",
																					Lexeme:      []byte("b"),
																					TC:          218,
																					StartLine:   17,
																					StartColumn: 5,
																					EndLine:     17,
																					EndColumn:   5,
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
																		Func: &functionStatement{
																			IsAnonymous: true,
																			Body: &body{
																				Qty: 1,
																				Blocks: map[uint64]block{
																					0: {
																						Return: &returnStatement{
																							Explist: &explist{
																								List: []*exp{
																									{
																										Element: &element{
																											Token: &lexmachine.Token{
																												Type:        nNumber,
																												Value:       "4",
																												Lexeme:      []byte("4"),
																												TC:          248,
																												StartLine:   18,
																												StartColumn: 16,
																												EndLine:     18,
																												EndColumn:   16,
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
						3: {
							Statement: statement{
								Assignment: &assignmentStatement{
									VarList: &explist{
										List: []*exp{
											{
												Prefixexp: &prefixexpStatement{
													Element: &element{
														Token: &lexmachine.Token{
															Type:        nID,
															Value:       "c",
															Lexeme:      []byte("c"),
															TC:          262,
															StartLine:   21,
															StartColumn: 1,
															EndLine:     21,
															EndColumn:   1,
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
												Prefixexp: &prefixexpStatement{
													FuncCall: &funcCallStatement{
														Prefixexp: &prefixexpStatement{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "b",
																	Lexeme:      []byte("b"),
																	TC:          266,
																	StartLine:   21,
																	StartColumn: 5,
																	EndLine:     21,
																	EndColumn:   5,
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
									},
								},
							},
						},
						4: {
							Statement: statement{
								FuncCall: &funcCallStatement{
									Prefixexp: &prefixexpStatement{
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nID,
												Value:       "print",
												Lexeme:      []byte("print"),
												TC:          270,
												StartLine:   22,
												StartColumn: 1,
												EndLine:     22,
												EndColumn:   5,
											},
										},
									},
									Explist: &explist{
										List: []*exp{
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "c",
														Lexeme:      []byte("c"),
														TC:          276,
														StartLine:   22,
														StartColumn: 7,
														EndLine:     22,
														EndColumn:   7,
													},
												},
											},
										},
									},
								},
							},
						},
					},
					Qty: 5,
				},
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "condition statement one var",
			args: args{
				code: []byte(`
if a ~= 1 then
    return 22
end
if subsystem == 'http' then
    require "resty.core.response"
    require "resty.core.phase"
end
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]block{
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
												Value:       string(keywords[nInequality]),
												Lexeme:      keywords[nInequality],
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
									Body: &body{
										Qty: 1,
										Blocks: map[uint64]block{
											0: {
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
						1: {
							Statement: statement{
								If: &ifStatement{
									Exp: &exp{
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nID,
												Value:       "subsystem",
												Lexeme:      []byte("subsystem"),
												TC:          37,
												StartLine:   5,
												StartColumn: 4,
												EndLine:     5,
												EndColumn:   12,
											},
										},
										Binop: &element{
											Token: &lexmachine.Token{
												Type:        nEquality,
												Value:       string(keywords[nEquality]),
												Lexeme:      keywords[nEquality],
												TC:          47,
												StartLine:   5,
												StartColumn: 14,
												EndLine:     5,
												EndColumn:   15,
											},
										},
										Exp: &exp{
											Element: &element{
												Token: &lexmachine.Token{
													Type:        nString,
													Value:       `'http'`,
													Lexeme:      []byte(`'http'`),
													TC:          50,
													StartLine:   5,
													StartColumn: 17,
													EndLine:     5,
													EndColumn:   22,
												},
											},
										},
									},
									Body: &body{
										Qty: 2,
										Blocks: map[uint64]block{
											0: {
												Statement: statement{
													FuncCall: &funcCallStatement{
														Prefixexp: &prefixexpStatement{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "require",
																	Lexeme:      []byte("require"),
																	TC:          66,
																	StartLine:   6,
																	StartColumn: 5,
																	EndLine:     6,
																	EndColumn:   11,
																},
															},
														},
														Explist: &explist{
															List: []*exp{
																{
																	Element: &element{
																		Token: &lexmachine.Token{
																			Type:        nString,
																			Value:       `"resty.core.response"`,
																			Lexeme:      []byte(`"resty.core.response"`),
																			TC:          74,
																			StartLine:   6,
																			StartColumn: 13,
																			EndLine:     6,
																			EndColumn:   33,
																		},
																	},
																},
															},
														},
													},
												},
											},
											1: {
												Statement: statement{
													FuncCall: &funcCallStatement{
														Prefixexp: &prefixexpStatement{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "require",
																	Lexeme:      []byte("require"),
																	TC:          100,
																	StartLine:   7,
																	StartColumn: 5,
																	EndLine:     7,
																	EndColumn:   11,
																},
															},
														},
														Explist: &explist{
															List: []*exp{
																{
																	Element: &element{
																		Token: &lexmachine.Token{
																			Type:        nString,
																			Value:       `"resty.core.phase"`,
																			Lexeme:      []byte(`"resty.core.phase"`),
																			TC:          108,
																			StartLine:   7,
																			StartColumn: 13,
																			EndLine:     7,
																			EndColumn:   30,
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
					Qty: 2,
				},
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
				Body: &body{
					Blocks: map[uint64]block{
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
												Value:       string(keywords[nEquality]),
												Lexeme:      keywords[nEquality],
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
									Body: &body{
										Blocks: map[uint64]block{
											0: {
												Statement: statement{
													Break: &breakStatement{},
												},
											},
										},
										Qty: 1,
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
														Value:       string(keywords[nEquality]),
														Lexeme:      keywords[nEquality],
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
											Body: &body{
												Blocks: map[uint64]block{
													0: {
														Statement: statement{
															Break: &breakStatement{},
														},
													},
												},
												Qty: 1,
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
				Body: &body{
					Blocks: map[uint64]block{
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
												Value:       string(keywords[nEquality]),
												Lexeme:      keywords[nEquality],
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
									Body: &body{
										Blocks: map[uint64]block{
											0: {
												Statement: statement{
													Break: &breakStatement{},
												},
											},
										},
										Qty: 1,
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
														Value:       string(keywords[nEquality]),
														Lexeme:      keywords[nEquality],
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
											Body: &body{
												Blocks: map[uint64]block{
													0: {
														Statement: statement{
															Break: &breakStatement{},
														},
													},
												},
												Qty: 1,
											},
										},
									},
									ElsePart: &elseStatement{
										Body: &body{
											Blocks: map[uint64]block{
												0: {
													Statement: statement{
														Break: &breakStatement{},
													},
												},
											},
											Qty: 1,
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
