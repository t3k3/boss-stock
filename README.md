# boss-stock

## Boss Stock Inventory Management Application for Phone Stores

  

## Telefoncular için stok takip uygulaması

  

### Endpoints:

`app.Get("/", status)`

//TODO: Edit Endpoints

`app.Get("/api/v1/product", product.GetAllProducts)`

`app.Post("/api/v1/product", product.SaveProduct)`

  
  

### Address:

`127.0.0.1:3000/api/v1/product`

`app.Listen(":3000")`

  

Repo clone sonrası aşağıdaki çalıştır.

  

`mkdir boss-stock`

`cd boss-stock`

`go mod init boss-stock`

  

`go get -u github.com/gofiber/fiber/v2`

`go get -u github.com/mattn/go-sqlite3`

`go get -u gorm.io/gorm`

`go get -u gorm.io/driver/sqlite`

  

`go run main.go`




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


