![Unit tests and puzzles](https://github.com/healeycodes/andoma/workflows/Unit%20tests%20and%20puzzles/badge.svg)

# ♟ Andoma
> My blog post: [Building My Own Chess Engine](https://healeycodes.com/building-my-own-chess-engine/)

<br>

A chess engine called _Andoma_ (after its authors _Andrew_ and _Roma_).

It implements:
- [Alpha-beta pruning](https://en.wikipedia.org/wiki/Alpha%E2%80%93beta_pruning) for move searching
- [Move ordering](https://www.chessprogramming.org/Move_Ordering) based off heuristics like captures and promotions
- Tomasz Michniewski's [Simplified Evaluation Function](https://www.chessprogramming.org/Simplified_Evaluation_Function) for board evaluation and piece-square tables
- A slice of the Universal Chess Interface (UCI) to allow challenges via lichess.org

<br>

An example interaction with the engine (responses are commented):

```bash
uci
# id name Andoma
# id author Andrew Healey & Roma Parramore
# uciok
position startpos moves e2e4
go
# bestmove g8f6
```

See the [UCI interface doc](https://github.com/healeycodes/andoma/blob/main/uci-interface.md) for more information on communicating with the engine.

Debug information (the number of positions searched, the time taken) is sent to stderr. The engine's response is sent to stdout.

<br>

## Lichess.org

The UCI protocol slice that's implemented by this engine means you can play it via lichess.org by using [ShailChoksi/lichess-bot](https://github.com/ShailChoksi/lichess-bot) (a bridge between Lichess API and chess engines) and a BOT account.

The engine file required by `lichess-bot` may be generated using [pyinstaller](https://www.pyinstaller.org/).

<br>

## Use it!

Start the engine with:

`python main.py`

Give it a fen:

`position fen rnbqk1nr/p1ppppbp/1p4p1/8/2P5/2Q5/PP1PPPPP/RNB1KBNR b KQkq - 0 1`

Get a move:

`go`

<hr>

```bash
$ python  main.py
position fen rnbqk1nr/p1ppppbp/1p4p1/8/2P5/2Q5/PP1PPPPP/RNB1KBNR b KQkq - 0 1
>>> position fen rnbqk1nr/p1ppppbp/1p4p1/8/2P5/2Q5/PP1PPPPP/RNB1KBNR b KQkq - 0 1
go
>>> go
>>> {'positions': 9128, 'time': 1.6894989013671875}
bestmove g7c3
```

<br>

## Tests

There are unit tests for the engine and evaluation modules, and mate-in-two/mate-in-three puzzles are being added.

`python -m unittest discover test/`
