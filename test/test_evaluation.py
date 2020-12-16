import chess
import unittest
from evaluate import evaluate_board, move_value, check_end_game


class TestEvaluation(unittest.TestCase):
    def test_move_value(self):
        board = chess.Board(
            'rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1')
        move = chess.Move.from_uci('e4d5')
        # pawn takes pawn
        pawn_for_pawn = move_value(board, move, check_end_game(board))

        board = chess.Board(
            'rnb1kbnr/ppp1pppp/8/3p1q2/4P3/3P1P2/PPP1B1PP/RNBQK1NR b KQkq - 0 1')
        move = chess.Move.from_uci('f5e4')
        # queen takes pawn
        queen_for_pawn = move_value(board, move, check_end_game(board))

        board = chess.Board(
            'rnb1kbnr/ppp1pppp/8/3p1q2/4P3/3P1P2/PPP1B1PP/RNBQK1NR w KQkq - 0 1')
        move = chess.Move.from_uci('e4f5')
        # pawn takes queen
        pawn_for_queen = move_value(board, move, check_end_game(board))

        correct_order = list(
            sorted([pawn_for_queen, pawn_for_pawn, queen_for_pawn]))
        self.assertEqual(
            correct_order, [pawn_for_queen, pawn_for_pawn, queen_for_pawn])

    def test_end_game(self):
        # starting fen
        board = chess.Board(
            'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1')
        self.assertEqual(check_end_game(board), False)

        # no queens
        board = chess.Board(
            'rnb1kbnr/ppp1pppp/8/8/8/8/PPP1PPPP/RNB1KBNR w KQkq - 0 1')
        self.assertEqual(check_end_game(board), True)

        # one queen and one minor piece each
        board = chess.Board('3k3q/3p4/8/8/8/8/4P3/Q3K3 w - - 0 1')
        self.assertEqual(check_end_game(board), True)

    def test_evaluate_board(self):
        starting_fen = chess.Board(
            'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1')

        white_down_one_pawn = chess.Board(
            'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPP1/RNBQKBNR w KQkq - 0 1')

        self.assertTrue(evaluate_board(starting_fen) >
                        evaluate_board(white_down_one_pawn))
