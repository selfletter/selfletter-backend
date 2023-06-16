# selfletter-backend

### Building:
You should have a working [Go 1.20+ environment](https://go.dev/dl/).
```
git clone https://github.com/selfletter/selfletter-backend
cd selfletter-backend
go get -d ./...
go build cmd/main.go
```

### Running:
You should have a running PostgreSQL database.

Adjust `config.json` to your settings. Refer to [docs/config.md](docs/config.md).

Then launch the executable that you built.

For api documentation refer to [docs/api/README.md](docs/api/README.md).


### Contributing:
Contributions are welcome! Both issues and pull requests.

### Licensing:
This software is licensed with [AGPL-3.0](LICENSE.txt) license.