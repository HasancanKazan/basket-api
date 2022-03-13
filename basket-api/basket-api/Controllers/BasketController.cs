using basket_api.Manager;
using basket_api.Model;
using Microsoft.AspNetCore.Mvc;
using MizuCaseStudy.Core.Utilities;
using System.Net.Http;
using System.Threading.Tasks;

namespace basket_api.Controllers
{
    [ApiController]
    [Route("api/basket")]
    public class BasketController : ControllerBase
    {
        private readonly IBasketManager basketManager;
        public BasketController(IBasketManager basketManager)
        {
            this.basketManager = basketManager;
        }


        [HttpGet("{userId}")]
        public async Task<IActionResult> Get(string userId)
        {
            var basket = await basketManager.GetBasket(userId);

            if (basket.AllItems == null)
                return NoContent();

            return Ok(basket);
        }


        //Auth yapısını yapmadığımız için userId bilgisini header üzerinden alarak hangi kullanıcı için işlem yapılması gerektiğini belirliyorum.
        [HttpPut]
        public async Task<IDataResult<UserBasket>> Put(AddBasketRequestModel basketItem, [FromHeader(Name = "userId")] string userId)
        {
            return await basketManager.AddToBasket(basketItem, userId);
        }



    }
}
