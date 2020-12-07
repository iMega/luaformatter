package formatter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
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
a = -1 - -1
a = -1 - -b
a = -1 .. -1
a = -1 .. -b
a = -1 * -1
a = -1 * -b
a = -1 / -1
a = -1 / -b
a = -1 & -1
a = -1 & -b
a = -1 % -1
a = -1 % -b
a = -1 ^ -1
a = -1 ^ -b
a = -1 + -1
a = -1 + -b
a = -1 < -1
a = -1 < -b
a = -1 << -1
a = -1 << -b
a = -1 <= -1
a = -1 <= -b
a = -1 == -1
a = -1 == -b
a = -1 > -1
a = -1 > -b
a = -1 >= -1
a = -1 >= -b
a = -1 >> -1
a = -1 >> -b
a = -1 | -1
a = -1 | -b
a = -1 ~ -1
a = -1 ~ -b
a = -1 ~= -1
a = -1 ~= -b
a = -1 and -1
a = -1 and -b
a = -1 or -1
a = -1 or -b
a = -b
a = -b - -1
a = -b .. -1
a = -b * -1
a = -b / -1
a = -b & -1
a = -b % -1
a = -b ^ -1
a = -b + -1
a = -b < -1
a = -b << -1
a = -b <= -1
a = -b == -1
a = -b > -1
a = -b >= -1
a = -b >> -1
a = -b | -1
a = -b ~ -1
a = -b ~= -1
a = -b and -1
a = -b or -1
a = ...
a = ... == ...
a = ... ~= ...
a = ... and ...
a = ... or ...
a = ""
a = "" .. ""
a = "" < ""
a = "" <= ""
a = "" == ""
a = "" > ""
a = "" >= ""
a = "" ~= ""
a = "" and ""
a = "" or ""
a = {}
a = #b
a = 0
a = 1 - 1
a = 1 - b
a = 1 .. 1
a = 1 .. b
a = 1 * 1
a = 1 * b
a = 1 / 1
a = 1 / b
a = 1 & 1
a = 1 & b
a = 1 % 1
a = 1 % b
a = 1 ^ 1
a = 1 ^ b
a = 1 + 1
a = 1 + b
a = 1 < 1
a = 1 < b
a = 1 << 1
a = 1 << b
a = 1 <= 1
a = 1 <= b
a = 1 == 1
a = 1 == b
a = 1 > 1
a = 1 > b
a = 1 >= 1
a = 1 >= b
a = 1 >> 1
a = 1 >> b
a = 1 | 1
a = 1 | b
a = 1 ~ 1
a = 1 ~ b
a = 1 ~= 1
a = 1 ~= b
a = 1 and 1
a = 1 and b
a = 1 or 1
a = 1 or b
a = b
a = false
a = false == nil
a = false == true
a = false ~= nil
a = false ~= true
a = false and nil
a = false and true
a = false or nil
a = false or true
a = func:call""
a = func:call()
a = func:call{}
a = func.call""
a = func.call()
a = func.call{}
a = func[0].ca:ll""
a = func[0].ca:ll()
a = func[0].ca:ll{}
a = func[0].call""
a = func[0].call()
a = func[0].call{}
a = funccall""
a = funccall()
a = funccall{}
a = nil
a = nil == true
a = nil ~= true
a = nil and true
a = nil or true
a = true
a = true == false
a = true ~= false
a = true and false
a = true or false
a = function()
end
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
c = -(x^2)
`),
			},
			wantW: `
a = -1
a = -1 - -1
a = -1 - -b
a = -1 .. -1
a = -1 .. -b
a = -1 * -1
a = -1 * -b
a = -1 / -1
a = -1 / -b
a = -1 & -1
a = -1 & -b
a = -1 % -1
a = -1 % -b
a = -1 ^ -1
a = -1 ^ -b
a = -1 + -1
a = -1 + -b
a = -1 < -1
a = -1 < -b
a = -1 << -1
a = -1 << -b
a = -1 <= -1
a = -1 <= -b
a = -1 == -1
a = -1 == -b
a = -1 > -1
a = -1 > -b
a = -1 >= -1
a = -1 >= -b
a = -1 >> -1
a = -1 >> -b
a = -1 | -1
a = -1 | -b
a = -1 ~ -1
a = -1 ~ -b
a = -1 ~= -1
a = -1 ~= -b
a = -1 and -1
a = -1 and -b
a = -1 or -1
a = -1 or -b
a = -b
a = -b - -1
a = -b .. -1
a = -b * -1
a = -b / -1
a = -b & -1
a = -b % -1
a = -b ^ -1
a = -b + -1
a = -b < -1
a = -b << -1
a = -b <= -1
a = -b == -1
a = -b > -1
a = -b >= -1
a = -b >> -1
a = -b | -1
a = -b ~ -1
a = -b ~= -1
a = -b and -1
a = -b or -1
a = ...
a = ... == ...
a = ... ~= ...
a = ... and ...
a = ... or ...
a = ""
a = "" .. ""
a = "" < ""
a = "" <= ""
a = "" == ""
a = "" > ""
a = "" >= ""
a = "" ~= ""
a = "" and ""
a = "" or ""
a = {}
a = #b
a = 0
a = 1 - 1
a = 1 - b
a = 1 .. 1
a = 1 .. b
a = 1 * 1
a = 1 * b
a = 1 / 1
a = 1 / b
a = 1 & 1
a = 1 & b
a = 1 % 1
a = 1 % b
a = 1 ^ 1
a = 1 ^ b
a = 1 + 1
a = 1 + b
a = 1 < 1
a = 1 < b
a = 1 << 1
a = 1 << b
a = 1 <= 1
a = 1 <= b
a = 1 == 1
a = 1 == b
a = 1 > 1
a = 1 > b
a = 1 >= 1
a = 1 >= b
a = 1 >> 1
a = 1 >> b
a = 1 | 1
a = 1 | b
a = 1 ~ 1
a = 1 ~ b
a = 1 ~= 1
a = 1 ~= b
a = 1 and 1
a = 1 and b
a = 1 or 1
a = 1 or b
a = b
a = false
a = false == nil
a = false == true
a = false ~= nil
a = false ~= true
a = false and nil
a = false and true
a = false or nil
a = false or true
a = func:call("")
a = func:call()
a = func:call({})
a = func.call("")
a = func.call()
a = func.call({})
a = func[0].ca:ll("")
a = func[0].ca:ll()
a = func[0].ca:ll({})
a = func[0].call("")
a = func[0].call()
a = func[0].call({})
a = funccall("")
a = funccall()
a = funccall({})
a = nil
a = nil == true
a = nil ~= true
a = nil and true
a = nil or true
a = true
a = true == false
a = true ~= false
a = true and false
a = true or false
a = function()
end
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
c = -(x ^ 2)
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

			if !assert.Equal(t, tt.wantW, w.String()) {
				t.Error("failed to format")
			}
		})
	}
}
