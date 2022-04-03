[WIP]

The log5ws package provides a simple logger wrapper dedicated to 5W input/output.

# Installation

```
go get -u github.com/Espigah/log5ws
```

# Example

```go
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	log5ws.New(&adapter.Native{}).ID("Native").Who("me").What("something").When("today").Where("here").Why("because").Info()

	log5ws.New(&adapter.Zerolog{Logger: log.Logger}).ID("Zerolog").Who("me").What("something").When("today").Where("here").Why("because").Info()

	log5ws.New(&adapter.Zap{Logger: *logger}).ID("Zap").Who("me").What("something").When("today").Where("here").Why("because").Info()

}

```

output
```
ID:Native | Who:me | What:something | When:today | Where:here | Why:because
{"level":"info","id":"Zerolog","who":"me","what":"something","when":"today","where":"here","why":"because","time":1649014839}
{"level":"info","ts":1649014839.878693,"caller":"adapter/zap.go:13","msg":"","id":"Zap","who":"me","what":"something","when":"today","where":"here","why":"because"}
```

# Links

* Palestra sobre os 5w do log - https://dev.to/solidusio/the-5-w-s-of-application-logging-77e
* Um pouco da história por trás dos 5w - https://en.wikipedia.org/wiki/Five_Ws
* Artigo que fala como fazer log direito - https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.173.2198&rep=rep1&type=pdf
* Realizando logging de maneira sistemática no Dev Eficiente  https://www.youtube.com/watch?v=gRrDUKxcqLM