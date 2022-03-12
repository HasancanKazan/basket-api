using basket_api.Model;
using System.Threading.Tasks;

namespace basket_api.Manager
{
    public interface IBasketManager
    {
        Task<UserBasket> AddToBasket(AddBasketRequestModel basketRequestModel, string userId);
        Task<UserBasket> GetBasket(string userId);
    }
}
