package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseAssignment(t *testing.T) {
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
			name: "assignment statement one var",
			args: args{
				code: []byte(`
myvar = 1
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
		{
			skip: true,
			name: "assignment statement two vars",
			args: args{
				code: []byte(`
myvar = 1
myvar2 = 2
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
					1: {
						Statement: statement{
							Assignment: &assignmentStatement{
								LastTokenType: nNumber,
								Namelist: []*element{
									{
										Token: &lexmachine.Token{
											Type:        nID,
											Value:       "myvar2",
											Lexeme:      []byte("myvar2"),
											TC:          11,
											StartLine:   3,
											StartColumn: 1,
											EndLine:     3,
											EndColumn:   6,
										},
									},
								},
								EqPart: &element{
									Token: &lexmachine.Token{
										Type:        nEq,
										Value:       keywords[nEq],
										Lexeme:      []byte(keywords[nEq]),
										TC:          18,
										StartLine:   3,
										StartColumn: 8,
										EndLine:     3,
										EndColumn:   8,
									},
								},
								Explist: []*exp{
									{
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nNumber,
												Value:       "2",
												Lexeme:      []byte("2"),
												TC:          20,
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
				QtyBlocks: 2,
			},
			wantErr: false,
		},
		{
			skip: true,
			name: "assignment statement",
			args: args{
				code: []byte(`
myvar3, myvar4 = 3, 4
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
											Value:       "myvar3",
											Lexeme:      []byte("myvar3"),
											TC:          1,
											StartLine:   2,
											StartColumn: 1,
											EndLine:     2,
											EndColumn:   6,
										},
									},
									{
										Token: &lexmachine.Token{
											Type:        nComma,
											Value:       keywords[nComma],
											Lexeme:      []byte(keywords[nComma]),
											TC:          7,
											StartLine:   2,
											StartColumn: 7,
											EndLine:     2,
											EndColumn:   7,
										},
									},
									{
										Token: &lexmachine.Token{
											Type:        nID,
											Value:       "myvar4",
											Lexeme:      []byte("myvar4"),
											TC:          9,
											StartLine:   2,
											StartColumn: 9,
											EndLine:     2,
											EndColumn:   14,
										},
									},
								},
								EqPart: &element{
									Token: &lexmachine.Token{
										Type:        nEq,
										Value:       keywords[nEq],
										Lexeme:      []byte(keywords[nEq]),
										TC:          16,
										StartLine:   2,
										StartColumn: 16,
										EndLine:     2,
										EndColumn:   16,
									},
								},
								Explist: []*exp{
									{
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nNumber,
												Value:       "3",
												Lexeme:      []byte("3"),
												TC:          18,
												StartLine:   2,
												StartColumn: 18,
												EndLine:     2,
												EndColumn:   18,
											},
										},
									},
									{
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nNumber,
												Value:       "4",
												Lexeme:      []byte("4"),
												TC:          21,
												StartLine:   2,
												StartColumn: 21,
												EndLine:     2,
												EndColumn:   21,
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
			skip: true,
			name: "assignment statement one function",
			args: args{
				code: []byte(`
myvar = function () end
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
		{
			skip: true,
			name: "assignment statement one function",
			args: args{
				code: []byte(`
myvar = function ()
    return function()
    end
end
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							Assignment: &assignmentStatement{
								LastTokenType: nEq,
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
										Func: &functionStatement{
											IDStatement: &element{
												Token: &lexmachine.Token{
													Type:        nFunction,
													Value:       keywords[nFunction],
													Lexeme:      []byte(keywords[nFunction]),
													TC:          9,
													StartLine:   2,
													StartColumn: 9,
													EndLine:     2,
													EndColumn:   16,
												},
											},
											Body: []Block{
												{
													Return: &returnStatement{
														Explist: []*exp{
															{
																Func: &functionStatement{
																	IDStatement: &element{
																		Token: &lexmachine.Token{
																			Type:        nFunction,
																			Value:       keywords[nFunction],
																			Lexeme:      []byte(keywords[nFunction]),
																			TC:          32,
																			StartLine:   3,
																			StartColumn: 12,
																			EndLine:     3,
																			EndColumn:   19,
																		},
																	},
																	EndElement: &element{
																		Token: &lexmachine.Token{
																			Type:        nEnd,
																			Value:       keywords[nEnd],
																			Lexeme:      []byte(keywords[nEnd]),
																			TC:          47,
																			StartLine:   4,
																			StartColumn: 5,
																			EndLine:     4,
																			EndColumn:   7,
																		},
																	},
																	Anonymous: true,
																},
															},
														},
													},
												},
											},
											EndElement: &element{
												Token: &lexmachine.Token{
													Type:        nEnd,
													Value:       keywords[nEnd],
													Lexeme:      []byte(keywords[nEnd]),
													TC:          51,
													StartLine:   5,
													StartColumn: 1,
													EndLine:     5,
													EndColumn:   3,
												},
											},
											Anonymous: true,
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
			name: "assignment statement one function",
			args: args{
				code: []byte(`
myvar = function ()
    return function()
        return 1
    end
end
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							Assignment: &assignmentStatement{
								LastTokenType: nEq,
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
										Func: &functionStatement{
											IDStatement: &element{
												Token: &lexmachine.Token{
													Type:        nFunction,
													Value:       keywords[nFunction],
													Lexeme:      []byte(keywords[nFunction]),
													TC:          9,
													StartLine:   2,
													StartColumn: 9,
													EndLine:     2,
													EndColumn:   16,
												},
											},
											Body: []Block{
												{
													Return: &returnStatement{
														Explist: []*exp{
															{
																Func: &functionStatement{
																	IDStatement: &element{
																		Token: &lexmachine.Token{
																			Type:        nFunction,
																			Value:       keywords[nFunction],
																			Lexeme:      []byte(keywords[nFunction]),
																			TC:          32,
																			StartLine:   3,
																			StartColumn: 12,
																			EndLine:     3,
																			EndColumn:   19,
																		},
																	},
																	Body: []Block{
																		{
																			Return: &returnStatement{Explist: []*exp{
																				{
																					Element: &element{Token: &lexmachine.Token{
																						Type:        nNumber,
																						Value:       "1",
																						Lexeme:      []byte("1"),
																						TC:          58,
																						StartLine:   4,
																						StartColumn: 16,
																						EndLine:     4,
																						EndColumn:   16,
																					}},
																				},
																			}},
																		},
																	},
																	EndElement: &element{
																		Token: &lexmachine.Token{
																			Type:        nEnd,
																			Value:       keywords[nEnd],
																			Lexeme:      []byte(keywords[nEnd]),
																			TC:          64,
																			StartLine:   5,
																			StartColumn: 5,
																			EndLine:     5,
																			EndColumn:   7,
																		},
																	},
																	Anonymous: true,
																},
															},
														},
													},
												},
											},
											EndElement: &element{
												Token: &lexmachine.Token{
													Type:        nEnd,
													Value:       keywords[nEnd],
													Lexeme:      []byte(keywords[nEnd]),
													TC:          68,
													StartLine:   6,
													StartColumn: 1,
													EndLine:     6,
													EndColumn:   3,
												},
											},
											Anonymous: true,
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
