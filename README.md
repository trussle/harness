# Harness

Useful helpers when testing.

## Generators

Go has built in fuzz testing ([quick](https://golang.org/pkg/testing/quick/) package) 
that makes it simple to test lots of various different permutations. The issue
is that there isn't very many types that implement the `Generate` interface, so
it has limited usage. 

Hopefully some of these can remedy that issue.

```go
fn := func(a generators.ASCII) bool {
  return validASCII(a.String())
}
if err := quick.Check(fn, nil); err != nil {
  t.Fatal(err)
}
```

It is also fair to say that the `quick` package isn't not in the same vein as 
other languages, as this performs no reduction upon an error to target the 
point of failure. 


## Requests

Simple functions that perform HTTP requests that fails using the `testing.T` 
instance methods upon failure. 

Don't use these in non-test code!
