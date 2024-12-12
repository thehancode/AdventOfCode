import { stdin, stdout } from 'node:process';

// Initialize arrays to hold the first and second columns
const marchCount = (puzzleString: string[]) => {
  let puzzle: string[][] = puzzleString.map(line => line.split(''))

  console.log("Puzzle : ");
  puzzle.forEach(
    line => console.log(line.join(''))
  )

  let start: number[] = [0, 0];
  puzzle.forEach(
    (line, i) => {
      line.forEach((c, j) => {
        if (c === "^")
          start = [i, j]
      });
    }
  )
  let directions: number[][] = [[-1, 0], [0, 1], [1, 0], [0, -1],]
  let directionIdx = 0
  let currentDirection = directions[directionIdx]
  let counter: number = 0

  let currentPos = start
  let nextPos = [currentPos[0] + currentDirection[0], currentPos[1] + currentDirection[1]]
  console.log("Pos :", currentPos);
  /*
   *
    while (next if not out of bounds) {
      while (next is not obstacle)
      current = next and put # if new
        add step
    change direction
    }
   * */


  const checkPosInBounds = (cpos, cpuzzle) => (cpos[0] >= 0 && cpos[1] >= 0 && cpos[0] < cpuzzle[0].length && cpos[1] < cpuzzle.length)
  while (checkPosInBounds(nextPos, puzzle)) {

    while (puzzle[nextPos[0]][nextPos[1]] !== '#' && checkPosInBounds(nextPos, puzzle)) {

      if (puzzle[currentPos[0]][currentPos[1]] !== 'x') {
        puzzle[currentPos[0]][currentPos[1]] = 'x'
        counter++
      }
      currentPos = nextPos
      nextPos = [currentPos[0] + currentDirection[0], currentPos[1] + currentDirection[1]]
    }
    if (checkPosInBounds(nextPos, puzzle)) {
      break
    }
    currentDirection = directions[(directionIdx++) % directions.length]
    nextPos = [currentPos[0] + currentDirection[0], currentPos[1] + currentDirection[1]]
    console.log("Pos :", currentPos);
    console.log("dir :", currentDirection);
  }

  console.log("end : ");
  puzzle.forEach(
    line => console.log(line.join(''))
  )
  return counter
}


// Collect all input data
let input = '';

// Set encoding to UTF-8
stdin.setEncoding('utf-8');

// Listen for data chunks
stdin.on('data', (chunk) => {
  input += chunk;
});

// When input ends, process it
stdin.on('end', () => {
  // Split input into lines, handling both Unix and Windows line endings
  const lines = input.trim().split(/\r?\n/);


  // Output the two arrays
  console.log('Count  :', marchCount(lines));
});

// Resume stdin to start reading
stdin.resume();
