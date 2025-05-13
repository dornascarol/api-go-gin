# Pagodão :microphone:

_Project feature I learned at Alura_

## Introduction
Developing a CRUD API using the Gin-Gonic framework with the objective of deepening technical knowledge in building a web application with Go, with a theme related to the universe of the pagode musical genre and using PostgreSQL as a database.


## Technologies used
* VS Code
* Clean Architecture
* Go (Golang)
* Package gin-gonic/gin
* Package net/http
* Package log
* Package gorm.io/driver/postgres
* Package gorm.io/gorm
* Docker
* Package validator.v2
* Package testing
* Package net/http/httptest
* Package stretchr/testify/assert


## Specs
➥ Gin-gonic

➥ Validations

➥ Tests


## Tools
The <a href="https://go.dev/doc/install" target="_blank"> Go</a> programming language was installed on Windows, version 1.24.0.

The Go extension was installed in <a href="https://code.visualstudio.com/download" target="_blank"> VS Code</a>, version 0.46.1.

The program <a href="https://insomnia.rest/download" target="_blank" > Insomnia</a> was used to simulate route requests.

The container management tool <a href="https://www.docker.com/products/docker-desktop/" target="_blank" > Docker Desktop</a> was installed, using the version compatible with the operating system.


## Running the project
1. **To get the repository, clone it:**
	```bash
	git clone https://github.com/dornascarol/api-go-gin.git
	cd api-go-gin
	```

2. **Command to check, update, and download the package manager modules (go.mod):**
	```
	go mod tidy
	```

3. **Command to run the project in the terminal:**
	```
	go run main.go
	```

4. **To stop the execution in the terminal: press `"Ctrl"` + `"C"`.**
<br>  

### _Docker_

1. **Command to run the `docker-compose.yml` file with two service images, one for Postgres and another for pgAdmin:**
	```
	docker-compose up
	```

_Another way to run Docker_

1. **To find out the full name of the containers, write the command in the terminal:**
	```
	docker ps -a
	```
  
2. **To run the containers, write the next command in the terminal:**
	```
	docker start gin-api-rest-pgadmin-compose-1 gin-api-rest-postgres-1
	```
  
3. **Stop the containers with the command in the terminal:**
	```
	docker stop gin-api-rest-pgadmin-compose-1 gin-api-rest-postgres-1
	```

- URL to access pgAdmin:
```
localhost:54321
```
<br>  

### _Tests_

1. **To run the tests simultaneously:**
   ```
   go test -v ./...
   ```

_Another way to run Tests_

1. **To run tests individually with more detailed outputs:**
	```
	go test -run TestFailed -v
	```

	```
	go test -run TestGreetingStatusCode -v
	```
 
_Importing package testify/assert_

1. **Command:**
	```
	go get github.com/stretchr/testify/assert@v1.10.0
	```


## Endpoints

| Method    | URL                 | Description                                        |
| --------- | ------------------- | -------------------------------------------------- |
| 1) GET    | /singers            | List all singers                                   | 
| 2) POST   | /singers            | Create a new singer                                | 
| 3) GET    | /singers/:id        | Get information of a specific singer               | 
| 4) DELETE | /singers/:id        | Delete a specific singer                           | 
| 5) PATCH  | /singers/:id        | Edit information of a specific singer              | 
| 6) GET    | /singers/name/:name | Search specific singer by artist name              | 
| 7) GET    | /:name              | Returns the name passed in params inside a message |

.

### 1) All singers
- URL params: localhost:8080/singers
- Method: GET
- Request body: empty
- Response:
Example of 200 - OK 
```
[
	{
		"ID": 1,
		"CreatedAt": "2025-04-25T22:00:12.859752-03:00",
		"UpdatedAt": "2025-04-25T22:00:12.859752-03:00",
		"DeletedAt": null,
		"artist_name": "Péricles",
		"song_name": "Até que durou",
		"musical_genre": "Pagode"
	},
	{
		"ID": 2,
		"CreatedAt": "2025-04-25T22:00:39.474979-03:00",
		"UpdatedAt": "2025-04-25T22:00:39.474979-03:00",
		"DeletedAt": null,
		"artist_name": "Soweto",
		"song_name": "Antes de dizer adeus",
		"musical_genre": "Pagode"
	},
	{
		"ID": 3,
		"CreatedAt": "2025-04-29T20:59:39.757229-03:00",
		"UpdatedAt": "2025-04-29T20:59:39.757229-03:00",
		"DeletedAt": null,
		"artist_name": "Exaltasamba",
		"song_name": "Gamei",
		"musical_genre": "Pagode"
	}
]
```

