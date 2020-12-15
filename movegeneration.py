from evaluate import evaluate_board, move_value

# heavy inspiration taken from the JavaScript chess engine https://github.com/lhartikk/simple-chess-ai
# improvements over the original:
# - move ordering
# - end game specific evaluation

def next_move(depth, board):
    '''
    What is the next best move?
    '''
    return minimax_root(depth, board, True)


def get_ordered_moves(board):
    '''
    Get legal moves.
    Attempt to sort moves by best to worst.
    Use piece values to weight captures, otherwise use positional gains.
    '''
    in_order = sorted(board.legal_moves,
                      key=lambda move: move_value(board, move))
    return list(in_order)


def minimax_root(depth, board, is_maximising_player):
    best_move = -9999
    best_move_found = None

    moves = get_ordered_moves(board)
    for move in moves:
        board.push(move)
        value = minimax(depth - 1, board, -10000,
                        10000, not is_maximising_player)
        board.pop()
        if value >= best_move:
            best_move = value
            best_move_found = move

    return best_move_found


def minimax(depth, board, alpha, beta, is_maximising_player):
    if depth == 0 or board.is_game_over():
        return -evaluate_board(board)

    if is_maximising_player:
        best_move = -9999
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
        best_move = 9999
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
