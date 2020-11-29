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
a = -1
a = ""
a = "" .. ""
a = {}
a = #b
a = 0
a = 1 - 1
a = 1 .. 1
a = 1 * 1
a = 1 / 1
a = 1 & 1
a = 1 % 1
a = 1 ^ 1
a = 1 + 1
a = 1 < 1
a = 1 << 1
a = 1 <= 1
a = 1 == 1
a = 1 > 1
a = 1 >= 1
a = 1 >> 1
a = 1 | 1
a = 1 ~ 1
a = 1 ~= 1
a = 1 and 1
a = 1 or 1
a = b
a = false
a = function()
end
a = nil
a = true
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
a = -1
a = ""
a = "" .. ""
a = {}
a = #b
a = 0
a = 1 - 1
a = 1 .. 1
a = 1 * 1
a = 1 / 1
a = 1 & 1
a = 1 % 1
a = 1 ^ 1
a = 1 + 1
a = 1 < 1
a = 1 << 1
a = 1 <= 1
a = 1 == 1
a = 1 > 1
a = 1 >= 1
a = 1 >> 1
a = 1 | 1
a = 1 ~ 1
a = 1 ~= 1
a = 1 and 1
a = 1 or 1
a = b
a = false
a = function()
end
a = nil
a = true
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
		{
			name: "for statement",
			args: args{
				c: DefaultConfig(),
				b: []byte(`
for a in b do
end
for a in b do
    -- comment
end
for a, b in c do
end
for a, b in c(d) do
end
for a = 1, 1 do
end
for a = 1, 1 do
    -- comment
end
for a = 1, 1, -1 do
end
for a = 1, 1, -1 do
    -- comment
end
`),
			},
			wantW: `
for a in b do
end
for a in b do
    -- comment
end
for a, b in c do
end
for a, b in c(d) do
end
for a = 1, 1 do
end
for a = 1, 1 do
    -- comment
end
for a = 1, 1, -1 do
end
for a = 1, 1, -1 do
    -- comment
end
`,
			wantErr: false,
		},
		{
			name: "table statement",
			args: args{
				c: DefaultConfig(),
				b: []byte(`
a = {}
a = {
    b = c
}
a = {
    b = c,
}
`),
			},
			wantW: `
a = {}
a = {
    b = c,
}
a = {
    b = c,
}
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
