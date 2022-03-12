namespace basket_api.Model
{
    public class AddBasketRequestModel
    {
        public string ProductId { get; set; }
        public string VariantId { get; set; }
        public int Quantity { get; set; }
    }
}
