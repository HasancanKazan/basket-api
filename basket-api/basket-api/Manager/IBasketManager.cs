using basket_api.Model;
using MizuCaseStudy.Core.Utilities;
using System.Threading.Tasks;

namespace basket_api.Manager
{
    public interface IBasketManager
    {
        Task<DataResult<UserBasket>> AddToBasket(AddBasketRequestModel basketItem, string userId);
        Task<UserBasket> GetBasket(string userId);
    }
}
