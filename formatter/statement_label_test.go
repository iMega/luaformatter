package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseLabel(t *testing.T) {
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
			name: "label statement",
			args: args{
				code: []byte(`
:: label ::
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							Label: &labelStatement{
								Element: &element{
									Token: &lexmachine.Token{
										Type:        nID,
										Value:       "label",
										Lexeme:      []byte("label"),
										TC:          4,
										StartLine:   2,
										StartColumn: 4,
										EndLine:     2,
										EndColumn:   8,
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
			name: "label statement",
			args: args{
				code: []byte(`
::
      label
::
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							Label: &labelStatement{
								Element: &element{
									Token: &lexmachine.Token{
										Type:        nID,
										Value:       "label",
										Lexeme:      []byte("label"),
										TC:          10,
										StartLine:   3,
										StartColumn: 7,
										EndLine:     3,
										EndColumn:   11,
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
