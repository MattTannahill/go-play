# Go Play

This is some code I wrote to learn Go.

## Docker
```
$ docker build -t go-play .
$ docker run --rm -p 8080:8080 go-play
```

## Sample Usage
```
$ curl "localhost:8080?greeting=Sup&name=Matt"
```