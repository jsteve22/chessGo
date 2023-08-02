export function FENByteToNumber(token:string) {
  switch (token) {
    case 'P':
      return 1;
    case 'N':
      return 2;
    case 'B':
      return 3;
    case 'R':
      return 4;
    case 'Q':
      return 5;
    case 'K':
      return 6;
    case 'p':
      return 1 + 8;
    case 'n':
      return 2 + 8;
    case 'b':
      return 3 + 8;
    case 'r':
      return 4 + 8;
    case 'q':
      return 5 + 8;
    case 'k':
      return 6 + 8;
  }
  return 0;
}

export function NumberToFENByte(piece:number) {
  switch (piece) {
    case 1:
      return 'P';
    case 2:
      return 'N';
    case 3:
      return 'B';
    case 4:
      return 'R';
    case 5:
      return 'Q';
    case 6:
      return 'K';
    case 1 + 8:
      return 'p';
    case 2 + 8:
      return 'n';
    case 3 + 8:
      return 'b';
    case 4 + 8:
      return 'r';
    case 5 + 8:
      return 'q';
    case 6 + 8:
      return 'k';
  }
  return '';
}

export function IndexToChessNotation(index:number) {
  if (index === -1) {
    return '-';
  }

  const FILE = index&7;
  const RANK = index>>3;
  const alpha = 'abcdefgh';

  const notation = `${alpha[FILE]}${8-RANK}`;
  return notation;
}

export function ChessNotationToIndex(notation:string) {
  const FILE = notation[0].charCodeAt(0) - 97;
  const RANK = 8-Number(notation[1]);

  const index = (RANK<<3)+FILE;
  return index;
}

export function LoadCastlingFEN(castling:string) {
  if (castling === '-') {
    return [false, false, false, false];
  }

  const castle = [false, false, false, false];
  for (let i = 0; i < castling.length; i++) {
    switch (castling[i]) {
      case 'K':
        castle[0] = true;
        break
      case 'Q':
        castle[1] = true;
        break
      case 'k':
        castle[2] = true;
        break
      case 'q':
        castle[3] = true;
        break
    }
  }
  return castle;
}

export function GenerateCastlingFEN(castle:Array<boolean>) {
  let castling = '';

  if (castle[0]) {
    castling += 'K';
  }
  if (castle[1]) {
    castling += 'Q';
  }
  if (castle[2]) {
    castling += 'k';
  }
  if (castle[3]) {
    castling += 'q';
  }

  if (!castling) {
    return '-';
  }
  return castling;
}