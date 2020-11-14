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
			skip: false,
			name: "global assignment statement with one function",
			args: args{
				code: []byte(`
myvar = function() end
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
														Blocks: make(map[uint64]block),
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
			name: "global assignment statement with one var",
			args: args{
				code: []byte(`
myvar = 1
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
					},
					Qty: 1,
				},
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "global assignment statement with two vars",
			args: args{
				code: []byte(`
a, b = 1, 2
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
									IsLocal: false,
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
												},
											},
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "b",
														Lexeme:      []byte("b"),
														TC:          4,
														StartLine:   2,
														StartColumn: 4,
														EndLine:     2,
														EndColumn:   4,
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
														TC:          8,
														StartLine:   2,
														StartColumn: 8,
														EndLine:     2,
														EndColumn:   8,
													},
												},
											},
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nNumber,
														Value:       "2",
														Lexeme:      []byte("2"),
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
					},
					Qty: 1,
				},
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "local assignment statement with define one var",
			args: args{
				code: []byte(`
local a
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
			name: "local assignment statement with one var",
			args: args{
				code: []byte(`
local a = b
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
														Type:        nID,
														Value:       "b",
														Lexeme:      []byte("b"),
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
					},
					Qty: 1,
				},
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "local assignment statement with two var",
			args: args{
				code: []byte(`
local a, b
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
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "b",
														Lexeme:      []byte("b"),
														TC:          10,
														StartLine:   2,
														StartColumn: 10,
														EndLine:     2,
														EndColumn:   10,
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
			name: "local assignment statement with two var",
			args: args{
				code: []byte(`
local a, b = c, d
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
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "b",
														Lexeme:      []byte("b"),
														TC:          10,
														StartLine:   2,
														StartColumn: 10,
														EndLine:     2,
														EndColumn:   10,
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
														Type:        nID,
														Value:       "c",
														Lexeme:      []byte("c"),
														TC:          14,
														StartLine:   2,
														StartColumn: 14,
														EndLine:     2,
														EndColumn:   14,
													},
												},
											},
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "d",
														Lexeme:      []byte("d"),
														TC:          17,
														StartLine:   2,
														StartColumn: 17,
														EndLine:     2,
														EndColumn:   17,
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
			name: "local assignment statement with one funccall",
			args: args{
				code: []byte(`
local base = require "resty.core.base"
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
									IsLocal: true,
									VarList: &explist{
										List: []*exp{
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "base",
														Lexeme:      []byte("base"),
														TC:          7,
														StartLine:   2,
														StartColumn: 7,
														EndLine:     2,
														EndColumn:   10,
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
																	Value:       "require",
																	Lexeme:      []byte("require"),
																	TC:          14,
																	StartLine:   2,
																	StartColumn: 14,
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
																			Type:        nString,
																			Value:       `"resty.core.base"`,
																			Lexeme:      []byte(`"resty.core.base"`),
																			TC:          22,
																			StartLine:   2,
																			StartColumn: 22,
																			EndLine:     2,
																			EndColumn:   38,
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
