import chess
import unittest
from io import StringIO
from unittest.mock import patch
from communication import command


class TestCommunication(unittest.TestCase):
    def test_uci_command(self):
        board = chess.Board()
        with patch('sys.stdout', new=StringIO()) as patched_output:
            command(3, board, 'uci')
            lines = patched_output.getvalue().strip().split('\n')
            self.assertEqual(len(lines), 3)

    def test_startpos_command(self):
        board = chess.Board()
        command(3, board, 'position startpos moves e2e4')
        self.assertEqual(board.fen(),
                         'rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq - 0 1')

    def test_go_command(self):
        board = chess.Board()
        with patch('sys.stdout', new=StringIO()) as patched_output:
            command(
                3, board, 'position fen 3r4/8/1R4pk/1P3p1p/3bn2P/3R2P1/6K1/3B4 b - - 0 1')
            command(3, board, 'go')

            # black bishop should take a undefended rook
            self.assertEqual(patched_output.getvalue().strip(), 'd4b6')

        board = chess.Board()
        with patch('sys.stdout', new=StringIO()) as patched_output:
            command(
                3, board, 'position fen rnbqk1nr/p1ppppbp/1p4p1/8/2P5/2Q5/PP1PPPPP/RNB1KBNR b KQkq - 0 1')
            command(3, board, 'go')

            # black will trade a bishop for a queen
            self.assertEqual(patched_output.getvalue().strip(), 'g7c3')


        board = chess.Board()
        with patch('sys.stdout', new=StringIO()) as patched_output:
            command(
                3, board, 'position fen rn1qk2r/pb1nbppp/1p2p3/2ppP3/3P4/P2B1N2/1PP1NPPP/R1BQ1RK1 b kq - 0 1')
            command(3, board, 'go')

            # black will threaten a bishop with a pawn (a very strong but not instantly obvious move)
            self.assertEqual(patched_output.getvalue().strip(), 'c5c4')