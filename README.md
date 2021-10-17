### Generate GraphQL Schema
```shell
go get github.com/99designs/gqlgen/cmd@v0.14.0
go run github.com/99designs/gqlgen generate
```

### Generate Entgo Schema
```shell
go get -d entgo.io/ent/cmd/ent@v0.9.1
go run entgo.io/ent/cmd/ent generate ./ent/schema
```