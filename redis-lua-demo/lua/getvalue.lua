local exists = redis.call('exists', KEYS[1])
if exists == 1 then
    return redis.call('GET', KEYS[1])
else
    redis.call('SET', KEYS[1], ARGV[1])
    return ARGV[1]
end
-- statements



-- 命令行启动脚本
-- redis-cli --eval getvalue.lua bb , 234555
