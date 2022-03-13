using basket_api.Constant;
using basket_api.Core.Redis;
using basket_api.Model;
using basket_api.Model.DTOs;
using Microsoft.Extensions.Configuration;
using MizuCaseStudy.Core.Utilities;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace basket_api.Manager
{
    public class BasketManager : IBasketManager
    {
        private static readonly HttpClient client = new HttpClient();
        private readonly IRedisManager redisManager;
        private readonly IConfiguration configuration;

        public BasketManager(IRedisManager redisManager, IConfiguration configuration)
        {
            this.redisManager = redisManager;
            this.configuration = configuration;
        }



        public async Task<DataResult<UserBasket>> AddToBasket(AddBasketRequestModel basketItem, string userId)
        {
            if (string.IsNullOrEmpty(userId))
                return new ErrorDataResult<UserBasket>(Messages.UserNotFound);

            var product = await GetProductById(basketItem.ProductId, basketItem.VariantId);
            var stock = await GetStockById(basketItem.ProductId, basketItem.VariantId);

            var validateResult = BasketValidate(product, stock, basketItem.Quantity);
            if (!validateResult.Success)
                return new ErrorDataResult<UserBasket>(validateResult.Message);

            var basket = await GetBasket(userId);
            UpsertBasket(basket, product, userId, basketItem.Quantity);

            return new SuccessDataResult<UserBasket>(basket, Messages.ProductAddSuccessfully);
        }

        public async Task<UserBasket> GetBasket(string userId)
        {
            var userBasket = await redisManager.Get("basket_" + userId);
            if (userBasket == null)
                return new UserBasket();

            return JsonConvert.DeserializeObject<UserBasket>(userBasket.ToString());
        }

        private async Task<ProductDetailDTO> GetProductById(string productId, string variantId)
        {
            try
            {
                //Golang product servisinden ürün çekilir.
                HttpResponseMessage productResponse = await client.GetAsync(configuration.GetSection("Services:ProductsService").Value + $"?productId={productId}&variantId={variantId}");
                productResponse.EnsureSuccessStatusCode();
                var responseBody = await productResponse.Content.ReadAsStringAsync();

                return JsonConvert.DeserializeObject<ProductDetailDTO>(responseBody);
            }
            catch (Exception)
            {
                throw;
            }
        }

        private async Task<StockDetailDTO> GetStockById(string productId, string variantId)
        {
            try
            {
                //Golang stock servisinden ürüne ait stok çekilir.
                HttpResponseMessage stockResponse = await client.GetAsync(configuration.GetSection("Services:StocksService").Value + $"?productId={productId}&variantId={variantId}");
                stockResponse.EnsureSuccessStatusCode();
                var responseBody = await stockResponse.Content.ReadAsStringAsync();

                return JsonConvert.DeserializeObject<StockDetailDTO>(responseBody);
            }
            catch (Exception)
            {
                throw;
            }
        }


        private IResult BasketValidate(ProductDetailDTO product, StockDetailDTO stock, int quantity)
        {
            if (String.IsNullOrWhiteSpace(product.Id))
                return new ErrorResult(Messages.ProductNotFound);

            if (quantity <= 0)
                return new ErrorResult(Messages.QuantityInvalid);

            if (String.IsNullOrWhiteSpace(stock.Id))
                return new ErrorResult(Messages.StockInformationNotFound);

            if (quantity > stock.RemainingStock)
                return new ErrorResult(Messages.OutOfStock);

            if (product.Price < 0)
                return new ErrorResult(Messages.PriceCannotBeNegative);

            return new SuccessResult();
        }


        private async void UpsertBasket(UserBasket basket, ProductDetailDTO product, string userId, int quantity)
        {
            if (basket.AllItems == null)
            {
                basket.UserId = userId;
                basket.Discount = Discount();
                basket.AllItems = new List<BasketItem>
                {
                    new BasketItem
                    {
                        Id = Convert.ToString(Guid.NewGuid()),
                        ProductId = product.Id,
                        VariantId = product.VariantId,
                        ProductName = product.Name,
                        Price = product.Price,
                        Quantity = quantity,
                        CreatedAt = DateTime.UtcNow,
                        UpdatedAt = DateTime.UtcNow
                    }
                };
            }
            else
            {
                BasketItem item = basket.AllItems.Find(x => x.ProductId == product.Id && x.VariantId == product.VariantId);

                if (item != null)
                {
                    item.Quantity += quantity;
                    item.UpdatedAt = DateTime.UtcNow;
                }
                else
                    basket.AllItems.Add(new BasketItem
                    {
                        Id = Convert.ToString(Guid.NewGuid()),
                        ProductId = product.Id,
                        VariantId = product.VariantId,
                        ProductName = product.Name,
                        Price = product.Price,
                        Quantity = quantity,
                        CreatedAt = DateTime.UtcNow,
                        UpdatedAt = DateTime.UtcNow
                    });
            }
            await redisManager.Set("basket_" + userId, basket, DateTimeOffset.UtcNow, null, false);
        }


        //Discount servisinde ürüne ait indirim varsa uygulanır. Her ürün için varmış gibi davranıyorum.
        private decimal Discount()
        {
            Random discount = new Random();
            return discount.Next(10, 50);
        }
    }
}
