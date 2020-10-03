package main

import (
	"bytes"
	"log"

	"github.com/imega/luaformatter/formatter"
)

var code = `
local box='a'

local box2="a2"

local box3=[[
	qweqwe qweqwe
]]

local log=require('log')

box.cfg{
    feedback_enabled = false,
    log_format = 'json',
    log_level = 5,
    -- memtx_memory = 1.4 * 1024 * 1024 * 1024,
}

box.cfg(
	{
		feedback_enabled = false,
		log_format = 'json',
		log_level = 5,
		memtx_memory = 1.4 * 1024 * 1024 * 1024,
	},
	werwer,
	{
		feedback_enabled = false,
		log_format = 'json',
		log_level = 5,
		memtx_memory = 1.4 * 1024 * 1024 * 1024,
	},
)
`

func main() {
	// buf := bytes.NewBuffer([]byte(`local box = 'a'`))

	w := &bytes.Buffer{}
	if err := formatter.Format([]byte(code), w); err != nil {
		log.Fatalf("failed to format code, %s", err)
	}
}
