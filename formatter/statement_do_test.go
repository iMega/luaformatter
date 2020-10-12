package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDo(t *testing.T) {
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
			name: "do statement",
			args: args{
				code: []byte(`
do
    break
end
`,
				),
			},
			want: &document{
				Body: map[uint64]Block{
					0: {
						Statement: statement{
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
