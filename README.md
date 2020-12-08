# fullstack-go

## Copy example.env

> cp example.env .env

---

## Change ENV file to your credentials and values of the Database 

---

## Install depedences

```go
go get github.com/badoux/checkmail
go get github.com/jinzhu/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm/dialects/mysql
go get github.com/joho/godotenv
go get gopkg.in/go-playground/assert.v1
go get github.com/nsf/jsondiff 
```

## Run application withou run the containers

> go run main.go

## Testing the Endpoints in Postman

### There is a file with all the endpoints to import into the postman

- [Home](http://localhost:8080/)               -  Without token
- [PostLogin](http://localhost:8080/posts)     -  email: sergiolns75@gmail.com, password: password
- [GetPosts](http://localhost:8080/posts)      -  Without token
- [CreatePost](http://localhost:8080/posts)    -  **With token**
- [UpdatePost](http://localhost:8080/posts/1)  -  **With token**
- [GetPost](http://localhost:8080/posts/1)     -  Without token
- [DeletePost](http://localhost:8080/posts/1)  -  **With token**
- [GetUsers](http://localhost:8080/users)      -  Without token
- [UpdateUser](http://localhost:8080/users/2)  -  **With token**
- [DeleteUser](http://localhost:8080/users/2)  -  **With token**
- [CreateUser](http://localhost:8080/users)    -  Without token
- [GetUser](http://localhost:8080/users/1)     -  Without token
- [PostCompare](http://localhost:8080/compare) -  Without token(**to compare two JSONs**)

## Run containers

> docker-compose up -d

### After run the containers, Test the endpoints with Postman