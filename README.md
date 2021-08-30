# Deferrer

A tiny Go module to defer defers in your unit tests to outer scopes, for
instance, when resources span multiple "By" blocks in
[Ginkgo](https://github.com/onsi/ginkgo).

## Usage

An illustrative example when using [Ginkgo](https://github.com/onsi/ginkgo):

```go
It("is a multi-By test", func(){
    outer := Deferrer{}
     defer outer.Cleanup()

     var f *os.File // ...a By-spanning resource
     By("allocates some resource", func(){
         f, _ := os.Open("foobar")
         // f.Close() returns something, so we need to wrap it.
         outer.Defer(func(){_ = f.Close()})
     })

     By("uses some resource", func(){
         var b byte[256]
           _, _ = f.Read(b)
     })
})
```

## ⚖️ Copyright and License

`deferrer` is Copyright 2021 Harald Albrecht, and licensed under the Apache
License, Version 2.0.
