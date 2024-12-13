productRegister-GoLang-restAPI

A Rest API in GoLang, you can register a new product, update a product and consult by id or not.
Using DDD architecture, the system its already configured to be used in Docker, Containing a postgres container for the database, and the main project.
Using the libraries gin-gonic/gin, net/http, strconv.

To run, make sure that you have GoLang in your machine and run the follow commands:
> [!TIP]
>in that case project-name = main.go

### Projet init 
```
go mod init project-name
```
### Install the libraries 
```
go get github.com/gin-gonic/gin
go get strconv
got get net/http
```

### Running the project (Navigate to the cmd file)
> [!IMPORTANT]
> Remember to run at least the postgres container.
```
cd cmd
go run project-name
```


With docker installed, open your terminal and navigate to the origin of the project (where contains the docker-compose file).
Run the following commands bellow:

### Building the api image 
```
docker build -t go-api .
```
### Running docker compose
```
docker-compose up -d
```

Now the app will run in background, if you want to stop the service just open your terminal again, navigate to the origin of the project and run the follow:

### Stopping the application 
```
docker-compose down
```
