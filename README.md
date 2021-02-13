# Lua formatter

[![Build Status](https://travis-ci.org/iMega/luaformatter.svg?branch=master)](https://travis-ci.org/iMega/luaformatter) [![codecov](https://codecov.io/gh/iMega/luaformatter/branch/master/graph/badge.svg?token=O1619F5DYR)](https://codecov.io/gh/iMega/luaformatter) [![Go Report Card](https://goreportcard.com/badge/github.com/imega/luaformatter)](https://goreportcard.com/report/github.com/imega/luaformatter) [![codebeat badge](https://codebeat.co/badges/27b31bf7-fe96-421f-a1dd-0d3fe7613c60)](https://codebeat.co/projects/github-com-imega-luaformatter-master)

## Code formatting rules

### Variables

The option alignment.variables will be align list of variables.

```lua
local a  = 1  -- comment
local bb = 22 -- comment

local c = "short text" -- comment
```

### Comments

Skipped space between a double hyphen and text will be revert. Semantic line
(only hyphens) will not has the space.

```lua
--------------------------------
-- Description
-- params:  - input
--          - output
```

### Tables

If you create an empty table or with only field key or
with one field then will be written in one line.

```lua
local t = {}
local t = {"one", "two", 3, 4, 5}
local t = {a = 1, b = 2, c = 3}
```

If you write more five fields with only key then table will be written in
multiline. Same behaviour will if more four fields with key and value.
If the table is multiline then every field will has trailing comma.

```lua
local t = {
    "one",
    "two",
    3,
    4,
    5,
    "six",
}
local t = {
    a = 1,
    b = 2,
    c = 3,
    d = 4,
}
```

The option alignment.table will be align table.

```lua
local t = {
    a      = 1,    -- comment
    ["bb"] = "bb", -- comment
    -- comment
    ccc    = "ccc",    -- comment
    dddeee = "dddeee", -- comment
    f = function()
        return 1
    end,
    g = 7, -- comment
    -- comment
}
```

### Newline

If return statement inside block with other statements.

```lua
function isFalse()
    return false
end

function calc(a, b)
    local c = a + b

    return c
end
```

A function with other statements.

```lua
function isFalse()
    return false
end

isFalse()
```

If a function has parameters more max line length.

```lua
function very_long_name(
    very_long_name_parameter_1,
    very_long_name_parameter_2,
    very_long_name_parameter_3
)
    -- stuff
end
```

If a function call has parameters more max line length.

```lua
very_long_name_function_call(
    {a = 1, b = 2, c = 3},
    very_long_name_parameter,
    {
        a = 1,
        b = 2,
        c = 3,
        d = 4,
    }
)
```
