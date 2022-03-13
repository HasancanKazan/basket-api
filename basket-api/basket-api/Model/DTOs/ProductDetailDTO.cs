namespace basket_api.Model.DTOs
{
    public class ProductDetailDTO
    {
        public string Id { get; set; }
        public string VariantId { get; set; }
        public string Code { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public decimal Price { get; set; }
        public string Currency { get; set; }
    }
}
