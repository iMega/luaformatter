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
				Body: &body{
					Blocks: map[uint64]block{
						0: {
							Statement: statement{
								Label: &labelStatement{
									Element: &element{
										Token: &lexmachine.Token{
											Type:        nLabel,
											Value:       "label",
											Lexeme:      []byte("label"),
											TC:          1,
											StartLine:   2,
											StartColumn: 1,
											EndLine:     2,
											EndColumn:   11,
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
			name: "label statement with newline",
			args: args{
				code: []byte(`
::
      label
::
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]block{
						0: {
							Statement: statement{
								Label: &labelStatement{
									Element: &element{
										Token: &lexmachine.Token{
											Type:        nLabel,
											Value:       "label",
											Lexeme:      []byte("label"),
											TC:          1,
											StartLine:   2,
											StartColumn: 1,
											EndLine:     4,
											EndColumn:   2,
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
			name: "two label statement",
			args: args{
				code: []byte(`
::
      label
::

::
      label2
::
`,
				),
			},
			want: &document{
				Body: &body{
					Blocks: map[uint64]block{
						0: {
							Statement: statement{
								Label: &labelStatement{
									Element: &element{
										Token: &lexmachine.Token{
											Type:        nLabel,
											Value:       "label",
											Lexeme:      []byte("label"),
											TC:          1,
											StartLine:   2,
											StartColumn: 1,
											EndLine:     4,
											EndColumn:   2,
										},
									},
								},
							},
						},
						1: {
							Statement: statement{
								NewLine: &newlineStatement{},
							},
						},
						2: {
							Statement: statement{
								Label: &labelStatement{
									Element: &element{
										Token: &lexmachine.Token{
											Type:        nLabel,
											Value:       "label2",
											Lexeme:      []byte("label2"),
											TC:          20,
											StartLine:   6,
											StartColumn: 1,
											EndLine:     8,
											EndColumn:   2,
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
