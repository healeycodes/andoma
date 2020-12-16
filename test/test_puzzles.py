import chess
import unittest
from io import StringIO
from unittest.mock import patch
from movegeneration import next_move


class TestPuzzles(unittest.TestCase):
    def test_mate_in_two_puzzles(self):
        # Siegbert Tarrasch vs. Max Kurschner
        # 1. Qg6+ hxg6 2. Bxg6#
        board = chess.Board(
            'r2qk2r/pb4pp/1n2Pb2/2B2Q2/p1p5/2P5/2B2PPP/RN2R1K1 w - - 1 0')
        move = next_move(3, board)
        self.assertEqual(move.uci(), 'f5g6')
