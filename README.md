# structure-app-mvc

## Estrutura Atual

<img width="260" alt="Captura de Tela 2023-05-08 às 18 46 03" src="https://user-images.githubusercontent.com/2454085/236943195-c2b9431c-3502-41a5-9eac-b62a511966fa.png">


## Bibliotecas Necessárias - Comando de Instalação

JWT - https://github.com/golang-jwt/jwt
- go get -u github.com/golang-jwt/jwt/v5

Gorm - https://gorm.io/docs/index.html
- go get -u gorm.io/gorm
- go get -u gorm.io/driver/postgres

DotEnv - https://github.com/joho/godotenv
- go get github.com/joho/godotenv

Viper - https://github.com/spf13/viper
- go get github.com/spf13/viper


Compile Daemon - https://github.com/githubnemo/CompileDaemon
- go get github.com/githubnemo/CompileDaemon
- To compile my app command : ./CompileDaemon —command=“./myproject”

OBS.: O compile Daemon irá manter o servidor ativo durante o seu processo de codificação, assim suas mudanças no código serão refletidas em tempo real, sem haver a necessidade de derrubar o servidor sempre que quisermos observar 
alguma alteração realizada.

Gin Framework - https://github.com/gin-gonic/gin
- go get -u github.com/gin-gonic/gin

Crypto - 
- go get -u golang.org/x/crypto/bcrypt

Validator
- go get github.com/go-playground/validator/v10

Swagger - https://github.com/swaggo/gin-swagger
- go get -u github.com/swaggo/gin-swagger
- go get -u github.com/swaggo/files

## Documentações 

Go Documentation - https://go.dev/doc/
Middleware Golang - https://drstearns.github.io/tutorials/gomiddleware/
Router Golang - https://pkg.go.dev/go.chromium.org/luci/server/router
Logger - https://github.com/google/logger
