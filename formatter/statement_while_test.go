package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timtadh/lexmachine"
)

func TestParseWhile(t *testing.T) {
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
			name: "while statement",
			args: args{
				code: []byte(`
while a do
    break
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
								While: &whileStatement{
									Exp: &exp{
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
									Do: &doStatement{
										Body: []Block{
											{
												Statement: statement{
													Break: &breakStatement{},
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
