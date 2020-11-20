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
								ForNumerical: &forNumericalStatement{
									VarPart: &field{
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
									LimitPart: &exp{
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
									StepPart: &exp{
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
