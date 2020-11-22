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
    name = "Jack",
    ["1394-E"] = val1,
    ["UTF-8"] = val2,
    ["and"] = val2,
}
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
