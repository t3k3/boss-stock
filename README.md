# boss-stock

## Boss Stock Inventory Management Application for Phone Stores

## Telefoncular için stok takip uygulaması

  

### Endpoints:

BASE_URL : http://<ip_address>:3000/api/v1/
BASE_PORT : 3000

`app.Get("/", status)` //API status

//TODO: Edit Endpoints

### Ürün Listele
`BASE_URL:3000/api/v1/product`
### Ürün Ekle
`BASE_URL:3000/api/v1/product/new`


### Tamirdeki Cihazları Listele
`BASE_URL:3000/api/v1/repair"`
### Tamir İçin Cihaz Ekle
`BASE_URL:3000/api/v1/repair/new`

### Mevcut Mağazaları Getir
`BASE_URL:3000/api/v1/store`
### Yeni Mağaza Ekle
`BASE_URL:3000/api/v1/store/new`
  
  

  

Repo clone sonrası aşağıdakileri çalıştır.

  

`mkdir boss-stock`

`cd boss-stock`

`go mod init boss-stock`

  

`go get -u github.com/gofiber/fiber/v2`

`go get -u github.com/mattn/go-sqlite3`

`go get -u gorm.io/gorm`

`go get -u gorm.io/driver/sqlite`

  

`go run main.go`



### NOTLAR
gorm.Find 
gorm.Model

ya da

gorm.Where
gorm.Model

peşpeşe çağırılarak update yapılabilir

örnek:
```
var burki Southwind.Employee
db.Find(&burki, "ID=?", 1) //Önce
db.Model(&burki).Update("LastName", "Selim Senyurt")
WriteToScreen(burki)
````

```
var buffon Southwind.Employee

db.Model(&buffon).Where("ID=?", 2).Updates(map[string]interface{}{"FirstName": "Cianluici", "LastName": "Buffon"})
db.First(&buffon, 2) //Direkt primary key üstünden(varsayılan olarak ID) arama yapar
WriteToScreen(buffon)
} else {
fmt.Println(err.Error())
}
```




