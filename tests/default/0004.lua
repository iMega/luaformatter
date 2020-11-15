local a = 1
local b, c
if a == 1 then
    b = function()
        return 1
    end
elseif a == 2 then
    b = function()
        return 2
    end
elseif a == 3 then
    b = function()
        return 3
    end
else
    b = function()
        return 4
    end
end
c = b()
print(c)
