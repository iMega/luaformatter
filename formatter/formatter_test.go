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
a = b
a = ""
a = 0
local a
local a, b
local a = b
local a, b = c, d
local a = b()
local a = b({})
local a = b({
    c = d,
})
local a = b.c({
    d = e,
})
local a = b("1")
local a = b("")
`),
			},
			wantW: `
a = b
a = ""
a = 0
local a
local a, b
local a = b
local a, b = c, d
local a = b()
local a = b({})
local a = b({
    c = d,
})
local a = b.c({
    d = e,
})
local a = b("1")
local a = b("")
`,
			wantErr: false,
		},
		{
			name: "function call statement",
			args: args{
				c: DefaultConfig(),
				b: []byte(`
a()
a{}
a({})
a""
`),
			},
			wantW: `
a()
a({})
a({})
a("")
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
