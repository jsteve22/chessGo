#!/usr/bin/env python3

def main():

  new_line_count = 4

  print('{')
  for i in range(64):
    bitmap_moves = generate_knight_psuedomoves(i)
    hex_bitmap = hex(bitmap_moves)
    if (len(hex_bitmap) != 18):
      padded_zeros = 18-len(hex_bitmap)
      hex_bitmap = '0x' + '0'*padded_zeros + hex_bitmap[2:]
    print(f'{hex_bitmap}', end=', ')
    # bitboard_print(bitmap_moves)
    # print()
    if i % new_line_count == new_line_count-1:
      print()
  print('}')
  return

def bitboard_print(bitboard):
  for i in range(64):
    p = 'x' if index_to_bitboard(i)&bitboard > 0 else 'o'
    print(f'{p}', end=' ')
    if i % 8 == 7:
      print()

def index_to_bitboard(index):
  return 1 << (index)

def generate_knight_psuedomoves(index):
  bottom_mask = 8-1
  rank = index >> 3
  file = index & bottom_mask

  bitmap_moves = 0

  for vert in [-2, 2]:
    for horz in [-1, 1]:
      next_rank = rank + vert
      next_file = file + horz
      if (next_rank < 0) or (next_rank >= 8) or (next_file < 0) or (next_file >= 8):
        continue
      bitmap_moves += index_to_bitboard( (next_rank<<3) + next_file )
  
  for horz in [-2, 2]:
    for vert in [-1, 1]:
      next_rank = rank + vert
      next_file = file + horz
      if (next_rank < 0) or (next_rank >= 8) or (next_file < 0) or (next_file >= 8):
        continue
      bitmap_moves += index_to_bitboard( (next_rank<<3) + next_file )
  return bitmap_moves

if __name__ == '__main__':
  main()