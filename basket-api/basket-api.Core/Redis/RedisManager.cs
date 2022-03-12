using Microsoft.Extensions.Configuration;
using Newtonsoft.Json;
using System;
using System.Threading.Tasks;

namespace basket_api.Core.Redis
{
    public class RedisManager : IRedisManager
    {
        private readonly IConfiguration _configuration;
        private RedisServer _redisServer;

        public RedisManager(IConfiguration configuration)
        {
            _configuration = configuration;
            _redisServer = new RedisServer(_configuration);
        }

        public async Task<bool> Exists(string key)
        {
            return await _redisServer.Database.KeyExistsAsync(key);
        }

        public async Task<object> Get(string key)
        {
            bool result = await Exists(key);
            if (result)
            {
                string jsonData = _redisServer.Database.StringGet(key);
                return jsonData == null ? null : jsonData;
            }
            return null;
        }

        public async Task Remove(string key)
        {
            await _redisServer.Database.KeyDeleteAsync(key);
        }

        public void RemoveAll()
        {
            _redisServer.FlushDatabase();
        }

        public async Task Set(string key, object value, DateTimeOffset absoluteExpiration, TimeSpan? slidingExpiration, bool removable)
        {
            if (slidingExpiration != null)
                await _redisServer.Database.KeyExpireAsync(key, slidingExpiration);

            string jsonData = value == null ? null : JsonConvert.SerializeObject(value);
            _redisServer.Database.StringSet(key, jsonData);
        }
    }
}
