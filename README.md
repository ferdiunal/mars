# Mars

Mars provides a fast authentication system for SPAs(Single Page Applications), mobile application.

## Install

```
go get -u https://github.com/ferdiunal/mars@v1.0.0-alpha
```

## ⚡️ Quickstart
```go
package main
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
 
  "https://github.com/ferdiunal/mars"
)

func main() {
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    
    mars := NewMars(db)
    
    token := mars.CreateToken(1, "Login") // or set abilities mars.CreateToken(1, "Login", ["*"])
    fmt.Println(token) 
    // Output: {846e92272d6f17ebf769a9e8f7367312c592bb718dae05ccc3069b22d469b2bf 2021-09-22 00:44:05.4878583 +0300 +03 m=+0.006680401 [*]}
}

```
