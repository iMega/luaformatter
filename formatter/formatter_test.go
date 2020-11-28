package formatter

import (
	"bytes"
	"testing"
)

func TestFormat(t *testing.T) {
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
			name: "assignment statement",
			args: args{
				c: DefaultConfig(),
				b: []byte(`
a = 1
local a
local a = b
`),
			},
			wantW: `
a = 1
local a
local a = b
`,
			wantErr: false,
		},
		{
			name: "function call statement with empty table",
			args: args{
				c: DefaultConfig(),
				b: []byte(`
a{}
a({})
b""
`),
			},
			wantW: `
a({})
a({})
b("")
`,
			wantErr: false,
		},
		{
			name: "function call statement with table",
			args: args{
				c: DefaultConfig(),
				b: []byte("a{b=1}"),
			},
			wantW: `
a({
    b = 1,
})
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			w.Write([]byte("\n"))
			if err := Format(tt.args.c, tt.args.b, w); (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Format() = \n%v, want \n%v", gotW, tt.wantW)
			}
		})
	}
}
