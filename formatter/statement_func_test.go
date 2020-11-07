package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseFunction(t *testing.T) {
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
			name: "empty function",
			args: args{
				code: []byte(`
function a()
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
								Function: &functionStatement{
									Name: &element{
										Token: &lexmachine.Token{
											Type:        nID,
											Value:       "a",
											Lexeme:      []byte("a"),
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
					Qty: 1,
				},
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "empty two function",
			args: args{
				code: []byte(`
function a()
end

function b()
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
								Function: &functionStatement{
									Name: &element{
										Token: &lexmachine.Token{
											Type:        nID,
											Value:       "a",
											Lexeme:      []byte("a"),
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
						1: {Statement: statement{
							NewLine: &newlineStatement{},
						}},
						2: {
							Statement: statement{
								Function: &functionStatement{
									Name: &element{
										Token: &lexmachine.Token{
											Type:        nID,
											Value:       "b",
											Lexeme:      []byte("b"),
											TC:          28,
											StartLine:   5,
											StartColumn: 10,
											EndLine:     5,
											EndColumn:   10,
										},
									},
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
			name: "function with arguments",
			args: args{
				code: []byte(`
function sum(a, b)
    return a + b
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
								Function: &functionStatement{
									Name: &element{
										Token: &lexmachine.Token{
											Type:        nID,
											Value:       "sum",
											Lexeme:      []byte("sum"),
											TC:          10,
											StartLine:   2,
											StartColumn: 10,
											EndLine:     2,
											EndColumn:   12,
										},
									},
									Parlist: &explist{
										List: []*exp{
											{
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       "a",
														Lexeme:      []byte("a"),
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
														Value:       "b",
														Lexeme:      []byte("b"),
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
									Body: []Block{
										{
											Return: &returnStatement{
												Explist: &explist{
													List: []*exp{
														{
															Element: &element{
																Token: &lexmachine.Token{
																	Type:        nID,
																	Value:       "a",
																	Lexeme:      []byte("a"),
																	TC:          31,
																	StartLine:   3,
																	StartColumn: 12,
																	EndLine:     3,
																	EndColumn:   12,
																},
															},
															Binop: &element{
																Token: &lexmachine.Token{
																	Type:        nAddition,
																	Value:       "+",
																	Lexeme:      []byte("+"),
																	TC:          33,
																	StartLine:   3,
																	StartColumn: 14,
																	EndLine:     3,
																	EndColumn:   14,
																},
															},
															Exp: &exp{
																Element: &element{
																	Token: &lexmachine.Token{
																		Type:        nID,
																		Value:       "b",
																		Lexeme:      []byte("b"),
																		TC:          35,
																		StartLine:   3,
																		StartColumn: 16,
																		EndLine:     3,
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
