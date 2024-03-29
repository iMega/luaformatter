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

	conf := DefaultConfig()
	conf.Highlight = false

	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "assignment statement",
			args: args{
				c: conf,
				b: []byte(`
function funccall() end
b = 1
x = 2
t = {
    call = function() end,
    ca = {
        ll = function() end,
    },
}
t0 = {}
t0[0] = {
    ca = {
        ll = function() end
    },
    call = function() end,
}
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
a = -1 and nil
a = -1 and false
a = -1 and true
a = -1 and ""
a = -1 and ...
a = -1 and function() end
a = -1 and funccall()
a = -1 and {}
a = -1 and 1 and -1
a = -1 or -1
a = -1 or -b
a = -1 or false
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
a = #t
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
a = t:call""
a = t:call()
a = t:call{}
a = t.call""
a = t.call()
a = t.call{}
a = t0[0].ca:ll""
a = t0[0].ca:ll()
a = t0[0].ca:ll{}
a = t0[0].call""
a = t0[0].call()
a = t0[0].call{}
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
local a = funccall()
local a = funccall({})
local a = funccall({
    c = d,
})
local a = t.call({
    d = e,
})
local a = funccall("1")
local a = funccall("")
c = -(x^2)
c = -x^2
c = (a < y) and (y <= z)
c = (a+i) < ((b/2)+1)
c = 5+((x^2)*8)
c = 5+x^2*8
c = a < y and y <= z
c = a+i < b/2+1
c = x^(y^z)
c = x^y^z
c = ((1 + 1) + 1) + 1
c = ((1 * 1) * 1) * 1
c = ((1 - 1) - 1) - 1
c = ((((((a((1 + 2)) + b()))))))
c = (a(((1+2)+1)) + b(1-2/2))
c = (a(((-1 + -2) + -1)) + b(-1 - -2 / -2))
`),
			},
			wantW: `
function funccall()
end

b = 1
x = 2
t = {
    call = function()
    end,
    ca = {
        ll = function()
        end,
    },
}
t0 = {}
t0[0] = {
    ca = {
        ll = function()
        end,
    },
    call = function()
    end,
}
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
a = -1 and nil
a = -1 and false
a = -1 and true
a = -1 and ""
a = -1 and ...
a = -1 and function()
end
a = -1 and funccall()
a = -1 and {}
a = -1 and 1 and -1
a = -1 or -1
a = -1 or -b
a = -1 or false
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
a = #t
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
a = t:call("")
a = t:call()
a = t:call({})
a = t.call("")
a = t.call()
a = t.call({})
a = t0[0].ca:ll("")
a = t0[0].ca:ll()
a = t0[0].ca:ll({})
a = t0[0].call("")
a = t0[0].call()
a = t0[0].call({})
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
local a = funccall()
local a = funccall({})
local a = funccall({c = d})
local a = t.call({d = e})
local a = funccall("1")
local a = funccall("")
c = -(x ^ 2)
c = -x ^ 2
c = (a < y) and (y <= z)
c = (a + i) < ((b / 2) + 1)
c = 5 + ((x ^ 2) * 8)
c = 5 + x ^ 2 * 8
c = a < y and y <= z
c = a + i < b / 2 + 1
c = x ^ (y ^ z)
c = x ^ y ^ z
c = ((1 + 1) + 1) + 1
c = ((1 * 1) * 1) * 1
c = ((1 - 1) - 1) - 1
c = ((((((a((1 + 2)) + b()))))))
c = (a(((1 + 2) + 1)) + b(1 - 2 / 2))
c = (a(((-1 + -2) + -1)) + b(-1 - -2 / -2))
`,
			wantErr: false,
		},
		{
			name: "function call statement",
			args: args{
				c: conf,
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
				c: conf,
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
				c: Config{
					Highlight:     false,
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
a = {}
a = {b = c}
a = {
    b = c,
    d = e,
}
local numbers = {1,2,3}
a = {
    b = 1,
    c = 2,
    d = 3,
}
a = {
    b = 1,
    c = 2,
    d = 3,
    e = 4,
}
local strings = {"012345", "012345", "012345", "012345", "012345"}
local strings = {"012345", "012345", "012345", "012345", "01234555555555555555"}
local strings = {"012345", "012345", "012345", "012345", "012345555555555555550"}
local strings = {"012345", "012345", "012345", "012345", "012345", "012345"}
`),
			},
			wantW: `
a = {}
a = {b = c}
a = {b = c, d = e}
local numbers = {1, 2, 3}
a = {b = 1, c = 2, d = 3}
a = {
    b = 1,
    c = 2,
    d = 3,
    e = 4,
}
local strings = {"012345", "012345", "012345", "012345", "012345"}
local strings = {"012345", "012345", "012345", "012345", "01234555555555555555"}
local strings = {
    "012345",
    "012345",
    "012345",
    "012345",
    "012345555555555555550",
}
local strings = {
    "012345",
    "012345",
    "012345",
    "012345",
    "012345",
    "012345",
}
`,
			wantErr: false,
		},
		{
			name: "break statement",
			args: args{
				c: conf,
				b: []byte(`
if true then break end
`),
			},
			wantW: `
if true then
    break
end
`,
			wantErr: false,
		},
		{
			name: "config for table alignment",
			args: args{
				c: Config{
					Highlight:     false,
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
table = {
    ["a()"] = false, -- comm 1
    [1+1] = true, -- comm 2
    bb = function () return 1 end, -- comm 3
    ["1394-E"] = val1, -- comm 4
    ["UTF-8"] = val2, -- comm 5
    ["and"] = val3, -- comm 6
    [true] = 1, -- comm 7
    aa = nil, -- comm 8
}
`),
			},
			wantW: `
table = {
    ["a()"] = false, -- comm 1
    [1 + 1] = true,  -- comm 2
    bb = function()
        return 1
    end, -- comm 3
    ["1394-E"] = val1, -- comm 4
    ["UTF-8"]  = val2, -- comm 5
    ["and"]    = val3, -- comm 6
    [true]     = 1,    -- comm 7
    aa         = nil,  -- comm 8
}
`,
			wantErr: false,
		},
		{
			name: "config for table alignment",
			args: args{
				c: Config{
					Highlight:     false,
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
table = {
    -- comm 0
    ["a()"] = false, -- comm 1
    [1+1] = true, -- comm 2
    bb = function () return 1 end, -- comm 3
    ["1394-E"] = val1, -- comm 4
    ["UTF-8"] = val2, -- comm 5
    ["and"] = val3, -- comm 6
    [true] = 1, -- comm 7
    aa = nil, -- comm 8
}
`),
			},
			wantW: `
table = {
    -- comm 0
    ["a()"] = false, -- comm 1
    [1 + 1] = true,  -- comm 2
    bb = function()
        return 1
    end, -- comm 3
    ["1394-E"] = val1, -- comm 4
    ["UTF-8"]  = val2, -- comm 5
    ["and"]    = val3, -- comm 6
    [true]     = 1,    -- comm 7
    aa         = nil,  -- comm 8
}
`,
			wantErr: false,
		},
		{
			name: "config for table alignment",
			args: args{
				c: Config{
					Highlight:     false,
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
table = {
    ["a()"] = false, -- comm 1
    -- comm 0
    [1+1] = true, -- comm 2
    bb = function () return 1 end, -- comm 3
    ["1394-E"] = val1, -- comm 4
    ["UTF-8"] = val2, -- comm 5
    ["and"] = val3, -- comm 6
    [true] = 1, -- comm 7
    aa = nil, -- comm 8
}
`),
			},
			wantW: `
table = {
    ["a()"] = false, -- comm 1
    -- comm 0
    [1 + 1] = true, -- comm 2
    bb = function()
        return 1
    end, -- comm 3
    ["1394-E"] = val1, -- comm 4
    ["UTF-8"]  = val2, -- comm 5
    ["and"]    = val3, -- comm 6
    [true]     = 1,    -- comm 7
    aa         = nil,  -- comm 8
}
`,
			wantErr: false,
		},
		{
			name: "config for table alignment",
			args: args{
				c: Config{
					Highlight:     false,
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
table = {
    -- comm a
    -- comm b
    ["a()"] = false, -- comm 1
    -- comm c
    -- comm d
    [1+1] = true, -- comm 2
    -- comm e
    -- comm f
    bb = function () return 1 end, -- comm 3
    -- comm g
    -- comm h
    ["1394-E"] = val1, -- comm 4
    ["UTF-8"] = val2, -- comm 5
    ["and"] = val3, -- comm 6
    [true] = 1, -- comm 7
    aa = nil, -- comm 8
    -- comm i
    -- comm j
}
`),
			},
			wantW: `
table = {
    -- comm a
    -- comm b
    ["a()"] = false, -- comm 1
    -- comm c
    -- comm d
    [1 + 1] = true, -- comm 2
    -- comm e
    -- comm f
    bb = function()
        return 1
    end, -- comm 3
    -- comm g
    -- comm h
    ["1394-E"] = val1, -- comm 4
    ["UTF-8"]  = val2, -- comm 5
    ["and"]    = val3, -- comm 6
    [true]     = 1,    -- comm 7
    aa         = nil,  -- comm 8
    -- comm i
    -- comm j
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
