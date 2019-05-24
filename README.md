# go-ctxwait

Wait on a context without wasting a goroutine per context.


## Usage

```go
ctxwait.OnDone(ctx, func() {
  // do something.
})
```
