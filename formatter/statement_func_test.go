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
		name    string
		args    args
		want    *document
		wantErr bool
	}{
		{
			name: "empty function",
			args: args{
				code: []byte(`
function a()
end
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							Function: &functionStatement{
								NamePart: &element{
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
				QtyBlocks: 1,
			},
			wantErr: false,
		},
		{
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
				Body: map[uint64]Block{
					0: {
						Statement: statement{
							Function: &functionStatement{
								NamePart: &element{
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
					1: {
						Statement: statement{
							Function: &functionStatement{
								NamePart: &element{
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
				QtyBlocks: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
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
