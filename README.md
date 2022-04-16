# basket-api
Bir e-ticaret uygulamasının sepete ekleme bölümünü .net core 5 ile basitçe simüle etmiş, uygulama içerisinde sepete eklenen ürünleri redis'e yazıyor ve okuyor olduğumuz,
product ve stock için mock datalar barındıran 2 tane küçük golang apisine sahip, her bir projenin dockerize edilerek son durumda docker-compose ile tek komutta sistemin
çalışabilir olduğu bir sepete ekleme uygulaması.


Proje test edilebilmesi ve servislerin uygun şekilde çalışabilmesi için dockerize edilmiştir. Powershell ya da cmd açılarak docker yüklü bir ortamda aşağıdaki komut
çalıştırılmalı ve uygulamanın host edilmesi sağlanmalıdır. **Docker Desktop üzerinde Switch to Linux Containers seçilerek aşağıdaki komutlar çalıştırılmalıdır.**

**Not: docker-compose.yml’ın yüklü olduğu dizinde komutu çalıştırıyor olmalıyız.**

<img width="482" alt="image" src="https://user-images.githubusercontent.com/37241737/163678475-4f5575d0-cd80-4c4a-be7c-e8d591be02ee.png">

<img width="246" alt="image" src="https://user-images.githubusercontent.com/37241737/163678493-abbe9aea-8fc9-4738-9fe7-6b336255ac4d.png">

**1.	docker compose up -d (Uygulamamızın servislerini ayağa kaldıracaktır.)**

**2.	docker ps (ayakta olan containerlarımızı listeliyoruz.)**

**3.	basket-api_basket-api_1 isimli container’ın port bilgisini öğreniyoruz. Örneğimizde localhost:36322 şeklinde gelmiştir.**

**4.	Ardından browser üzerinden localhost:36322/swagger çağrılmalıdır.**

![1_bqdHZjHFuoy-qDu2200x3g](https://user-images.githubusercontent.com/37241737/163678630-a2d3a9cb-239f-4607-ad03-287f1182ec0a.png)

Proje temel olarak 4 ayrı servisten oluşmaktadır. Bu servisler sırasıyla aşağıdaki gibidir.

**1.	products-api (golang) → localhost:8001 portunda çalışır**

**2.	stocks-api (golang) → localhost:8002 portunda çalışır**	

**3.	redis (basket datalarını burada tutuyoruz) → redisin default portu localhost:6379 da çalışır**

**4.	basket-api(.net 5 core api) → container ayağa kalktığında boş olan bir porta atacaktır.**

products ve stocks apileri mock data sunan 2 ayrı api şeklinde tasarlanmıştır ve aşağıdaki endpointlere sahiptir.

•	http://localhost:8001/allproducts (Varolan tüm mock product listesini döndürür.)

•	http://localhost:8001/products?productId=xxxxx&variantId=zzzzzz (İçeride bulunan belli bir
product ve variant Id bilgisine göre 1 adet product döndürür.)

•	http://localhost:8002/stocks?productId=xxxxx&variantId=yyyyy (Belli bir product ve variantId bilgisine göre stok toplam stok ve kalan stok bilgisini döner.)

Basket apisi içerisinde 2 adet endpoint yer almaktadır. Bunlar /GET ve /PUT şeklindedir. Swagger üzerinde beklenen model ve parametreler yer almaktadır.

**NOT: Uygulamada herhangi bir AUTH işlemi ve session gibi bilgiler bulunmadığı için sepete eklenecek ürünün hangi kullanıcıya ait olduğu bilgisi header üzerindeki
userId parametresi ile alınarak bir kullanıcı varmış gibi davranılmıştır.**

**Örnek Request 1**
curl  -X  PUT  "http://localhost:4759/api/basket"  -H    "accept:  text/plain"  -H    "userId:  5"  -H    "Content-Type:  application/json"  
-d "{\"productId\":\"86287ba9-e561-4a62-bb7d-4e958e49c15a\",\"variantId\":\"62421c08-6302-40be-b59e-2c70777a9ab9\",\"quantity\":5}"

**Örnek Request 2**
curl  -X  PUT  "http://localhost:4759/api/basket"  -H    "accept:  text/plain"  -H    "userId:  5"  -H    "Content-Type:  application/json"  
-d "{\"productId\":\"f630adf1-248e-4f80-8f63-63263fbeb50d\",\"variantId\":\"207cee54-e725-4b65-b4f3-ccee7104ca1b\",\"quantity\":1}"

**Örnek Response**
{
"data": { "userId":  "5", "allItems":  [
{
"id":  "990aca00-8c0c-4a6e-8cf3-0aa2aa155615", "productId":  "86287ba9-e561-4a62-bb7d-4e958e49c15a", "variantId":  "62421c08-6302-40be-b59e-2c70777a9ab9", "productName":  "Afrodit  Meyve  Sepeti",
"price":  199,
"quantity":  5,
"createdAt":  "2022-02-28T19:30:09.6814229Z", "updatedAt":  "2022-02-28T19:30:09.6814494Z"
},
{
"id":  "547550db-5232-4f88-8407-97a86e6c7e02",
"productId":  "f630adf1-248e-4f80-8f63-63263fbeb50d", "variantId":  "207cee54-e725-4b65-b4f3-ccee7104ca1b", "productName":  "Karışık  Meyve  Sepeti",
"price":  134.96,
"quantity":  1,
"createdAt":  "2022-02-28T19:30:52.4851159Z", "updatedAt":  "2022-02-28T19:30:52.4851162Z"
}
],
"discount":  14,
"totalPrice":  1115.96
},
"success":  true,
"message":  "Ürün  Başarıyla  Sepete  Eklendi"
}

