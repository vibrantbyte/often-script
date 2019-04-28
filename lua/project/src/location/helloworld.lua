
--- 启动调
local mobdebug = require("src.initial.mobdebug");
mobdebug.start();

local print_log = require("src.server.plugin_print");

local function get_name()
    local first_name = "a";
    local second_name = "b";
    return first_name..second_name;
end

ngx.say(get_name()..print_log.path);


local incr_param = ARGV[1];
local result = 1024 * 8;
if (incr_param == '1') then
    result = redis.call('INCR',KEYS[1]);
elseif (incr_param == '-1') then
    result = redis.call('DECR',KEYS[1]);
end

if(redis.call('TTL',KEYS[1]) < 0) then
    redis.call('EXPIRE',KEYS[1],ARGV[1])
end

return result;







local incr_param = ARGV[1];local result = 1024 * 8;if (incr_param == '1') then result = redis.call('INCR',KEYS[1]); else result = redis.call('DECR',KEYS[1]); end if( redis.call('TTL',KEYS[1]) < 0) then redis.call('EXPIRE',KEYS[1],ARGV[2]); end return result;
