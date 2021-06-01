import chess
import unittest
from evaluate import evaluate_board, move_value, check_end_game


class TestEvaluation(unittest.TestCase):
    def test_move_value_white(self):
        """
        Test move_value function playing with white pieces
        """
        board = chess.Board(
            "rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 0 1"
        )
        move = chess.Move.from_uci("e4d5")
        # pawn takes pawn
        pawn_for_pawn = move_value(board, move, check_end_game(board))
        self.assertEqual(
            pawn_for_pawn, 5, f"Pawn for pawn {pawn_for_pawn} but it should be 5"
        )

        board = chess.Board(
            "rnbqkbnr/ppp1pppp/8/3p4/2Q1P3/3P1P2/PPP1B1PP/RNB1K1NR w KQkq - 0 1"
        )
        move = chess.Move.from_uci("c4d5")
        # queen takes pawn
        queen_for_pawn = move_value(board, move, check_end_game(board))

        self.assertEqual(
            queen_for_pawn,
            -800,
            f"Queen for pawn {queen_for_pawn} but it should be -800",
        )

        board = chess.Board(
            "rnb1kbnr/ppp1pppp/8/3p1q2/4P3/3P1P2/PPP1B1PP/RNBQK1NR w KQkq - 0 1"
        )
        move = chess.Move.from_uci("e4f5")
        # pawn takes queen
        pawn_for_queen = move_value(board, move, check_end_game(board))
        self.assertEqual(
            pawn_for_queen, 790, f"Pawn for Queen {pawn_for_queen} but it should be 790"
        )

        board = chess.Board("8/4P3/2k5/8/8/3K4/8/8 w - - 0 1")
        move = chess.Move.from_uci("e7e8q")
        # pawn promotes
        pawn_promotes = move_value(board, move, check_end_game(board))
        self.assertEqual(
            pawn_promotes,
            float("inf"),
            f"Pawn promotes was {pawn_promotes} but it should be inf",
        )

        worst_to_best = list(
            sorted([pawn_promotes, pawn_for_queen, pawn_for_pawn, queen_for_pawn])
        )
        self.assertEqual(
            worst_to_best,
            [queen_for_pawn, pawn_for_pawn, pawn_for_queen, pawn_promotes],
        )

    def test_move_value_black(self):
        """
        Test move_value function playing with black pieces
        """
        board = chess.Board(
            "rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1"
        )
        move = chess.Move.from_uci("d5e4")
        # pawn takes pawn
        pawn_for_pawn = move_value(board, move, check_end_game(board))
        self.assertEqual(
            pawn_for_pawn, -5, f"Pawn for pawn {pawn_for_pawn} but it should be -5"
        )

        board = chess.Board(
            "rnb1kbnr/ppp1pppp/8/3p1q2/4P3/3P1P2/PPP1B1PP/RNBQK1NR b KQkq - 0 1"
        )
        move = chess.Move.from_uci("f5e4")
        # queen takes pawn
        queen_for_pawn = move_value(board, move, check_end_game(board))
        self.assertEqual(
            queen_for_pawn, 800, f"Queen for pawn {queen_for_pawn} but it should be 800"
        )

        board = chess.Board(
            "rnbqkbnr/ppp1pppp/8/3p4/2Q1P3/3P1P2/PPP1B1PP/RNB1K1NR b KQkq - 0 1"
        )
        move = chess.Move.from_uci("d5c4")
        # pawn takes queen
        pawn_for_queen = move_value(board, move, check_end_game(board))

        self.assertEqual(
            pawn_for_queen,
            -790,
            f"Pawn for queen {pawn_for_queen} but it should be -790",
        )

        best_to_worst = list(sorted([pawn_for_queen, pawn_for_pawn, queen_for_pawn]))
        self.assertEqual(best_to_worst, [pawn_for_queen, pawn_for_pawn, queen_for_pawn])

    def test_end_game(self):
        # starting fen
        board = chess.Board("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
        self.assertEqual(check_end_game(board), False)

        # no queens
        board = chess.Board("rnb1kbnr/ppp1pppp/8/8/8/8/PPP1PPPP/RNB1KBNR w KQkq - 0 1")
        self.assertEqual(check_end_game(board), True)

        # one queen and one minor piece each
        board = chess.Board("3k3q/3p4/8/8/8/8/4P3/Q3K3 w - - 0 1")
        self.assertEqual(check_end_game(board), True)

    def test_evaluate_board(self):
        white_down_one_pawn = chess.Board(
            "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPP1/RNBQKBNR w KQkq - 0 1"
        )

        self.assertTrue(
            evaluate_board(chess.Board()) > evaluate_board(white_down_one_pawn)
        )

        white_played_e2e4 = chess.Board(
            "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1"
        )

        self.assertTrue(
            evaluate_board(chess.Board()) < evaluate_board(white_played_e2e4)
        )

        black_played_b8c6 = chess.Board(
            "r1bqkbnr/pppppppp/2n5/8/4P3/8/PPPP1PPP/RNBQKBNR w KQkq - 1 2"
        )

        self.assertTrue(
            evaluate_board(black_played_b8c6) < evaluate_board(white_played_e2e4)
        )
