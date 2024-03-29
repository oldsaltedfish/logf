# golang logger facade

golang logger facade,  zap, logrus,and you also can implement other log lib by logf interface.

## example

```go
func main() {
    l, _ := zap.NewProduction()
    var Logger Logf = logger.NewZapLogger(l)
    Logger.Info("hello world")
}
```