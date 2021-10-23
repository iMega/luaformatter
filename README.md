# Lua formatter

[![Test](https://github.com/iMega/luaformatter/actions/workflows/test.yml/badge.svg)](https://github.com/iMega/luaformatter/actions/workflows/test.yml) [![codecov](https://codecov.io/gh/iMega/luaformatter/branch/master/graph/badge.svg?token=O1619F5DYR)](https://codecov.io/gh/iMega/luaformatter) [![Go Report Card](https://goreportcard.com/badge/github.com/imega/luaformatter)](https://goreportcard.com/report/github.com/imega/luaformatter) [![codebeat badge](https://codebeat.co/badges/27b31bf7-fe96-421f-a1dd-0d3fe7613c60)](https://codebeat.co/projects/github-com-imega-luaformatter-master)

## Code formatting rules

### Variables

The option alignment.variables will align list of variables.

```lua
local a  = 1  -- comment
local bb = 22 -- comment

local c = "short text" -- comment
```

### Comments

Skipped space between a double hyphen and text will be revert. Semantic line
(only hyphens) will not have the space.

```lua
--------------------------------
-- Description
-- params:  - input
--          - output
```

### Tables

If you create an empty table or table with only key fields or
table with one field they will be written in one line.

```lua
local t = {}
local t = {"one", "two", 3, 4, 5}
local t = {a = 1, b = 2, c = 3}
```

If you write more than five fields with only key then table will be written in
multiline. Same behaviour will be if you write more than four fields
with key and value. If the table is multiline then every field will have
trailing comma.

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

The option alignment.table will align table.

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

If return statement is inside the block with other statements then a newline
will be added before it.

```lua
function isFalse()
    return false
end

function calc(a, b)
    local c = a + b

    return c
end
```

If a function is inside the block with other statements then a newline
will be added before it.

```lua
function isFalse()
    return false
end

isFalse()
```

If a function has parameters more than max line length then
every parameter will be moved to the next line.

```lua
function very_long_name(
    very_long_name_parameter_1,
    very_long_name_parameter_2,
    very_long_name_parameter_3
)
    -- stuff
end
```

If a function call has parameters more than max line length then
every parameter will be moved to the next line.

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
