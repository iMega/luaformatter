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
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							Comment: &commentStatement{
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
					},
				},
				QtyBlocks: 1,
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
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							Comment: &commentStatement{
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
					},
					1: {
						Statement: statement{
							Comment: &commentStatement{
								Element: &element{
									Token: &lexmachine.Token{
										Type:        nComment,
										Value:       "comment2",
										Lexeme:      []byte("comment2"),
										TC:          12,
										StartLine:   3,
										StartColumn: 1,
										EndLine:     3,
										EndColumn:   15,
									},
								},
							},
						},
					},
				},
				QtyBlocks: 2,
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
