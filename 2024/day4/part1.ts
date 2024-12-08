import { stdin, stdout } from 'node:process';

// Initialize arrays to hold the first and second columns
const allLevel: number[][] = [];

const checkIfSafe = (levels: number[]) => {
}
const wordSearch = (puzzle: string[], word: string) => {

  let starts: number[][] = []

  puzzle.forEach(
    (line, i) => {
      line.split('').forEach((c, j) => {
        if (word[0] === c)
          starts.push([i, j])
      });
    }
  )
  let directions = [[1, 0], [1, 1], [0, 1], [-1, 1], [-1, 0], [-1, -1], [0, -1], [1, -1]]


  let matches = 0;
  starts.forEach(
    start => {
      directions.forEach(
        direction => {
          let counter = 0
          word.split('').forEach((c, i) => {
            let posi = start[0] + direction[0] * i
            let posj = start[1] + direction[1] * i
            if (posi >= 0 && posj >= 0 && posi < puzzle[0].length && posj < puzzle.length)
              if (c === puzzle[posi][posj]) {
                counter++
              }
          })
          if (counter === word.length) {
            matches++
          }
        }
      )
    }
  )
  return matches
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

  let searchCount = wordSearch(lines, "XMAS");

  // Output the two arrays
  console.log('Count  :', searchCount);
});

// Resume stdin to start reading
stdin.resume();
