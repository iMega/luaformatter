local t = {a = 1}

if 1 > 0 then
    box.cfg {
        feedback_enabled = false,
        log_format = 'json',
        log_level = 5,
        -- memtx_memory = 1.4 * 1024 * 1024 * 1024,
    }
end

for i = 1, i < 10, 3 do
    --
end

qwerty({
    feedback_enabled = false,
    log_format = 'json',
    memtx_memory = 1.4 * 1024 * 1024 * 1024,
}, qweeryy(), qweeryy(), qweeryy(), qweeryy(), qweeryy(), qweeryy(), qweeryy(),
       qweeryy()).qwerty{
    feedback_enabled = false,
    log_format = 'json',
    memtx_memory = 1.4 * 1024 * 1024 * 1024,
}.qwerty{
    feedback_enabled = false,
    log_format = 'json',
    memtx_memory = 1.4 * 1024 * 1024 * 1024,
}.qwerty{
    feedback_enabled = false,
    log_format = 'json',
    memtx_memory = 1.4 * 1024 * 1024 * 1024,
}.qwerty {
    feedback_enabled = false,
    log_format = 'json',
    memtx_memory = 1.4 * 1024 * 1024 * 1024,
}

function aaa()
    return function()
        local a = 1
        return 22 + 1
    end, 22, 123123123, "dsfsdfsfgdfgdfg",
           {a = 1, sdfsdf = 222234, sdfsdf = 2434234}, function()
        local a = 1
        return 22 + 1
    end, function()
        local a = 1
        return 22 + 1
    end
end

goto Name

for i = 10, 100, 2 do
    ::Name::
    print(i)
end

local a = {["a"] = 111, [function() return 1 end] = 22, [aaa[a]] = 1}
local b = {false + false}

local fn = {
    [function(a, b) return a + b end] = 1,
    [function(a, b) return a - b end] = 1,
}

for k, v in pairs(fn) do f(v) end

