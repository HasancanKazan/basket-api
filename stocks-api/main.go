package main

import "github.com/labstack/echo"

type StockResponse struct {
	ID             string `json:"Id"`
	ProductId      string
	VariantId      string
	TotalStock     int
	RemainingStock int
}

var mockStocks []StockResponse

func main() {
	mockStockPrepare()
	e := echo.New()
	e.GET("/stocks", getStock)
	e.Start(":8002")
}

func getStock(c echo.Context) error {
	productId := c.QueryParam("productId")
	variantId := c.QueryParam("variantId")

	for _, stock := range mockStocks {
		if stock.ProductId == productId && stock.VariantId == variantId {
			return c.JSON(200, stock)
		}
	}
	return c.JSON(200, StockResponse{})
}

func mockStockPrepare() {
	stock1 := StockResponse{
		ID:             "1fa1f6b2-8933-442d-9d05-d980c97cb085",
		ProductId:      "86287ba9-e561-4a62-bb7d-4e958e49c15a",
		VariantId:      "62421c08-6302-40be-b59e-2c70777a9ab9",
		RemainingStock: 50,
		TotalStock:     250,
	}
	stock1_1 := StockResponse{
		ID:             "0c22c07b-94fb-4647-8131-41dc0c0c1a87",
		ProductId:      "86287ba9-e561-4a62-bb7d-4e958e49c15a",
		VariantId:      "d7c6b6ce-a584-44a3-a45b-0f8b5296c9ed",
		RemainingStock: 50,
		TotalStock:     250,
	}
	stock2 := StockResponse{
		ID:             "9fd6f33b-f334-485a-af19-b9980307a215",
		ProductId:      "f630adf1-248e-4f80-8f63-63263fbeb50d",
		VariantId:      "207cee54-e725-4b65-b4f3-ccee7104ca1b",
		RemainingStock: 1,
		TotalStock:     152,
	}
	stock3 := StockResponse{
		ID:             "9443e0b9-b6b7-430e-9574-6ad869eac76d",
		ProductId:      "9c20fca2-dda9-4122-8e77-aa70c5ee08c5",
		VariantId:      "06e28fd0-be73-4262-8e75-42c568248729",
		RemainingStock: 0,
		TotalStock:     60,
	}
	stock4 := StockResponse{
		ID:             "7b522ad3-560c-48b8-ba50-ea0c4ac01eee",
		ProductId:      "a0a4e7fb-d8c4-42a4-bbcd-52a2e0dc8de3",
		VariantId:      "ddba161c-8ac4-44f3-9233-6b81b943c975",
		RemainingStock: 12,
		TotalStock:     98,
	}
	stock5 := StockResponse{
		ID:             "a08d19a9-1604-490d-89ae-2149ec0f7162",
		ProductId:      "5ebb9786-dcb1-4217-b747-dfd196c34f06",
		VariantId:      "c69e94be-ae2b-478b-91cb-282c33e33bb5",
		RemainingStock: 0,
		TotalStock:     450,
	}
	stock6 := StockResponse{
		ID:             "8ec23779-4689-4d42-8c35-2733b8fde342",
		ProductId:      "b12ddd73-66e5-4728-b489-a1e556ce2a63",
		VariantId:      "9d6d58b9-f4dd-40b2-8bad-b5d0f01cb88e",
		RemainingStock: 0,
		TotalStock:     35,
	}
	stock7 := StockResponse{
		ID:             "db5355f2-f280-4dba-8eaf-2e039701a7e9",
		ProductId:      "2b95606c-9f0d-4ee3-8b22-624e42b8e1f5",
		VariantId:      "13ce0f3a-cc5e-442d-b494-648703da4f07",
		RemainingStock: 36,
		TotalStock:     65,
	}
	stock8 := StockResponse{
		ID:             "5eb6fbd1-545e-409d-b8e3-a6c84dffac8f",
		ProductId:      "370048db-cacb-495c-8d96-de85e07821ea",
		VariantId:      "fe67533d-fef6-47e9-a023-116b183d32fb",
		RemainingStock: 154,
		TotalStock:     357,
	}
	stock9 := StockResponse{
		ID:             "530335ee-013b-4262-8ddb-52df59e71394",
		ProductId:      "4519c0f6-ae7a-4a78-9531-a4ccb5b091f4",
		VariantId:      "6e65f498-fb47-4814-9594-11f49746a55c",
		RemainingStock: 254,
		TotalStock:     1200,
	}
	stock10 := StockResponse{
		ID:             "244d638a-a178-439c-ad31-025713fe33fb",
		ProductId:      "fb62685e-8e23-4971-b049-e60a31b46a92",
		VariantId:      "94d3c5d1-1538-4444-8919-7a2f3813b4bd",
		RemainingStock: 0,
		TotalStock:     25,
	}
	mockStocks = append(mockStocks, stock1)
	mockStocks = append(mockStocks, stock1_1)
	mockStocks = append(mockStocks, stock2)
	mockStocks = append(mockStocks, stock3)
	mockStocks = append(mockStocks, stock4)
	mockStocks = append(mockStocks, stock5)
	mockStocks = append(mockStocks, stock6)
	mockStocks = append(mockStocks, stock7)
	mockStocks = append(mockStocks, stock8)
	mockStocks = append(mockStocks, stock9)
	mockStocks = append(mockStocks, stock10)
}
