import sys
import chess
import argparse
from movegeneration import next_move


search_depth: int


def talk():
    """
    The main input/output loop.
    This implements a slice of the UCI protocol.
    """
    global search_depth
    board = chess.Board()
    search_depth = get_depth()

    while True:
        msg = input()
        command(search_depth, board, msg)


def command(depth: int, board: chess.Board, msg: str):
    """
    Accept UCI commands and respond.
    The board state is also updated.
    """
    msg = msg.strip()
    tokens = msg.split(" ")
    while "" in tokens:
        tokens.remove("")
    if len(tokens) <= 0:
        return

    if msg == "quit":
        sys.exit()

    if msg == "uci":
        print("id name Andoma")  # Andrew/Roma -> And/oma
        print("id author Andrew Healey & Roma Parramore")
        print("option name SearchDepth type spin default 3 min 1 max 20")
        print("uciok")
        return

    if msg == "isready":
        print("readyok")
        return

    if msg == "ucinewgame":
        return

    if tokens[0] == "position":
        if len(tokens) < 2:
            return

        # Set starting position
        if tokens[1] == "startpos":
            board.reset()
            moves_start = 2
        elif tokens[1] == "fen":
            fen = " ".join(tokens[2:8])
            board.set_fen(fen)
            moves_start = 8
        else:
            return

        # Apply moves
        if len(tokens) <= moves_start or tokens[moves_start] != "moves":
            return

        for move in tokens[(moves_start+1):]:
            board.push_uci(move)

    if msg == "d":
        # Non-standard command, but supported by Stockfish and helps debugging
        print(board)
        print(board.fen())
        print(f"Depth: {depth}")

    if tokens[0] == "go":
        _move = next_move(depth, board)
        print(f"bestmove {_move}")
        return

    if tokens[0] == "setoption":
        if len(tokens) < 3 or tokens[1] != "name":
            return

        if tokens[2] == "SearchDepth":
            if len(tokens) < 5 or tokens[3] != "value":
                return
            new_depth = int(tokens[4])
            if new_depth < 1 or new_depth > 20:
                return
            global search_depth
            search_depth = new_depth
        else:
            # "SearchDepth" is the only option we support currently
            return


def get_depth() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--depth", default=3, help="provide an integer (default: 3)")
    args = parser.parse_args()
    return max([1, int(args.depth)])
