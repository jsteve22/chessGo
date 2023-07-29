#!/usr/bin/env python3

def main():

  new_line_count = 4

  print('{')
  for i in range(64):
    # bitmap_moves = generate_knight_pseudomoves(i)
    # bitmap_moves = generate_king_pseudomoves(i)
    # bitmap_moves = generate_bishop_pseudomoves(i)
    # bitmap_moves = generate_rook_pseudomoves(i)
    bitmap_moves = generate_queen_pseudomoves(i)
    # bitboard_print(bitmap_moves)
    # print()
    # continue
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

def generate_knight_pseudomoves(index):
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

def generate_king_pseudomoves(index):
  bottom_mask = 8-1
  rank = index >> 3
  file = index & bottom_mask

  bitmap_moves = 0

  for vert in [-1, 0, 1]:
    for horz in [-1, 0, 1]:
      if (vert == 0) and (horz == 0):
        continue
      next_rank = rank + vert
      next_file = file + horz
      if (next_rank < 0) or (next_rank >= 8) or (next_file < 0) or (next_file >= 8):
        continue
      bitmap_moves |= index_to_bitboard( (next_rank<<3) + next_file )

  return bitmap_moves

def generate_bishop_pseudomoves(index):
  bottom_mask = 8-1
  rank = index >> 3
  file = index & bottom_mask

  bitmap_moves = 0

  for dirs in ((1, 1), (-1, 1), (1, -1), (-1, -1)):
    rd, fd = dirs
    ri, fi = rank + rd, file + fd
    while (ri >= 0) and (ri < 8) and (fi >= 0) and (fi < 8):
      bitmap_moves |= index_to_bitboard( (ri<<3) + fi )
      ri, fi = ri + rd, fi + fd

  return bitmap_moves

def generate_rook_pseudomoves(index):
  bottom_mask = 8-1
  rank = index >> 3
  file = index & bottom_mask

  bitmap_moves = 0

  for dirs in ((1, 0), (-1, 0), (0, 1), (0, -1)):
    rd, fd = dirs
    ri, fi = rank + rd, file + fd
    while (ri >= 0) and (ri < 8) and (fi >= 0) and (fi < 8):
      bitmap_moves |= index_to_bitboard( (ri<<3) + fi )
      ri, fi = ri + rd, fi + fd

  return bitmap_moves

def generate_queen_pseudomoves(index):
  return generate_rook_pseudomoves(index) | generate_bishop_pseudomoves(index)

if __name__ == '__main__':
  main()