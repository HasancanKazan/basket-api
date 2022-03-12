using Microsoft.Extensions.Configuration;
using StackExchange.Redis;
using System;

namespace basket_api.Core.Redis
{
    public class RedisServer
    {
        private readonly IConfiguration _configuration;
        private readonly ConnectionMultiplexer _connectionMultiplexer;
        private readonly IDatabase _database;

        public IDatabase Database => _database;
        public RedisServer(IConfiguration configuration)
        {
            _configuration = configuration;

            _connectionMultiplexer = ConnectionMultiplexer.Connect(GetConnectionString());
            _database = _connectionMultiplexer.GetDatabase(GetCurrentDatabase());
        }


        private int GetCurrentDatabase()
        {
            var dbIndex = _configuration.GetSection("RedisConnection:Database").Value;
            return Convert.ToInt16(dbIndex);
        }

        private string GetConnectionString()
        {
            var host = _configuration.GetSection("RedisConnection:RedisHost").Value;
            var port = _configuration.GetSection("RedisConnection:RedisPort").Value;
            return host + ":" + port + ",abortConnect=False";
        }

        //Tüm Dbyi boşaltabilmek için FlushDatabase metodunu kullanacağız.
        public void FlushDatabase()
        {
            _connectionMultiplexer.GetServer(GetConnectionString()).FlushDatabase(GetCurrentDatabase());
        }
    }
}
