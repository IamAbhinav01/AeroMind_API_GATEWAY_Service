const { Redis } = require('ioredis');
const { REDIS_HOST, REDIS_PORT, REDIS_TIMEOUT } = require('./server.config');

const RedisConnection = new Redis({
    host: REDIS_HOST,
    port: REDIS_PORT,
    connectTimeout: REDIS_TIMEOUT,
});

RedisConnection.on('connect', () => {
    console.log('[Redis] Connected successfully');
});

RedisConnection.on('error', (err) => {
    console.error('[Redis] Connection error:', err.message);
});

module.exports = RedisConnection;
