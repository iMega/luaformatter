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
			name: "prefixexp statement with comments",
			args: args{
				code: []byte(`
--[[1]] r --[[2]] [ --[[3]] "ff" --[[4]] ] --[[5]] [ --[[6]] "dd" --[[7]] ]--[[8]].--[[9]]name--[[10]].--[[11]]name2--[[12]][--[[13]]"ee"--[[14]]]--[[15]](--[[16]])--[[17]]
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]statement{
						0: &commentStatement{
							Element: &element{
								Token: &lexmachine.Token{
									Type:        nCommentLong,
									Value:       "1",
									Lexeme:      []byte("1"),
									TC:          1,
									StartLine:   2,
									StartColumn: 1,
									EndLine:     2,
									EndColumn:   7,
								},
							},
						},
						1: &funcCallStatement{
							Prefixexp: &prefixexpStatement{
								Element: &element{
									Token: &lexmachine.Token{
										Type:        nID,
										Value:       "r",
										Lexeme:      []byte("r"),
										TC:          9,
										StartLine:   2,
										StartColumn: 9,
										EndLine:     2,
										EndColumn:   9,
									},
								},
								Comments: map[uint64]*element{
									0: {
										Token: &lexmachine.Token{
											Type:        nCommentLong,
											Value:       "2",
											Lexeme:      []byte("2"),
											TC:          11,
											StartLine:   2,
											StartColumn: 11,
											EndLine:     2,
											EndColumn:   17,
										},
									},
								},
								Prefixexp: &prefixexpStatement{
									Comments: map[uint64]*element{
										0: {
											Token: &lexmachine.Token{
												Type:        nCommentLong,
												Value:       "5",
												Lexeme:      []byte("5"),
												TC:          44,
												StartLine:   2,
												StartColumn: 44,
												EndLine:     2,
												EndColumn:   50,
											},
										},
									},
									FieldAccessorExp: &exp{
										Comments: map[uint64]*element{
											0: {
												Token: &lexmachine.Token{
													Type:        nCommentLong,
													Value:       "3",
													Lexeme:      []byte("3"),
													TC:          21,
													StartLine:   2,
													StartColumn: 21,
													EndLine:     2,
													EndColumn:   27,
												},
											},
											1: {
												Token: &lexmachine.Token{
													Type:        nCommentLong,
													Value:       "4",
													Lexeme:      []byte("4"),
													TC:          34,
													StartLine:   2,
													StartColumn: 34,
													EndLine:     2,
													EndColumn:   40,
												},
											},
										},
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nString,
												Value:       `"ff"`,
												Lexeme:      []byte(`"ff"`),
												TC:          29,
												StartLine:   2,
												StartColumn: 29,
												EndLine:     2,
												EndColumn:   32,
											},
										},
									},
									Prefixexp: &prefixexpStatement{
										Comments: map[uint64]*element{
											0: {
												Token: &lexmachine.Token{
													Type:        nCommentLong,
													Value:       "8",
													Lexeme:      []byte("8"),
													TC:          76,
													StartLine:   2,
													StartColumn: 76,
													EndLine:     2,
													EndColumn:   82,
												},
											},
										},
										FieldAccessorExp: &exp{
											Comments: map[uint64]*element{
												0: {
													Token: &lexmachine.Token{
														Type:        nCommentLong,
														Value:       "6",
														Lexeme:      []byte("6"),
														TC:          54,
														StartLine:   2,
														StartColumn: 54,
														EndLine:     2,
														EndColumn:   60,
													},
												},
												1: {
													Token: &lexmachine.Token{
														Type:        nCommentLong,
														Value:       "7",
														Lexeme:      []byte("7"),
														TC:          67,
														StartLine:   2,
														StartColumn: 67,
														EndLine:     2,
														EndColumn:   73,
													},
												},
											},
											Element: &element{
												Token: &lexmachine.Token{
													Type:        nString,
													Value:       `"dd"`,
													Lexeme:      []byte(`"dd"`),
													TC:          62,
													StartLine:   2,
													StartColumn: 62,
													EndLine:     2,
													EndColumn:   65,
												},
											},
										},
										Prefixexp: &prefixexpStatement{
											Comments: map[uint64]*element{
												0: {
													Token: &lexmachine.Token{
														Type:        nCommentLong,
														Value:       "9",
														Lexeme:      []byte("9"),
														TC:          84,
														StartLine:   2,
														StartColumn: 84,
														EndLine:     2,
														EndColumn:   90,
													},
												},
												1: {
													Token: &lexmachine.Token{
														Type:        nCommentLong,
														Value:       "10",
														Lexeme:      []byte("10"),
														TC:          95,
														StartLine:   2,
														StartColumn: 95,
														EndLine:     2,
														EndColumn:   102,
													},
												},
											},
											FieldAccessor: &element{
												Token: &lexmachine.Token{
													Type:        nID,
													Value:       "name",
													Lexeme:      []byte("name"),
													TC:          91,
													StartLine:   2,
													StartColumn: 91,
													EndLine:     2,
													EndColumn:   94,
												},
											},
											Prefixexp: &prefixexpStatement{
												Comments: map[uint64]*element{
													0: {
														Token: &lexmachine.Token{
															Type:        nCommentLong,
															Value:       "11",
															Lexeme:      []byte("11"),
															TC:          104,
															StartLine:   2,
															StartColumn: 104,
															EndLine:     2,
															EndColumn:   111,
														},
													},
													1: {
														Token: &lexmachine.Token{
															Type:        nCommentLong,
															Value:       "12",
															Lexeme:      []byte("12"),
															TC:          117,
															StartLine:   2,
															StartColumn: 117,
															EndLine:     2,
															EndColumn:   124,
														},
													},
												},
												FieldAccessor: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "name2",
														Lexeme:      []byte("name2"),
														TC:          112,
														StartLine:   2,
														StartColumn: 112,
														EndLine:     2,
														EndColumn:   116,
													},
												},
												Prefixexp: &prefixexpStatement{
													Comments: map[uint64]*element{
														0: {
															Token: &lexmachine.Token{
																Type:        nCommentLong,
																Value:       "15",
																Lexeme:      []byte("15"),
																TC:          147,
																StartLine:   2,
																StartColumn: 147,
																EndLine:     2,
																EndColumn:   154,
															},
														},
													},
													FieldAccessorExp: &exp{
														Comments: map[uint64]*element{
															0: {
																Token: &lexmachine.Token{
																	Type:        nCommentLong,
																	Value:       "13",
																	Lexeme:      []byte("13"),
																	TC:          126,
																	StartLine:   2,
																	StartColumn: 126,
																	EndLine:     2,
																	EndColumn:   133,
																},
															},
															1: {
																Token: &lexmachine.Token{
																	Type:        nCommentLong,
																	Value:       "14",
																	Lexeme:      []byte("14"),
																	TC:          138,
																	StartLine:   2,
																	StartColumn: 138,
																	EndLine:     2,
																	EndColumn:   145,
																},
															},
														},
														Element: &element{
															Token: &lexmachine.Token{
																Type:        nString,
																Value:       `"ee"`,
																Lexeme:      []byte(`"ee"`),
																TC:          134,
																StartLine:   2,
																StartColumn: 134,
																EndLine:     2,
																EndColumn:   137,
															},
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
									{
										Comments: map[uint64]*element{
											0: {
												Token: &lexmachine.Token{
													Type:        nCommentLong,
													Value:       "16",
													Lexeme:      []byte("16"),
													TC:          156,
													StartLine:   2,
													StartColumn: 156,
													EndLine:     2,
													EndColumn:   163,
												},
											},
										},
									},
								},
							},
						},
						2: &commentStatement{
							Element: &element{
								Token: &lexmachine.Token{
									Type:        nCommentLong,
									Value:       "17",
									Lexeme:      []byte("17"),
									TC:          165,
									StartLine:   2,
									StartColumn: 165,
									EndLine:     2,
									EndColumn:   172,
								},
							},
						},
					},
					Qty: 3,
				},
			},
			wantErr: false,
		},
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
				Body: &body{
					Blocks: map[uint64]statement{
						0: &funcCallStatement{
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
											FieldAccessor: &element{
												Token: &lexmachine.Token{
													Type:        nID,
													Value:       "name",
													Lexeme:      []byte("name"),
													TC:          15,
													StartLine:   2,
													StartColumn: 15,
													EndLine:     2,
													EndColumn:   18,
												},
											},
											Prefixexp: &prefixexpStatement{
												FieldAccessor: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "name2",
														Lexeme:      []byte("name2"),
														TC:          20,
														StartLine:   2,
														StartColumn: 20,
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
							},
							Explist: &explist{
								List: []*exp{
									{},
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
				Body: &body{
					Blocks: map[uint64]statement{
						0: &assignmentStatement{
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
														FieldAccessor: &element{
															Token: &lexmachine.Token{
																Type:        nID,
																Value:       "dd",
																Lexeme:      []byte("dd"),
																TC:          15,
																StartLine:   2,
																StartColumn: 15,
																EndLine:     2,
																EndColumn:   16,
															},
														},
														Prefixexp: &prefixexpStatement{
															FieldAccessor: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "ee",
																	Lexeme:      []byte("ee"),
																	TC:          18,
																	StartLine:   2,
																	StartColumn: 18,
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
