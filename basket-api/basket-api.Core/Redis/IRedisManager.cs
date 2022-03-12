using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace basket_api.Core.Redis
{
    public interface IRedisManager
    {
        Task<bool> Exists(string key);
        Task<object> Get(string key);
        Task Set(string key, object value, DateTimeOffset absoluteExpiration, TimeSpan? slidingExpiration, bool removable);
        Task Remove(string key);
        void RemoveAll();
    }
}
