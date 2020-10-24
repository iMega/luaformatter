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
			name: "prefixexp statement conver to function call",
			args: args{
				//a.b.c.v:r["ff"].c["bb"]("vvv").cc().dd
				code: []byte(`
r["ff"]["dd"].name.name2["ee"]()
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							FuncCall: &funcCallStatement{
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
												Element: &element{
													Token: &lexmachine.Token{
														Type:        nID,
														Value:       ".name.name2",
														Lexeme:      []byte(".name.name2"),
														TC:          14,
														StartLine:   2,
														StartColumn: 14,
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
								Explist: &explist{
									List: []*exp{
										{},
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
