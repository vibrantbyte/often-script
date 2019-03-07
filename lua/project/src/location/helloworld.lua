
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
