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
curl -i -X POST -H "Content-Type: application/json" -d '{"username":"shinigami", "password":"12345678"}' http://<host>:<port>/api/v1/public/signup
```

#### Response:

```
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 29 May 2022 06:24:41 GMT
Content-Length: 26

{"status":"User Created"}
```

Once the user is created you can sign in:

```
curl -i -X POST -H "Content-Type: application/json" -d '{"username":"shinigami", "password":"12345678"}' http://<host>:<port>/api/v1/public/login
```

#### Response:

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 29 May 2022 06:25:17 GMT
Content-Length: 176

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTM4MDU4MTcsImlzcyI6IlNhbSBTZXBpb2wiLCJBdHRyaWJ1dGUiOiJzaGluaWdhbWkifQ.CdEL0FqZxOHAit5C6zfpcX2HuhLESDpwcKQSzlowm2s"}
```

After obtaining the token, we can send it to a special endpoint that will give us the user's data as a response.

```
curl -i -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTM4MDU4MTcsImlzcyI6IlNhbSBTZXBpb2wiLCJBdHRyaWJ1dGUiOiJzaGluaWdhbWkifQ.CdEL0FqZxOHAit5C6zfpcX2HuhLESDpwcKQSzlowm2s" http://<host:port>/a
pi/v1/private/users/user
```

#### Response

```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 29 May 2022 17:44:38 GMT
Content-Length: 201

{"id":"02a542e5-fbd8-46ae-b0dd-7043bb226c9f","username":"shinigami","password":"","ID":0,"CreatedAt":"2022-05-05T07:25:37.560142-05:00","UpdatedAt":"2022-05-05T07:25:37.560142-05:00","DeletedAt":null}
```

**Note:** The password is already hashed but even like that, we don't want to expose our password hash


## Running Unit Test

```
go test -v -coverprofile=coverage.out ./... ./...
```


## Showing Coverage in Browser

```
go tool cover -html=coverage.out
```

## DEMO

A demo application is deployed here:


- [demo app](https://go-authentication-api.azurewebsites.net)
