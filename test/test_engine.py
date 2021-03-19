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

    def test_go_command_black(self):
        '''
        Test go command with Andoma playing with black pieces
        '''
        board = chess.Board()
        with patch('sys.stdout', new=StringIO()) as patched_output:
            command(
                3, board, 'position fen 3r4/8/1R4pk/1P3p1p/3bn2P/3R2P1/6K1/3B4 b - - 0 1')
            command(3, board, 'go')

            # black bishop should take a undefended rook
            self.assertEqual(
                patched_output.getvalue().strip(), 'bestmove d4b6')

        board = chess.Board()
        with patch('sys.stdout', new=StringIO()) as patched_output:
            command(
                3, board, 'position fen rnbqk1nr/p1ppppbp/1p4p1/8/2P5/2Q5/PP1PPPPP/RNB1KBNR b KQkq - 0 1')
            command(3, board, 'go')

            # black will trade a bishop for a queen
            self.assertEqual(
                patched_output.getvalue().strip(), 'bestmove g7c3')

        board = chess.Board()
        with patch('sys.stdout', new=StringIO()) as patched_output:
            command(
                3, board, 'position fen rn1qk2r/pb1nbppp/1p2p3/2ppP3/3P4/P2B1N2/1PP1NPPP/R1BQ1RK1 b kq - 0 1')
            command(3, board, 'go')

            # black will threaten a bishop with a pawn (a very strong but not instantly obvious move)
            self.assertEqual(
                patched_output.getvalue().strip(), 'bestmove c5c4')

    def test_go_command_white(self):
        '''
        Test go command with Andoma playing with white pieces
        '''
        board = chess.Board()
        with patch('sys.stdout', new=StringIO()) as patched_output:
            command(
                3, board, 'position fen 6r1/8/R5pk/1P3p1p/3bn2P/1B3RP1/6K1/8 w - - 0 1')
            command(3, board, 'go')

            # white bishop should take a undefended rook
            self.assertEqual(
                patched_output.getvalue().strip(), 'bestmove b3g8')

        board = chess.Board()
        with patch('sys.stdout', new=StringIO()) as patched_output:
            command(
                3, board, 'position fen rnb1kbnr/p1ppppqp/1p4p1/8/2P5/1P6/PB1PPPPP/RN2KBNR w KQkq - 0 1')
            command(3, board, 'go')

            # white will trade a bishop for a queen
            self.assertEqual(
                patched_output.getvalue().strip(), 'bestmove b2g7')

        board = chess.Board()
        with patch('sys.stdout', new=StringIO()) as patched_output:
            command(
                3, board, 'position fen r2qkb1r/pppn1pp1/2n1b2p/4p3/3pPP2/3P2P1/PPPBN1BP/R2QK1NR w KQkq - 0 1')
            command(3, board, 'go')

            # white will threaten a bishop with a pawn (a very strong but not instantly obvious move)
            self.assertEqual(
                patched_output.getvalue().strip(), 'bestmove f4f5')

    def test_draw(self):
        '''
        Test go command with Andoma on the verge of drawing due to threefold repetition
        '''
        board = chess.Board()
        with patch('sys.stdout', new=StringIO()) as patched_output:
            command(
                3, board, 'position startpos moves c2c4 d7d6 d1a4 c8d7 a4a5 b8c6 a5b5 a7a6 b5b7 c6e5 b1c3 e5c4 g1f3 d7c8 b7a8 g8f6 a8c6 c8d7 c6c4 d6d5 c4a6 e7e6 f3e5 f8d6 e5d7 d8d7 a6a8 d7d8 a8c6 d8d7 c6a8 d7d8 a8c6 d8d7')
            command(3, board, 'go')

            # Bot is in a favorable position, should avoid threefold repetition
            self.assertNotEqual(
                patched_output.getvalue().strip(), 'bestmove c6a8')
