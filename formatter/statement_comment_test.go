package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseComment(t *testing.T) {
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
			name: "comment statement",
			args: args{
				code: []byte(`
-- comment
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]statement{
						0: &commentStatement{
							Element: &element{
								Token: &lexmachine.Token{
									Type:        nComment,
									Value:       "comment",
									Lexeme:      []byte("comment"),
									TC:          1,
									StartLine:   2,
									StartColumn: 1,
									EndLine:     2,
									EndColumn:   10,
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
			name: "two comment statement",
			args: args{
				code: []byte(`
-- comment
--     comment2
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]statement{
						0: &commentStatement{
							Element: &element{
								Token: &lexmachine.Token{
									Type:        nComment,
									Value:       "comment",
									Lexeme:      []byte("comment"),
									TC:          1,
									StartLine:   2,
									StartColumn: 1,
									EndLine:     2,
									EndColumn:   10,
								},
							},
						},
						1: &commentStatement{
							Element: &element{
								Token: &lexmachine.Token{
									Type:        nComment,
									Value:       "    comment2",
									Lexeme:      []byte("    comment2"),
									TC:          12,
									StartLine:   3,
									StartColumn: 1,
									EndLine:     3,
									EndColumn:   15,
								},
							},
							IsNewline: true,
						},
					},
					Qty: 2,
				},
			},
			wantErr: false,
		},
		{
			skip: false,
			name: "two comment statement",
			args: args{
				code: []byte(`
---------------
-- comment
--     comment2
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]statement{
						0: &commentStatement{
							Element: &element{
								Token: &lexmachine.Token{
									Type:        nComment,
									Value:       "-------------",
									Lexeme:      []byte("-------------"),
									TC:          1,
									StartLine:   2,
									StartColumn: 1,
									EndLine:     2,
									EndColumn:   15,
								},
							},
						},
						1: &commentStatement{
							Element: &element{
								Token: &lexmachine.Token{
									Type:        nComment,
									Value:       "comment",
									Lexeme:      []byte("comment"),
									TC:          17,
									StartLine:   3,
									StartColumn: 1,
									EndLine:     3,
									EndColumn:   10,
								},
							},
							IsNewline: true,
						},
						2: &commentStatement{
							Element: &element{
								Token: &lexmachine.Token{
									Type:        nComment,
									Value:       "    comment2",
									Lexeme:      []byte("    comment2"),
									TC:          28,
									StartLine:   4,
									StartColumn: 1,
									EndLine:     4,
									EndColumn:   15,
								},
							},
							IsNewline: true,
						},
					},
					Qty: 3,
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
