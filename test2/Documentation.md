# submission to Rollee code challenge
---

the service is started with the following commands:
```
go mod download
```
to download all project dependencies, then:
```
go run main.go
```
or using Makefile:
```
make download
```
followed with
```
make run
```
to start the service
---

it can also be started in a docker container by using the Makefile:

```
make run-docker
```

and stopped with:
```
make stop-docker
```

---
once started, the service is available on:
```
http://localhost:8080/service
```

 sample post request:
```
http://localhost:8080/service?word=abc
```

 sample GET request:
```
http://localhost:8080/service?prefix=a
```
---

unit tests can be run using:
```
go test -v -cover ./...
```
or using Makefile:
```
make test
```

