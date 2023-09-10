# JWT secret generator
JWT secret generator is a small module for the Go programming language that helps to generate secret keys safely.

To start using this module you only have to use the following command in your project terminal.
```shell
go get github.com/Crazzy4999/jwt-secret-generator
```

After importing the module you will be able to generate secret keys with the following function.
```go
//64 is the length of the secret key's string

generator.GenerateSecret(64)

//example output: ]oW3qHC6;k@;>Zfff7WNOjMJuLlE3HM:dYk<`nEnE\7A85VMlqmOrQStw@jpqaCN
```
