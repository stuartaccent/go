# Learning GO

![Test](https://github.com/stuartaccent/go/workflows/Test/badge.svg)

A basic site with user authentication and email.

## Docker

Build and run the docker container

```bash
docker-compose build
docker-compose up
```

## Working with VS Code

1, In compose file replace:
```
command: app
```

with:
```
command: /bin/sh
tty: true
```

2, Up the container and use vscode to attach to the running container

3, Open `/go/src/app` as the workspace root

3, Run `go run app.go` to start the site

## Testing

run all tests

```bash
docker-compose exec app go test -v ./...
```

## Commands

Create a new user:
```bash
docker-compose exec app createuser -email=admin@example.com -firstname=Admin -lastname=User -password=password
```

Generate a base64 encoded random key for the auth session:
```bash
docker-compose exec app randomkey -length=32
```

## Links

Docs

* [https://golang.org/doc](https://golang.org/doc/)
* [https://gorm.io/docs](https://gorm.io/docs/)
* [http://www.gorillatoolkit.org/pkg/mux](http://www.gorillatoolkit.org/pkg/mux)
* [http://www.gorillatoolkit.org/pkg/securecookie](http://www.gorillatoolkit.org/pkg/securecookie)
* [http://www.gorillatoolkit.org/pkg/sessions](http://www.gorillatoolkit.org/pkg/sessions)
* [https://github.com/dchest/passwordreset](https://github.com/dchest/passwordreset)
* [https://github.com/dchest/authcookie](https://github.com/dchest/authcookie)
* [https://godoc.org/github.com/go-playground/validator](https://godoc.org/github.com/go-playground/validator)

Examples

* [https://gobyexample.com](https://gobyexample.com/)
* [https://github.com/avelino/awesome-go](https://github.com/avelino/awesome-go)
