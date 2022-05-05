# go-authentication-api

A simple authentication API.

### What you can do:
- add new users
- sign in with your user when is already registered
- get a JWT if authentication is successful

## Running Locally

- ```git clone https://github.com/Edmartt/go-authentication-api```


or ssh instead:


- ```git@github.com:Edmart/go-authentication-api```

- ```cd go-authentication-api```

- ```go run main.go```


## Using

**Note**

This is actually set for using with SQLite for simplicity, but you can switch this database manager for postgres if needed. Just go to internal/users/data/repository.go and change the line 10 typing **Postgres**

You can use any client you want. I'm using curl:

```
curl -i -X POST -H "Content-Type: application/json" -d '{"username":"shinigami", "password":"12345678"}' http://localhost:8081/api/v1/users/signup
```

This is my response:

**HTTP/1.1 201 Created
Content-Type: application/json
Date: Thu, 05 May 2022 11:33:46 GMT
Content-Length: 0**

Once the user is created you can sign in:

```
curl -i -X POST -H "Content-Type: application/json" -d '{"username":"shinigami", "password":"12345678"}' http://localhost:8081/api/v1/users/login
```

This is my response:

**HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 05 May 2022 11:36:33 GMT
Content-Length: 140

"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTE3NTA4OTMsIkF0dHJpYnV0ZSI6InNoaW5pZ2FtaSJ9.njMWOvu6PBusDZt9k9lyfIRGRBvCo61Mm_jg6xhFPL0"**
