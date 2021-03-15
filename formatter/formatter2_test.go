package formatter

import (
	"os"
	"testing"
)

func TestFormat2(t *testing.T) {
	type args struct {
		c Config
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "config for table alignment",
			args: args{
				c: Config{
					IndentSize:    4,
					MaxLineLength: 80,
					Alignment: Alignment{
						Table: AlignmentTable{
							KeyValuePairs: true,
							Comments:      true,
						},
					},
				},
				b: []byte(`
return (machine and machine.is_loaded_qwe) and "coffee brewing" or "fill your water"
`),
			},
			wantW: `
a.b().c().d()
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// w := &bytes.Buffer{}
			w := os.Stdout
			w.Write([]byte("\n"))
			if err := Format(tt.args.c, tt.args.b, w); (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// if !assert.Equal(t, tt.wantW, w.String()) {
			// 	t.Error("failed to format")
			// }
		})
	}
}
