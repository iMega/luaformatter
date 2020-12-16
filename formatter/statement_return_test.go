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
			name: "return statement with funccall",
			args: args{
				code: []byte(`
return io.write(string.format(fmt, unpack(arg)))
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]block{
						0: {
							Return: &returnStatement{
								Explist: &explist{
									List: []*exp{
										{
											Prefixexp: &prefixexpStatement{
												FuncCall: &funcCallStatement{
													Prefixexp: &prefixexpStatement{
														Element: &element{
															Token: &lexmachine.Token{
																Type:        nID,
																Value:       "io",
																Lexeme:      []byte("io"),
																TC:          8,
																StartLine:   2,
																StartColumn: 8,
																EndLine:     2,
																EndColumn:   9,
															},
														},
														FieldAccessor: &element{
															Token: &lexmachine.Token{
																Type:        nID,
																Value:       "write",
																Lexeme:      []byte("write"),
																TC:          11,
																StartLine:   2,
																StartColumn: 11,
																EndLine:     2,
																EndColumn:   15,
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
																					Value:       "string",
																					Lexeme:      []byte("string"),
																					TC:          17,
																					StartLine:   2,
																					StartColumn: 17,
																					EndLine:     2,
																					EndColumn:   22,
																				},
																			},
																			FieldAccessor: &element{
																				Token: &lexmachine.Token{
																					Type:        nID,
																					Value:       "format",
																					Lexeme:      []byte("format"),
																					TC:          24,
																					StartLine:   2,
																					StartColumn: 24,
																					EndLine:     2,
																					EndColumn:   29,
																				},
																			},
																		},
																		Explist: &explist{
																			List: []*exp{
																				{
																					Element: &element{
																						Token: &lexmachine.Token{
																							Type:        nID,
																							Value:       "fmt",
																							Lexeme:      []byte("fmt"),
																							TC:          31,
																							StartLine:   2,
																							StartColumn: 31,
																							EndLine:     2,
																							EndColumn:   33,
																						},
																					},
																				},
																				{
																					Prefixexp: &prefixexpStatement{
																						FuncCall: &funcCallStatement{
																							Prefixexp: &prefixexpStatement{
																								Element: &element{
																									Token: &lexmachine.Token{
																										Type:        nID,
																										Value:       "unpack",
																										Lexeme:      []byte("unpack"),
																										TC:          36,
																										StartLine:   2,
																										StartColumn: 36,
																										EndLine:     2,
																										EndColumn:   41,
																									},
																								},
																							},
																							Explist: &explist{
																								List: []*exp{
																									{
																										Element: &element{
																											Token: &lexmachine.Token{
																												Type:        nID,
																												Value:       "arg",
																												Lexeme:      []byte("arg"),
																												TC:          43,
																												StartLine:   2,
																												StartColumn: 43,
																												EndLine:     2,
																												EndColumn:   45,
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
			name: "return statement",
			args: args{
				code: []byte(`
return
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]block{
						0: {
							Return: &returnStatement{},
						},
					},
					Qty: 1,
				},
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "return statement with two exp",
			args: args{
				code: []byte(`
return 1+2, b
`,
				),
			},
			want: &document{
				Body: &body{
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
					Qty: 1,
				},
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "return statement with func",
			args: args{
				code: []byte(`
return function () end
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]block{
						0: {
							Return: &returnStatement{
								Explist: &explist{
									List: []*exp{
										{
											Func: &functionStatement{
												IsAnonymous: true,
												FuncCall: &funcCallStatement{
													Explist: &explist{
														List: []*exp{{}},
													},
												},
												Body: &body{
													Blocks: make(map[uint64]block),
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
			name: "return statement with two func",
			args: args{
				code: []byte(`
return function () end, function () end
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]block{
						0: {
							Return: &returnStatement{
								Explist: &explist{
									List: []*exp{
										{
											Func: &functionStatement{
												IsAnonymous: true,
												FuncCall: &funcCallStatement{
													Explist: &explist{
														List: []*exp{{}},
													},
												},
												Body: &body{
													Blocks: make(map[uint64]block),
												},
											},
										},
										{
											Func: &functionStatement{
												IsAnonymous: true,
												FuncCall: &funcCallStatement{
													Explist: &explist{
														List: []*exp{{}},
													},
												},
												Body: &body{
													Blocks: make(map[uint64]block),
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
				t.Errorf("Parse() = \n%v, wasnt \n%v", got, tt.want)
			}
		})
	}
}
