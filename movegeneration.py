import sys
import time
from evaluate import evaluate_board, move_value, check_end_game

debug = {}


def next_move(depth, board):
    '''
    What is the next best move?
    '''
    debug.clear()
    debug['positions'] = 0
    t0 = time.time()

    move = minimax_root(depth, board, True)

    debug['time'] = time.time() - t0
    print(f'>>> {debug}', file=sys.stderr)
    return move


def get_ordered_moves(board):
    '''
    Get legal moves.
    Attempt to sort moves by best to worst.
    Use piece values (and positional gains/losses) to weight captures.
    '''
    end_game = check_end_game(board)

    def orderer(move):
        return move_value(board, move, end_game)

    in_order = sorted(board.legal_moves, key=orderer, reverse=True)
    return list(in_order)


def minimax_root(depth, board, is_maximising_player):
    best_move = -float('inf')
    best_move_found = None

    moves = get_ordered_moves(board)
    for move in moves:
        board.push(move)
        value = minimax(depth - 1, board, -float('inf'),
                        float('inf'), not is_maximising_player)
        board.pop()
        if value >= best_move:
            best_move = value
            best_move_found = move

    return best_move_found


def minimax(depth, board, alpha, beta, is_maximising_player):
    debug['positions'] += 1

    if depth == 0 or board.is_game_over():
        return -evaluate_board(board)

    if is_maximising_player:
        best_move = -float('inf')
        moves = get_ordered_moves(board)
        for move in moves:
            board.push(move)
            best_move = max(best_move, minimax(
                depth - 1, board, alpha, beta, not is_maximising_player))
            board.pop()
            alpha = max(alpha, best_move)
            if beta <= alpha:
                return best_move
        return best_move
    else:
        best_move = float('inf')
        moves = get_ordered_moves(board)
        for move in moves:
            board.push(move)
            best_move = min(best_move, minimax(
                depth - 1, board, alpha, beta, not is_maximising_player))
            board.pop()
            beta = min(beta, best_move)
            if beta <= alpha:
                return best_move
        return best_move
