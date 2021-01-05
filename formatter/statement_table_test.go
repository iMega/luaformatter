package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseTable(t *testing.T) {
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
			name: "table statement",
			args: args{
				code: []byte(`
table = {
    ["1394-E"] = val1,
    ["UTF-8"] = val2,
    ["and"] = val3,
}
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
													Value:       "table",
													Lexeme:      []byte("table"),
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
										Table: &tableStatement{
											FieldList: &fieldlist{
												List: []*field{
													{
														Square: true,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"1394-E"`,
																	Lexeme:      []byte(`"1394-E"`),
																	TC:          16,
																	StartLine:   3,
																	StartColumn: 6,
																	EndLine:     3,
																	EndColumn:   13,
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val1",
																	Lexeme:      []byte("val1"),
																	TC:          28,
																	StartLine:   3,
																	StartColumn: 18,
																	EndLine:     3,
																	EndColumn:   21,
																},
															},
														},
													},
													{
														Square: true,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"UTF-8"`,
																	Lexeme:      []byte(`"UTF-8"`),
																	TC:          39,
																	StartLine:   4,
																	StartColumn: 6,
																	EndLine:     4,
																	EndColumn:   12,
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val2",
																	Lexeme:      []byte("val2"),
																	TC:          50,
																	StartLine:   4,
																	StartColumn: 17,
																	EndLine:     4,
																	EndColumn:   20,
																},
															},
														},
													},
													{
														Square: true,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"and"`,
																	Lexeme:      []byte(`"and"`),
																	TC:          61,
																	StartLine:   5,
																	StartColumn: 6,
																	EndLine:     5,
																	EndColumn:   10,
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val3",
																	Lexeme:      []byte("val3"),
																	TC:          70,
																	StartLine:   5,
																	StartColumn: 15,
																	EndLine:     5,
																	EndColumn:   18,
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
			name: "table statement",
			args: args{
				code: []byte(`
table = {
    name = "Jack",
    ["1394-E"] = val1,
    ["UTF-8"] = val2,
    ["and"] = val2,
}
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
													Value:       "table",
													Lexeme:      []byte("table"),
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
										Table: &tableStatement{
											FieldList: &fieldlist{
												List: []*field{
													{
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "name",
																	Lexeme:      []byte("name"),
																	TC:          15,
																	StartLine:   3,
																	StartColumn: 5,
																	EndLine:     3,
																	EndColumn:   8,
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"Jack"`,
																	Lexeme:      []byte(`"Jack"`),
																	TC:          22,
																	StartLine:   3,
																	StartColumn: 12,
																	EndLine:     3,
																	EndColumn:   17,
																},
															},
														},
													},
													{
														Square: true,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"1394-E"`,
																	Lexeme:      []byte(`"1394-E"`),
																	TC:          35,
																	StartLine:   4,
																	StartColumn: 6,
																	EndLine:     4,
																	EndColumn:   13,
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val1",
																	Lexeme:      []byte("val1"),
																	TC:          47,
																	StartLine:   4,
																	StartColumn: 18,
																	EndLine:     4,
																	EndColumn:   21,
																},
															},
														},
													},
													{
														Square: true,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"UTF-8"`,
																	Lexeme:      []byte(`"UTF-8"`),
																	TC:          58,
																	StartLine:   5,
																	StartColumn: 6,
																	EndLine:     5,
																	EndColumn:   12,
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val2",
																	Lexeme:      []byte("val2"),
																	TC:          69,
																	StartLine:   5,
																	StartColumn: 17,
																	EndLine:     5,
																	EndColumn:   20,
																},
															},
														},
													},
													{
														Square: true,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"and"`,
																	Lexeme:      []byte(`"and"`),
																	TC:          80,
																	StartLine:   6,
																	StartColumn: 6,
																	EndLine:     6,
																	EndColumn:   10,
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val2",
																	Lexeme:      []byte("val2"),
																	TC:          89,
																	StartLine:   6,
																	StartColumn: 15,
																	EndLine:     6,
																	EndColumn:   18,
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
			name: "table statement",
			args: args{
				code: []byte(`
table = { -- 1
    name = "Jack",
    ["1394-E"] = val1,
    ["UTF-8"] = val2,
    ["and"] = val2,
}
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
													Value:       "table",
													Lexeme:      []byte("table"),
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
										Table: &tableStatement{
											FieldList: &fieldlist{
												List: []*field{
													{
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "name",
																	Lexeme:      []byte("name"),
																	TC:          20,
																	StartLine:   3,
																	StartColumn: 5,
																	EndLine:     3,
																	EndColumn:   8,
																},
															},
															Comments: map[uint64]*element{
																0: {
																	Token: &lexmachine.Token{
																		Type:        nComment,
																		Value:       "1",
																		Lexeme:      []byte("1"),
																		TC:          11,
																		StartLine:   2,
																		StartColumn: 11,
																		EndLine:     2,
																		EndColumn:   14,
																	},
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"Jack"`,
																	Lexeme:      []byte(`"Jack"`),
																	TC:          27,
																	StartLine:   3,
																	StartColumn: 12,
																	EndLine:     3,
																	EndColumn:   17,
																},
															},
														},
													},
													{
														Square: true,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"1394-E"`,
																	Lexeme:      []byte(`"1394-E"`),
																	TC:          40,
																	StartLine:   4,
																	StartColumn: 6,
																	EndLine:     4,
																	EndColumn:   13,
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val1",
																	Lexeme:      []byte("val1"),
																	TC:          52,
																	StartLine:   4,
																	StartColumn: 18,
																	EndLine:     4,
																	EndColumn:   21,
																},
															},
														},
													},
													{
														Square: true,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"UTF-8"`,
																	Lexeme:      []byte(`"UTF-8"`),
																	TC:          63,
																	StartLine:   5,
																	StartColumn: 6,
																	EndLine:     5,
																	EndColumn:   12,
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val2",
																	Lexeme:      []byte("val2"),
																	TC:          74,
																	StartLine:   5,
																	StartColumn: 17,
																	EndLine:     5,
																	EndColumn:   20,
																},
															},
														},
													},
													{
														Square: true,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"and"`,
																	Lexeme:      []byte(`"and"`),
																	TC:          85,
																	StartLine:   6,
																	StartColumn: 6,
																	EndLine:     6,
																	EndColumn:   10,
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val2",
																	Lexeme:      []byte("val2"),
																	TC:          94,
																	StartLine:   6,
																	StartColumn: 15,
																	EndLine:     6,
																	EndColumn:   18,
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
			name: "table statement",
			args: args{
				code: []byte(`
table = { -- 0.1
    -- 0.2
    -- 0.3
    --[[1.1]] [ --[[1.2]] "1394-E" --[[1.3]] ] --[[1.4]] = --[[1.5]] val1 --[[1.6]] , -- 1.7
    -- 2
    --[[3.1]] [ --[[3.2]] "UTF-8" --[[3.3]] ] --[[3.4]] = --[[3.5]] val2 --[[3.6]] , -- 3.7
    val3, -- 4
    -- 5
    -- 6
}
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
													Value:       "table",
													Lexeme:      []byte("table"),
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
										Table: &tableStatement{
											FieldList: &fieldlist{
												List: []*field{
													{
														Square: true,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"1394-E"`,
																	Lexeme:      []byte(`"1394-E"`),
																	TC:          66,
																	StartLine:   5,
																	StartColumn: 27,
																	EndLine:     5,
																	EndColumn:   34,
																},
															},
															Comments: map[uint64]*element{
																0: {
																	Token: &lexmachine.Token{
																		Type:        nComment,
																		Value:       "0.1",
																		Lexeme:      []byte("0.1"),
																		TC:          11,
																		StartLine:   2,
																		StartColumn: 11,
																		EndLine:     2,
																		EndColumn:   16,
																	},
																},
																1: {
																	Token: &lexmachine.Token{
																		Type:        nComment,
																		Value:       "0.2",
																		Lexeme:      []byte("0.2"),
																		TC:          22,
																		StartLine:   3,
																		StartColumn: 5,
																		EndLine:     3,
																		EndColumn:   10,
																	},
																},
																2: {
																	Token: &lexmachine.Token{
																		Type:        nComment,
																		Value:       "0.3",
																		Lexeme:      []byte("0.3"),
																		TC:          33,
																		StartLine:   4,
																		StartColumn: 5,
																		EndLine:     4,
																		EndColumn:   10,
																	},
																},
																3: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "1.1",
																		Lexeme:      []byte("1.1"),
																		TC:          44,
																		StartLine:   5,
																		StartColumn: 5,
																		EndLine:     5,
																		EndColumn:   13,
																	},
																},
																4: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "1.2",
																		Lexeme:      []byte("1.2"),
																		TC:          56,
																		StartLine:   5,
																		StartColumn: 17,
																		EndLine:     5,
																		EndColumn:   25,
																	},
																},
																5: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "1.3",
																		Lexeme:      []byte("1.3"),
																		TC:          75,
																		StartLine:   5,
																		StartColumn: 36,
																		EndLine:     5,
																		EndColumn:   44,
																	},
																},
																6: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "1.4",
																		Lexeme:      []byte("1.4"),
																		TC:          87,
																		StartLine:   5,
																		StartColumn: 48,
																		EndLine:     5,
																		EndColumn:   56,
																	},
																},
																7: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "1.5",
																		Lexeme:      []byte("1.5"),
																		TC:          99,
																		StartLine:   5,
																		StartColumn: 60,
																		EndLine:     5,
																		EndColumn:   68,
																	},
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val1",
																	Lexeme:      []byte("val1"),
																	TC:          109,
																	StartLine:   5,
																	StartColumn: 70,
																	EndLine:     5,
																	EndColumn:   73,
																},
															},
															Comments: map[uint64]*element{
																0: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "1.6",
																		Lexeme:      []byte("1.6"),
																		TC:          114,
																		StartLine:   5,
																		StartColumn: 75,
																		EndLine:     5,
																		EndColumn:   83,
																	},
																},
															},
														},
													},
													{
														Square: true,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nString,
																	Value:       `"UTF-8"`,
																	Lexeme:      []byte(`"UTF-8"`),
																	TC:          168,
																	StartLine:   7,
																	StartColumn: 27,
																	EndLine:     7,
																	EndColumn:   33,
																},
															},
															Comments: map[uint64]*element{
																0: {
																	Token: &lexmachine.Token{
																		Type:        nComment,
																		Value:       "1.7",
																		Lexeme:      []byte("1.7"),
																		TC:          126,
																		StartLine:   5,
																		StartColumn: 87,
																		EndLine:     5,
																		EndColumn:   92,
																	},
																},
																1: {
																	Token: &lexmachine.Token{
																		Type:        nComment,
																		Value:       "2",
																		Lexeme:      []byte("2"),
																		TC:          137,
																		StartLine:   6,
																		StartColumn: 5,
																		EndLine:     6,
																		EndColumn:   8,
																	},
																},
																2: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "3.1",
																		Lexeme:      []byte("3.1"),
																		TC:          146,
																		StartLine:   7,
																		StartColumn: 5,
																		EndLine:     7,
																		EndColumn:   13,
																	},
																},
																3: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "3.2",
																		Lexeme:      []byte("3.2"),
																		TC:          158,
																		StartLine:   7,
																		StartColumn: 17,
																		EndLine:     7,
																		EndColumn:   25,
																	},
																},
																4: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "3.3",
																		Lexeme:      []byte("3.3"),
																		TC:          176,
																		StartLine:   7,
																		StartColumn: 35,
																		EndLine:     7,
																		EndColumn:   43,
																	},
																},
																5: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "3.4",
																		Lexeme:      []byte("3.4"),
																		TC:          188,
																		StartLine:   7,
																		StartColumn: 47,
																		EndLine:     7,
																		EndColumn:   55,
																	},
																},
																6: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "3.5",
																		Lexeme:      []byte("3.5"),
																		TC:          200,
																		StartLine:   7,
																		StartColumn: 59,
																		EndLine:     7,
																		EndColumn:   67,
																	},
																},
															},
														},
														Val: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val2",
																	Lexeme:      []byte("val2"),
																	TC:          210,
																	StartLine:   7,
																	StartColumn: 69,
																	EndLine:     7,
																	EndColumn:   72,
																},
															},
															Comments: map[uint64]*element{
																0: {
																	Token: &lexmachine.Token{
																		Type:        nCommentLong,
																		Value:       "3.6",
																		Lexeme:      []byte("3.6"),
																		TC:          215,
																		StartLine:   7,
																		StartColumn: 74,
																		EndLine:     7,
																		EndColumn:   82,
																	},
																},
															},
														},
													},
													{
														Square: false,
														Key: &exp{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "val3",
																	Lexeme:      []byte("val3"),
																	TC:          238,
																	StartLine:   8,
																	StartColumn: 5,
																	EndLine:     8,
																	EndColumn:   8,
																},
															},
															Comments: map[uint64]*element{
																0: {
																	Token: &lexmachine.Token{
																		Type:        nComment,
																		Value:       "3.7",
																		Lexeme:      []byte("3.7"),
																		TC:          227,
																		StartLine:   7,
																		StartColumn: 86,
																		EndLine:     7,
																		EndColumn:   91,
																	},
																},
															},
														},
													},
													{
														Square: false,
														Key: &exp{
															Comments: map[uint64]*element{
																0: {
																	Token: &lexmachine.Token{
																		Type:        nComment,
																		Value:       "4",
																		Lexeme:      []byte("4"),
																		TC:          244,
																		StartLine:   8,
																		StartColumn: 11,
																		EndLine:     8,
																		EndColumn:   14,
																	},
																},
																1: {
																	Token: &lexmachine.Token{
																		Type:        nComment,
																		Value:       "5",
																		Lexeme:      []byte("5"),
																		TC:          253,
																		StartLine:   9,
																		StartColumn: 5,
																		EndLine:     9,
																		EndColumn:   8,
																	},
																},
																2: {
																	Token: &lexmachine.Token{
																		Type:        nComment,
																		Value:       "6",
																		Lexeme:      []byte("6"),
																		TC:          262,
																		StartLine:   10,
																		StartColumn: 5,
																		EndLine:     10,
																		EndColumn:   8,
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
