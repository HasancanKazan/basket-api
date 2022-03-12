using basket_api.Model;
using System.Threading.Tasks;

namespace basket_api.Manager
{
    public class BasketManager : IBasketManager
    {
        public Task<UserBasket> AddToBasket(AddBasketRequestModel basketRequestModel, string userId)
        {
            throw new System.NotImplementedException();
        }

        public Task<UserBasket> GetBasket(string userId)
        {
            throw new System.NotImplementedException();
        }
    }
}
