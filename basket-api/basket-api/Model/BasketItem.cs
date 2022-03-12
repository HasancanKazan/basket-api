using System;

namespace basket_api.Model
{
    public class BasketItem
    {
        public string Id { get; set; }
        public string ProductId { get; set; }
        public string VariantId { get; set; }
        public string ProductName { get; set; }
        public decimal Price { get; set; }
        public int Quantity { get; set; }
        public DateTime CreatedAt { get; set; }
        public DateTime UpdatedAt { get; set; }
    }
}
