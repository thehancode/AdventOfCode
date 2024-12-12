import { stdin, stdout } from 'node:process';

// Initialize arrays to hold the first and second columns
const countMarchSteps = (originalPuzzle: string[], firstTime: boolean) => {
  let puzzleGrid: string[][] = originalPuzzle.map(line => line.split(''))


  let startPosition: number[] = [0, 0];
  puzzleGrid.forEach(
    (line, i) => {
      line.forEach((c, j) => {
        if (c === "^")
          startPosition = [i, j]
      });
    }
  )
  let DIRECTIONS: number[][] = [[-1, 0], [0, 1], [1, 0], [0, -1],]
  let directionIdx = 0
  let currentDirection = DIRECTIONS[directionIdx]
  let counter: number = 1
  puzzleGrid[startPosition[0]][startPosition[1]] = 'x'

  let currentPos = startPosition
  let nextPos = [currentPos[0] + currentDirection[0], currentPos[1] + currentDirection[1]]
  console.log("Pos :", currentPos);
  let viewedBlocks = new Set()
  let visitedPositions: number[][] = []
  let currentBlock = ''

  const checkPosInBounds = (cpos, cpuzzle) => (cpos[0] >= 0 && cpos[1] >= 0 && cpos[0] < cpuzzle[0].length && cpos[1] < cpuzzle.length)
  while (checkPosInBounds(nextPos, puzzleGrid)) {

    while (checkPosInBounds(nextPos, puzzleGrid) && puzzleGrid[nextPos[0]][nextPos[1]] !== '#') {
      currentPos = nextPos
      nextPos = [currentPos[0] + currentDirection[0], currentPos[1] + currentDirection[1]]

      if (puzzleGrid[currentPos[0]][currentPos[1]] !== 'x') {
        visitedPositions.push(currentPos)
        puzzleGrid[currentPos[0]][currentPos[1]] = 'x'
        counter++
      }

    }
    if (!checkPosInBounds(nextPos, puzzleGrid)) {
      break
    }
    currentBlock = `${currentPos[0]},${currentPos[1]},${nextPos[0]},${nextPos[1]}`

    if (viewedBlocks.has(currentBlock)) {

      return -1
    }
    viewedBlocks.add(currentBlock)
    currentDirection = DIRECTIONS[(++directionIdx) % DIRECTIONS.length]
    nextPos = [currentPos[0] + currentDirection[0], currentPos[1] + currentDirection[1]]
  }

  if (!firstTime)
    return counter
  let countLoops = 0
  visitedPositions.forEach(
    (step) => {
      let puzzleNewObstacle = structuredClone(originalPuzzle)
      puzzleNewObstacle[step[0]] = puzzleNewObstacle[step[0]].substring(0, step[1]) + '#' + puzzleNewObstacle[step[0]].substring(step[1] + 1);
      if (countMarchSteps(puzzleNewObstacle, false) === -1)
        countLoops++
    }
  )
  console.log('count loops:', countLoops);
  return 1


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
  console.log('Count  :', countMarchSteps(lines, true));
});

// Resume stdin to start reading
stdin.resume();
