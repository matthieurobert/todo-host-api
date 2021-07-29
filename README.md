# todo-host-api

Todo-host-api is a self hosted todo list api build in Go

## Requirements

 - Docker
 - docker-compose

## Configuration

Createthe environment variables from the .env.example, for example:
```   
  POSTGRES_HOST=postgres
  POSTGRES_USER=admin
  POSTGRES_PASSWORD=root
  POSTGRES_PORT=5432
  POSTGRES_DB=todohost
  API_PORT=8000
```

## Running

Use docker-compose to run the api 
```bash
  docker-compose up -d --build
```

## Use

go to http://localhost:8000/swagger/index.html to get api doc

![Swagger doc](/assets/swagger_doc.png)
