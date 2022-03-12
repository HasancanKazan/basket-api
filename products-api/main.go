package main

import "github.com/labstack/echo"

type GetProductResponse struct {
	ID          string  `json:"Id"`
	VariantId   string  `json:"VariantId"`
	Code        string  `json:"Code"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	Price       float64 `json:"Price"`
	Currency    string  `json:"Currency"`
}

var mockProducts []GetProductResponse

func main() {
	mockProductPrepare()
	e := echo.New()
	e.GET("/products", getProductById)
	e.GET("/allproducts", getProducts)
	e.Start(":8001")
}

func getProductById(c echo.Context) error {
	productId := c.QueryParam("productId")
	variantId := c.QueryParam("variantId")

	for _, product := range mockProducts {
		if product.ID == productId && product.VariantId == variantId {
			return c.JSON(200, product)
		}
	}
	return c.JSON(200, GetProductResponse{})
}

func getProducts(e echo.Context) error {
	return e.JSON(200, mockProducts)
}

func mockProductPrepare() {
	product1 := GetProductResponse{
		ID:          "86287ba9-e561-4a62-bb7d-4e958e49c15a",
		VariantId:   "62421c08-6302-40be-b59e-2c70777a9ab9",
		Code:        "kc2715404",
		Name:        "Bargello 112 Floral Edp 50 Ml Kadın Parfüm",
		Description: `Fresh, meyveli, şeker, sıcak baharat, beyaz çiçeksi %100 orjinaldir.`,
		Price:       80,
		Currency:    "TRY"}
	product1_1 := GetProductResponse{
		ID:          "86287ba9-e561-4a62-bb7d-4e958e49c15a",
		VariantId:   "d7c6b6ce-a584-44a3-a45b-0f8b5296c9ed",
		Code:        "kc2715404",
		Name:        "Bargello 112 Floral Edp 50 Ml Kadın Parfüm",
		Description: `Fresh, meyveli, sıcak baharat, kırmızı çiçeksi.`,
		Price:       80,
		Currency:    "TRY"}
	product2 := GetProductResponse{
		ID:          "f630adf1-248e-4f80-8f63-63263fbeb50d",
		VariantId:   "207cee54-e725-4b65-b4f3-ccee7104ca1b",
		Name:        "Karışık Meyve Sepeti",
		Code:        "kc7479030",
		Description: "Tatlı mı tatlı sevdiklerinize, en az onlar kadar tatlı bir hediye vermeye ne dersiniz? Karışık meyvelerden oluşan bu renkli hediye sepeti, sevdiklerinizin önce gözlerine sonra kalbine hitap edecek.",
		Price:       134.96,
		Currency:    "TRY"}
	product3 := GetProductResponse{
		ID:          "9c20fca2-dda9-4122-8e77-aa70c5ee08c5",
		VariantId:   "06e28fd0-be73-4262-8e75-42c568248729",
		Code:        "vbat234",
		Name:        "Love Garden Bouquet",
		Description: "Kelime olarak Arapça ‘’aşaka’’ kelimesinden türeyen aşk; sarmaşmak, sıkıca sarılmak, sarmaşık anlamına gelir. ",
		Price:       359.00,
		Currency:    "TRY"}
	product4 := GetProductResponse{
		ID:          "a0a4e7fb-d8c4-42a4-bbcd-52a2e0dc8de3",
		VariantId:   "ddba161c-8ac4-44f3-9233-6b81b943c975",
		Code:        "vbat489",
		Name:        "Flormar Kozmetik Renkli Göz Kalemleri 12adet 12renkrenk",
		Description: "12 adet renkli göz kalemi. Görseldeki renkler ya da görsele yakın tonlar gönderilecektir",
		Price:       54.50,
		Currency:    "TRY"}
	product5 := GetProductResponse{
		ID:          "5ebb9786-dcb1-4217-b747-dfd196c34f06",
		VariantId:   "c69e94be-ae2b-478b-91cb-282c33e33bb5",
		Code:        "vbat390",
		Name:        "Feel Happy",
		Description: "Lila Gül : 6 Adet,Eryngium Aquarius Qstar-Boğa Dikeni : 4 Adet,Rosa Tr Silver Shadow Lila Çardak Gül : 12 Adet",
		Price:       339.00,
		Currency:    "TRY"}
	product6 := GetProductResponse{
		ID:          "b12ddd73-66e5-4728-b489-a1e556ce2a63",
		VariantId:   "9d6d58b9-f4dd-40b2-8bad-b5d0f01cb88e",
		Code:        "kc5447381",
		Name:        "Kişiye Özel Spotify Barkodlu Plak",
		Description: "Sevdiklerinize sizin için özel anlamlar ifade eden şarkıları armağan edin ya da özel anlarınızı anlamlı şarkılarla taçlandırın!",
		Price:       54.90,
		Currency:    "TRY"}
	product7 := GetProductResponse{
		ID:          "2b95606c-9f0d-4ee3-8b22-624e42b8e1f5",
		VariantId:   "13ce0f3a-cc5e-442d-b494-648703da4f07",
		Code:        "VTK22-5422",
		Name:        "VATKALI Beli Bağlamalı Saten Elbise Mavi",
		Description: "Beli Bağlamalı Saten Elbise Mavi Ürün Detayları Gömlek yaka, önden yarım düğmeli, uzun kollu,manşeti düğmeli,büzgü detaylı, mini saten elbise Kumaş Bilgisi %100 Viskon Viskon kumaş: Kayın ağacından üretilerek elde edilen bir hammaddedir. Doğal olması ona birçok avantaj sağlamaktadır. Selüloz içeriğe sahip olan viskon, üretim sırasında akışkan bir hal alır ve bu sırada herhangi bir kimyasala uğramaz. Üretim sırasında selülozun değişikliğe uğramamasından dolayı pamuğa benzer bir ürün olmasına katkıda bulunur. ",
		Price:       254.99,
		Currency:    "TRY"}
	product8 := GetProductResponse{
		ID:          "370048db-cacb-495c-8d96-de85e07821ea",
		VariantId:   "fe67533d-fef6-47e9-a023-116b183d32fb",
		Code:        "kc6389380",
		Name:        "Guess GUU1336L3M Kol Saati",
		Description: "Özel Güllü Kutuda Yollanmaktadır.",
		Price:       630.00,
		Currency:    "TRY"}

	product9 := GetProductResponse{
		ID:          "4519c0f6-ae7a-4a78-9531-a4ccb5b091f4",
		VariantId:   "6e65f498-fb47-4814-9594-11f49746a55c",
		Code:        "kc6268667",
		Name:        "Kişiye Özel İsimli Cüzdan",
		Description: "Tesbih siyah cam olup imame süsleri vardı çakmak metal olup gaz doldurmalıdı metal olup siyah çevirmeli olarak açılı touch pen olup dokunmatik ekranlarda rahatlıkla kullanabilirsiniz.",
		Price:       39.99,
		Currency:    "TRY"}
	product10 := GetProductResponse{
		ID:          "fb62685e-8e23-4971-b049-e60a31b46a92",
		VariantId:   "94d3c5d1-1538-4444-8919-7a2f3813b4bd",
		Code:        "658723213456",
		Name:        "TEKNO DÜNYAM Objaks Grafik Digital Çocuk Yazı Tahtası Çizim Tableti Lcd 8.5 Inc Ekran Grafik 8.5 Inç Ekran J.b",
		Description: "DERS ÇALIŞIRKEN KAĞIT İSRAFINA SON GRAFİK DİGİTAL ÇOCUK YAZI ÇİZİM TABLETİ LCD 8.5 INC EKRANLI + BİLGİSAYAR KALEMLİ ÇOCUKLARINIZ EĞLENİRKEN ÖĞRENSİN ÖZELLİKLER: HAFİF RENK, GÖZLERİNİZE ZARAR VERMEZ, RADYASYON YOKTUR. BASINCA DUYARLI, KULLANIMI VE OYNAMASI KOLAYDIR. EKRANI SADECE BİR DÜĞME İLE TEMİZLEYİN. ULTRA İNCE VE HAFİF TASARIM, TAŞIMAK İÇİN UYGUN. YAZMA VE ÇİZİM, EV MESAJI VB. İÇİN UYGUNDUR. AÇIKLAMALARI: BİR CR2016 DÜĞME PİL İLE. YENİDEN YAZILABİLİR TABLET KAĞIT TASARRUFU VE ÇEVRE DOSTU. OKUL, EV VE OFİS KULLANIMI İÇİN UYGUN. ",
		Price:       54.65,
		Currency:    "TRY"}

	mockProducts = append(mockProducts, product1)
	mockProducts = append(mockProducts, product1_1)
	mockProducts = append(mockProducts, product2)
	mockProducts = append(mockProducts, product3)
	mockProducts = append(mockProducts, product4)
	mockProducts = append(mockProducts, product5)
	mockProducts = append(mockProducts, product6)
	mockProducts = append(mockProducts, product7)
	mockProducts = append(mockProducts, product8)
	mockProducts = append(mockProducts, product9)
	mockProducts = append(mockProducts, product10)
}
