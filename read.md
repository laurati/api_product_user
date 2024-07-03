Swagger

go install github.com/swaggo/swag/cmd/swag@latest

swag init -g cmd/server/main.go

Main imports:
httpSwagger "github.com/swaggo/http-swagger"
_ "github.com/laurati/api_product_user/docs"

Navegador: http://localhost:8000/docs/index.html

Authorize: Bearer token