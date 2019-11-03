# Shorts Backend

```
go run main.go
```

Example url: [http://localhost:8080/forecast.json?lat=51.5101095&lng=-0.0932817](http://localhost:8080/forecast.json?lat=51.5101095&lng=-0.0932817)


## Running in docker

```
docker build . -t forecast
docker run -i -t -p 8080:8080 forecast
```