----------------------------

### 2) Create a singer
- URL params: localhost:8080/singers
- Method: POST
- Request body JSON format: 
```
{
	"artist_name": "Soweto",
	"song_name": "Antes de dizer adeus",
	"musical_genre": "Pagode"
}
```
- Response:
Example of 200 - OK
```
{
	"ID": 2,
	"CreatedAt": "2025-04-25T22:00:39.4749791-03:00",
	"UpdatedAt": "2025-04-25T22:00:39.4749791-03:00",
	"DeletedAt": null,
	"artist_name": "Soweto",
	"song_name": "Antes de dizer adeus",
	"musical_genre": "Pagode"
}
```

----------------------------

### 3) Singer by ID
- URL params: localhost:8080/singers/1
- Method: GET
- Request body: empty
- Responses:
Example of 200 - OK
```
{
	"ID": 1,
	"CreatedAt": "2025-04-25T22:00:12.859752-03:00",
	"UpdatedAt": "2025-04-25T22:00:12.859752-03:00",
	"DeletedAt": null,
	"artist_name": "Péricles",
	"song_name": "Até que durou",
	"musical_genre": "Pagode"
}
```

Example of 404 - Not Found 
```
{
	"Not found": "Singer not found"
}
```

----------------------------

### 4) Delete singer
- URL params: localhost:8080/singers/4
- Method: DELETE
- Request body: empty
- Response:
Example of 200 - OK
```
{
	"Data": "Singer deleted successfully"
}
```

----------------------------

### 5) Edit singer by ID
- URL params: localhost:8080/singers/2
- Method: PATCH
- Request body JSON format. At least one attribute field would be required:
```
{
	"song_name": "Farol das estrelas"
}
```  
- Response:
Example of 200 - OK 
```
{
	"ID": 2,
	"CreatedAt": "2025-04-25T22:00:39.474979-03:00",
	"UpdatedAt": "2025-04-25T22:00:39.474979-03:00",
	"DeletedAt": null,
	"artist_name": "Soweto",
	"song_name": "Farol das estrelas",
	"musical_genre": "Pagode"
}
```

----------------------------

### 6) Singer by name
- URL params: localhost:8080/singers/name/Exaltasamba
- Method: GET
- Request body: empty
- Response:
Example of 200 - OK 
```
{
	"ID": 3,
	"CreatedAt": "2025-04-29T20:59:39.757229-03:00",
	"UpdatedAt": "2025-04-29T20:59:39.757229-03:00",
	"DeletedAt": null,
	"artist_name": "Exaltasamba",
	"song_name": "Gamei",
	"musical_genre": "Pagode"
}
```

Example of 404 - Not Found 
```
{
	"Not found": "Singer not found"
}
```

----------------------------

### 7) Greeting
- URL params: localhost:8080/Dornas
- Method: GET
- Request body: empty
- Response:
Example of 200 - OK 
```
{
	"API says:": "Okay, Dornas?"
}
```


## Postgres image
```
  postgres:
    environment:
      - POSTGRES_USER=root
      - POSTGRES_PASSWORD=root
      - POSTGRES_DB=root
    ports:
      - "5432:5432"

  pgadmin-compose:
    environment:
      PGADMIN_DEFAULT_EMAIL: "gui@alura.com"
      PGADMIN_DEFAULT_PASSWORD: "123456"
    ports:
      - "54321:80"
```


## Tests
This project includes automated testing to validate essential functionalities. Below are the main tests implemented:

### ❌ TestFailed
This is a purposefully failed test to demonstrate an error message.
```
=== RUN   TestFailed
    main_test.go:18: Test failed intentionally!
--- FAIL: TestFailed (0.00s)
FAIL
exit status 1
FAIL    github.com/dornascarol/api-go-gin       1.047s
```

### ✅ TestGreetingStatusCode
Test to verify that the greeting function returns the expected status code.
```
=== RUN   TestGreetingStatusCode
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /:name                    --> github.com/dornascarol/api-go-gin/controllers.Greeting (3 handlers)
[GIN] 2025/05/11 - 16:57:50 | 200 |       556.8µs |                 | GET      "/Dornas"
--- PASS: TestGreetingStatusCode (0.06s)
PASS
```

## Project status
:construction: Application in progress.
