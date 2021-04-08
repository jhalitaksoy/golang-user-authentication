## Install Dependencies

Step 1)
    Install gin

    go get -u github.com/gin-gonic/gin

Step 2)
    Install cors

    get get -u github.com/gin-contrib/cors

Step 3)
    Install godotenv

    go get -u github.com/joho/godotenv

Step 4)
    Install crypto

    go get -u golang.org/x/crypto/...

Step 5) 
    Install JWT

    go get github.com/dgrijalva/jwt-go
    
Step 6) 
    Install pg
    
    go get github.com/go-pg/pg/v10

## .env File

You have to set SECRET_KEY in .env file.

    SECRET_KEY=your_long_secret_key