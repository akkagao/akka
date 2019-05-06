return redis.call('GET', KEYS[1])

-- redis-cli --eval get.lua 34234
-- redis-cli -h 172.168.18.310 -p 6378 -a password  --eval get.lua aa