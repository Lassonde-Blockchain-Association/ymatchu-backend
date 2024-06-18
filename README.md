### YMATCHU BACKEND

### Setup `.env` file

```env
PORT=<ENTER PORT>
DB_USER=<ENTER DB_USER>
DB_PASSWORD=<ENTER DB_PASSWORD>
DB_NAME=<ENTER DB_NAME>
DB_HOST=<ENTER DB_HOST>
DB_PORT=<ENTER DB_PORT>
```

### To run the project

```golang
go mod tidy
go run cmd/main.go
```

### To run project with live reload

##### First install air
```terminal
go install github.com/air-verse/air@latest
```
##### `.air.toml` is set up so to run project use 
```terminal
air
```
More info about air [here](https://github.com/air-verse/air)
