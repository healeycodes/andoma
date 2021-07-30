from unittest.mock import patch
import chess
import unittest
from ui import render, get_move


class TestUI(unittest.TestCase):
    def test_render_displays_all_characters(self):
        """
        Test CLI rendering all chess characters
        """
        board = chess.Board(chess.STARTING_FEN)
        black_pieces = "♖♘♗♕♔♙"
        white_pieces = "♜♞♝♛♚♟"
        ranks = "12345678"
        files = "abcdefgh"

        display = render(board)
        for char in black_pieces + white_pieces + ranks + files:
            self.assertIn(char, display)

    @patch("ui.input", return_value="e2e4")
    def test_get_move(self, _):
        """
        Test CLI accepting a move via stdin
        """
        board = chess.Board(chess.STARTING_FEN)
        legal_move = get_move(board)
        self.assertIn(legal_move, list(board.legal_moves))
