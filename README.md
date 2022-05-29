# go-authentication-api

A simple authentication API.

### What you can do:
- add new users
- sign in with your user when is already registered
- get a JWT if authentication is successful

## Running Locally

```
git clone https://github.com/Edmartt/go-authentication-api
```


or ssh instead:


```
git@github.com:Edmart/go-authentication-api
```


```
cd go-authentication-api
```

```
go mod tidy
```

```
go run main.go
```


## Using

**Note**

This is actually set for using with SQLite for simplicity, but you can switch this database manager for postgres if needed. Just go to internal/users/data/repository.go and change the line 10 typing **Postgres**. If you use postgres, remember to set the .env.example values according to your setup. Note that I left the `:` as part of HTTP_PORT env var value, so if you need to change this port just delete the numbers.

You can use any client you want. I'm using curl:

```
curl -i -X POST -H "Content-Type: application/json" -d '{"username":"shinigami", "password":"12345678"}' http://localhost:8081/api/v1/users/signup
```

This is my response:

```
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 29 May 2022 06:24:41 GMT
Content-Length: 26

{"status":"User Created"}
```

Once the user is created you can sign in:

```
curl -i -X POST -H "Content-Type: application/json" -d '{"username":"shinigami", "password":"12345678"}' http://localhost:8081/api/v1/users/login
```

This is my response:

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 29 May 2022 06:25:17 GMT
Content-Length: 176

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTM4MDU4MTcsImlzcyI6IlNhbSBTZXBpb2wiLCJBdHRyaWJ1dGUiOiJzaGluaWdhbWkifQ.CdEL0FqZxOHAit5C6zfpcX2HuhLESDpwcKQSzlowm2s"}
```

## Running Unit Test

```
go test -v -coverprofile=coverage.out ./... ./...
```


## Showing Coverage in Browser

```
go tool cover -html=coverage.out
```
