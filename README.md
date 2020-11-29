# Andoma

A chess engine called _Andoma_.

It implements:
- Alpha-beta pruning for move searching
- Tomasz Michniewski's [Simplified Evaluation Function](https://www.chessprogramming.org/Simplified_Evaluation_Function) for board evaluation
- A slice of the Universal Chess Interface (UCI) to allow challenges via lichess.org

<br>

An example interaction with the engine (responses are commented):

```bash
uci
# id name Andoma
# id author Andrew Healey & Roma Parramore
# uciok
position fen rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
go
# bestmove h2h4
```

<br>

Start the engine with:

`go run main.go`

Build with:

`go build`

<br>

## Tests

There are some unit tests you can run with:

`go test ./...`