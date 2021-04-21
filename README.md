# Golang Simple User Authentication Example

This project is simple User Authentication example that using JWT, password hash and PostgreSQL in golang.

## Clone Project
    
    git clone github.com/jhalitaksoy/golang-user-authentication

## .env File

You have to set Enviroment Variables in .env file.

Example 

    PORT=:4000
    SECRET_KEY=rbmvsa34mvb948vs372vm9n2ş4jn32avnv
    DB_ADDR=:5432
    DB_USER=postgres
    DB_PASSWORD=1234
    DB_DATABASE=go-user-auth-demo

## Dependencies

    go get -u github.com/gin-gonic/gin

    get get -u github.com/gin-contrib/cors

    go get -u github.com/joho/godotenv

    go get -u golang.org/x/crypto/...

    go get github.com/dgrijalva/jwt-go
    
    go get github.com/go-pg/pg/v10