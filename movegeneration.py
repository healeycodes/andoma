from typing import Dict, List, Any
import chess
import chess.polyglot
import sys
import time
from evaluate import evaluate_board, move_value, check_end_game
from transposition_table import HashEntry, TranspositionTable

debug_info: Dict[str, Any] = {}


def next_move(
    depth: int,
    board: chess.Board,
    transposition_table: TranspositionTable = None,
    debug=True,
) -> chess.Move:
    """
    What is the next best move?
    """
    debug_info.clear()
    debug_info["nodes"] = 0
    t0 = time.time()
    if not transposition_table:
        transposition_table = TranspositionTable()
    move = minimax_root(depth, board, transposition_table)

    debug_info["time"] = time.time() - t0
    if debug == True:
        print(f">>> {debug_info}", file=sys.stderr)
    return move


def get_ordered_moves(
    board: chess.Board, only_capture: bool = False
) -> List[chess.Move]:
    """
    Get legal moves.
    Attempt to sort moves by best to worst.
    Use piece values (and positional gains/losses) to weight captures.
    """
    end_game = check_end_game(board)

    def orderer(move):
        return move_value(board, move, end_game)

    in_order = sorted(
        board.legal_moves, key=orderer, reverse=(board.turn == chess.WHITE)
    )
    if only_capture:
        return [move for move in in_order if board.is_capture(move)]
    return list(in_order)


def minimax_root(
    depth: int, board: chess.Board, transposition_table: TranspositionTable
) -> chess.Move:
    # White always wants to maximize (and black to minimize)
    # the board score according to evaluate_board()
    maximize = board.turn == chess.WHITE
    best_move = -float("inf")
    if not maximize:
        best_move = float("inf")

    moves = get_ordered_moves(board)
    best_move_found = moves[0]

    for move in moves:
        board.push(move)
        # Checking if draw can be claimed at this level, because the threefold repetition check
        # can be expensive. This should help the bot avoid a draw if it's not favorable
        # https://python-chess.readthedocs.io/en/latest/core.html#chess.Board.can_claim_draw
        if board.can_claim_draw():
            value = 0.0
        else:
            value = minimax(
                depth - 1,
                board,
                -float("inf"),
                float("inf"),
                not maximize,
                transposition_table,
            )
        board.pop()
        if maximize and value >= best_move:
            best_move = value
            best_move_found = move
        elif not maximize and value <= best_move:
            best_move = value
            best_move_found = move

    return best_move_found


# https://www.chessprogramming.org/Quiescence_Search
def quiesce(board: chess.Board, alpha: float, beta: float, depth: int = 1):
    stand_pat = evaluate_board(board)
    if stand_pat >= beta:
        return beta
    if alpha < stand_pat:
        alpha = stand_pat
    if depth > 0:
        moves = get_ordered_moves(board, True)
        for move in moves:
            board.push(move)
            score = quiesce(board, alpha, beta, depth - 1)
            board.pop()
            if score >= beta:
                return beta
            if score > alpha:
                alpha = score
    return alpha


def minimax(
    depth: int,
    board: chess.Board,
    alpha: float,
    beta: float,
    is_maximising_player: bool,
    transposition_table: TranspositionTable,
) -> float:
    debug_info["nodes"] += 1

    if board.is_checkmate():
        # The previous move resulted in checkmate
        return -float("inf") if is_maximising_player else float("inf")
    # When the game is over and it's not a checkmate it's a draw
    # In this case, don't evaluate. Just return a neutral result: zero
    elif board.is_game_over():
        return 0

    if depth == 0:
        return quiesce(board, alpha, beta)

    # Transposition Table
    # https://www.chessprogramming.org/Transposition_Table
    zobrist = chess.polyglot.zobrist_hash(board)
    stored_position = transposition_table.get(zobrist, depth)
    if stored_position is not None and stored_position.depth >= depth:
        # print('-- stored ==')
        return stored_position.value
    best_move = None

    if is_maximising_player:
        move_value = -float("inf")
        moves = get_ordered_moves(board)
        for move in moves:
            board.push(move)
            best_move = move
            move_value = max(
                move_value,
                minimax(
                    depth - 1,
                    board,
                    alpha,
                    beta,
                    not is_maximising_player,
                    transposition_table,
                ),
            )
            board.pop()
            alpha = max(alpha, move_value)
            if beta <= alpha:
                break
    else:
        move_value = float("inf")
        moves = get_ordered_moves(board)
        for move in moves:
            board.push(move)
            best_move = move
            move_value = min(
                move_value,
                minimax(
                    depth - 1,
                    board,
                    alpha,
                    beta,
                    not is_maximising_player,
                    transposition_table,
                ),
            )
            board.pop()
            beta = min(beta, move_value)
            if beta <= alpha:
                break
    new_entry = HashEntry(zobrist, best_move, depth, move_value, board.fullmove_number)
    transposition_table.replace(new_entry)
    return move_value
