package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseFunccall(t *testing.T) {
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
			name: "funccall statement",
			args: args{
				code: []byte(`
funccall()
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
										Value:       "funccall",
										Lexeme:      []byte("funccall"),
										TC:          1,
										StartLine:   2,
										StartColumn: 1,
										EndLine:     2,
										EndColumn:   8,
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
			name: "funccall statement with string literal",
			args: args{
				code: []byte(`
funccall "literal"
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
										Value:       "funccall",
										Lexeme:      []byte("funccall"),
										TC:          1,
										StartLine:   2,
										StartColumn: 1,
										EndLine:     2,
										EndColumn:   8,
									},
								},
							},
							Explist: &explist{
								List: []*exp{
									{
										Element: &element{
											Token: &lexmachine.Token{
												Type:        nString,
												Value:       `"literal"`,
												Lexeme:      []byte(`"literal"`),
												TC:          10,
												StartLine:   2,
												StartColumn: 10,
												EndLine:     2,
												EndColumn:   18,
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
