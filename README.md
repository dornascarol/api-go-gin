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


## Tools
The <a href="https://go.dev/doc/install" target="_blank"> Go</a> programming language was installed on Windows, version 1.24.0.

The Go extension was installed in <a href="https://code.visualstudio.com/download" target="_blank"> VS Code</a>, version 0.46.1.

The program <a href="https://insomnia.rest/download" target="_blank" > Insomnia</a> was used to simulate route requests.

The container management tool <a href="https://www.docker.com/products/docker-desktop/" target="_blank" > Docker Desktop</a> was installed, using the version compatible with the operating system.


## Running the project
- To get the repository, clone it:
```
https://github.com/dornascarol/api-go-gin.git
```

- Command to check, update, and download the package manager modules (go.mod):
```
go mod tidy
```

- Command to run the project in the terminal:
```
go run main.go
```

- To stop the execution in the terminal: press `"Ctrl"` + `"C"`.
<br>  

### _Docker_

- Command to run the `docker-compose.yml` file with two service images, one for Postgres and another for pgAdmin:
```
docker-compose up
```

_Another way to run Docker_

- To find out the full name of the containers, write the command in the terminal:
```
docker ps -a
```
  
- To run the containers, write the next command in the terminal:
```
docker start gin-api-rest-pgadmin-compose-1 gin-api-rest-postgres-1
```
  
- Stop the containers with the command in the terminal:
```
docker stop gin-api-rest-pgadmin-compose-1 gin-api-rest-postgres-1
```

- URL to access pgAdmin:
```
localhost:54321
```


## Endpoints

```
http://localhost:8080/singers
```

| Method | URL                 | Description                                      |
| ------ | ------------------- | ------------------------------------------------ |
| GET    | /singers            | List all singers                                 |
| POST   | /singers            | Create a new singer                              |
| GET    | /singers/:id        | Get information of a specific singer             |
| DELETE | /singers/:id        | Delete a specific singer                         |
| PATCH  | /singers/:id        | Update information of a specific singer          |
| GET    | /singers/name/:name | Search specific singer by artist name            |


## Request body
🔸 JSON format for the POST method to create a singer
```
{
	"artist_name": "Soweto",
	"song_name": "Antes de dizer adeus",
	"musical_genre": "Pagode"
}
```

🔸 JSON format for the PATCH method to update a singer. At least one attribute field would be required
```
{
	"song_name": "Farol das estrelas"
}
```


## Response
🔸 Example of 200 - OK response for GET all singers
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
.

🔸 Example of 200 - OK response for POST create singer
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
.

🔸 Example of 200 - OK response for GET singer by ID
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

🔸 Example of 404 - Not Found response for GET singer by ID
```
{
	"Not found": "Singer not found"
}
```
.

🔸 Example of 200 - OK response for DELETE singer
```
{
	"Data": "Singer deleted successfully"
}
```
.

🔸 Example of 200 - OK response for PATCH singer by ID
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
.

🔸 Example of 200 - OK response for GET singer by name
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

🔸 Example of 404 - Not Found response for GET singer by name
```
{
	"Not found": "Singer not found"
}
```


## Postgres Image
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


## Project Status
:construction: Application in progress.
