# boss-stock

## Boss Stock Inventory Management Application for Phone Stores

## Telefon Mağazaları için stok takip uygulaması

  
### Endpoints:

BASE_URL : http://<ip_address>:3000/api/v1/
BASE_PORT : 3000

`app.Get("/", status)` //API status

//TODO: Edit Endpoints

### Ürün Listele (GET)
`BASE_URL:3000/api/v1/product`
### Ürün Listele Single (GET)
`BASE_URL:3000/api/v1/product/:id`
### Ürün Listele by Barkod Number (GET)
`BASE_URL:3000/api/v1/product/barcode/:id`
### Ürün Listele by Kategori ID (GET)
`BASE_URL:3000/api/v1/product/category/:id`
### Ürün Sil (DELETE)
`BASE_URL:3000/api/v1/product/:id`
### Ürün Ekle (POST)
`BASE_URL:3000/api/v1/product/new`


### Tamirdeki Cihaz Listele (GET)
`BASE_URL:3000/api/v1/repair`
### Tamirdeki Cihaz Listele Single (GET)
`BASE_URL:3000/api/v1/repair/:id`
### Tamirdeki Cihaz Listele by Cihaz Durumu ID (GET)
`BASE_URL:3000/api/v1/repair/status/:id`
### Tamirdeki Cihaz Sil (DELETE)
`BASE_URL:3000/api/v1/repair/:id`
### Tamirdeki Cihaz Ekle (POST)
`BASE_URL:3000/api/v1/repair/new`


### Mağaza Listele (GET)
`BASE_URL:3000/api/v1/store`
### Mağaza Listele Single (GET)
`BASE_URL:3000/api/v1/store/:id`
### Mağaza Sil (DELETE)
`BASE_URL:3000/api/v1/store/:id`
### Mağaza Ekle (POST)
`BASE_URL:3000/api/v1/store/new`


### Kategori Listele (GET)
`BASE_URL:3000/api/v1/category`
### Kategori Listele Single (GET)
`BASE_URL:3000/api/v1/category/:id`
### Kategori Sil (DELETE)
`BASE_URL:3000/api/v1/category/:id`
### Kategori Ekle (POST)
`BASE_URL:3000/api/v1/category/new`


  

Repo clone sonrası aşağıdakileri çalıştır.

  

`mkdir boss-stock`

`cd boss-stock`

`go mod init boss-stock`

  

`go get -u github.com/gofiber/fiber/v2`

`go get -u github.com/mattn/go-sqlite3`

`go get -u gorm.io/gorm`

`go get -u gorm.io/driver/sqlite`

  

`go run main.go`



