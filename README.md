# Golang Simple User Authentication Example

This project is simple User Authentication example that using JWT, password hash and PostgreSQL in golang.

## Clone Project
    
    git clone https://github.com/jhalitaksoy/golang-user-authentication.git

    cd golang-user-authentication

    go get

## .env File

Set Enviroment Variables in .env file.

    PORT=:4000
    SECRET_KEY=rbmvsa34mvb948vs372vm9n2ş4jn32avnv
    DB_ADDR=:5432
    DB_USER=postgres
    DB_PASSWORD=1234
    DB_DATABASE=go-user-auth-demo

## Database

Create Database named as {DB_DATABASE}

Create Tables

    CREATE TABLE public.users
    (
        id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
        name text COLLATE pg_catalog."default",
        CONSTRAINT users_pkey PRIMARY KEY (id),
        CONSTRAINT name UNIQUE (name)
    )

    TABLESPACE pg_default;

    ALTER TABLE public.users
        OWNER to postgres;

    
    CREATE TABLE public.passwords
    (
        id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
        user_id integer NOT NULL,
        hash text COLLATE pg_catalog."default" NOT NULL,
        CONSTRAINT passwords_pkey PRIMARY KEY (id),
        CONSTRAINT user_id UNIQUE (user_id),
        CONSTRAINT "user" FOREIGN KEY (user_id)
            REFERENCES public.users (id) MATCH SIMPLE
            ON UPDATE NO ACTION
            ON DELETE NO ACTION
            NOT VALID
    )

    TABLESPACE pg_default;

    ALTER TABLE public.passwords
        OWNER to postgres;  


## Dependencies

    go get -u github.com/gin-gonic/gin

    get get -u github.com/gin-contrib/cors

    go get -u github.com/joho/godotenv

    go get -u golang.org/x/crypto/...

    go get github.com/dgrijalva/jwt-go
    
    go get github.com/go-pg/pg/v10