using System.Collections.Generic;
using System.Linq;

namespace basket_api.Model
{
    public class UserBasket
    {
        public string UserId { get; set; }
        public List<BasketItem> AllItems { get; set; }
        public decimal Discount { get; set; }
        public decimal TotalPrice
        {
            get
            {
                var total = AllItems.Sum(x => x.Price * x.Quantity);
                if (total - Discount < 0)
                    return total;
                return AllItems.Sum(x => x.Price * x.Quantity) - Discount;

            }
        }
    }
}
