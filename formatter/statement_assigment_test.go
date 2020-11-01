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
			name: "local assignment statement with one var",
			args: args{
				code: []byte(`
local a
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
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
				QtyBlocks: 1,
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
				Body: map[uint64]Block{
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
				QtyBlocks: 1,
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
				Body: map[uint64]Block{
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
				QtyBlocks: 1,
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
				Body: map[uint64]Block{
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
