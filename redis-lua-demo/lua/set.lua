return redis.call('set', KEYS[1], ARGV[1])

-- redis-cli --eval set.lua aa ,aa
