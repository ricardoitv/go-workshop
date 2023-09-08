# Commands

```sh
go run ./...
go build ./...
```

# Testing
To compare structs in tests - https://pkg.go.dev/github.com/google/go-cmp/cmp

# Error handling 

```go
err := TheThing()
if err != nil {
  // Do something.
}
err = TheOtherThing() // no := but just = to avoid errors
if err != nil {
  // Do something.
}
```

# Dependencies

```
go get github.com/sashabaranov/go-openai
```