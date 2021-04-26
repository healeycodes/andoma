import chess
import numpy as np
from enum import Enum
from typing import Optional


class HashEntry:
    def __init__(self, zobrist: int, best_move: Optional[chess.Move], depth: int, value: float, age: int):
        self.zobrist = zobrist
        self.best_move = best_move
        self.depth = depth
        self.value = value
        self.age = age

# Using 2^20 + 7 like in
# http://mediocrechess.blogspot.com/2007/01/guide-transposition-tables.html
table_size = 1048583

class TranspositionTable:

    def __init__(self):
        self.table = np.empty(table_size, dtype = HashEntry)

    # Replacement Strategy: Depth-Preferred + Aging
    def replace(self, hash_entry):
        index = hash_entry.zobrist % table_size
        stored_entry = self.table[index]

        if not stored_entry or hash_entry.age > stored_entry.age or hash_entry.depth > stored_entry.depth:
            np.put(self.table, index, hash_entry)

    def get(self, zobrist: int, depth: int):
        index = zobrist % table_size

        stored_entry = self.table[index]

        if stored_entry and zobrist == stored_entry.zobrist and stored_entry.depth >= depth:
            return stored_entry

        return None