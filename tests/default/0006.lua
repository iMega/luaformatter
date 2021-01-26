 local function brew_coffee(machine)
    return (machine and machine.is_loaded) and "coffee brewing" or "fill your water"
 end

 -- good
if test then break end

-- good
if not ok then return nil, "this failed for this reason: " .. reason end

-- good
use_callback(x, function(k) return k.last end)

-- good
if test then
  return false
end

-- bad
if test < 1 and do_complicated_function(test) == false or seven == 8 and nine == 10 then do_other_complicated_function() end

-- good
if test < 1 and do_complicated_function(test) == false or seven == 8 and nine == 10 then
   do_other_complicated_function()
   return false
end

-- bad
local whatever = "sure";
a = 1; b = 2

-- good
local whatever = "sure"
a = 1
b = 2

--bad
-- good

-- bad
local x = y*9
local numbers={1,2,3}
numbers={1 , 2 , 3}
numbers={1 ,2 ,3}
local strings = { "hello"
                , "Lua"
                , "world"
                }
dog.set( "attr",{
  age="1 year",
  breed="Bernese Mountain Dog"
})

-- good
local x = y * 9
local numbers = {1, 2, 3}
local strings = {
   "hello",
   "Lua",
   "world",
}
dog.set("attr", {
   age = "1 year",
   breed = "Bernese Mountain Dog",
})
