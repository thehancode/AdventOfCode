import { stdin, stdout } from 'node:process';

const wordSearch = (puzzle: string[], word: string) => {

  let center = word[1]
  let starts: number[][] = []

  puzzle.forEach(
    (line, i) => {
      line.split('').forEach((c, j) => {
        if (center === c)
          starts.push([i, j])
      });
    }
  )
  let directions = [[1, 1], [-1, 1], [-1, -1], [1, -1]]
  let matches = 0;

  starts.forEach(start => {
    let found: string[] = [];
    directions.forEach(direction => {
      let posi = start[0] + direction[0];
      let posj = start[1] + direction[1];
      if (posi >= 0 && posj >= 0 && posi < puzzle.length && posj < puzzle[0].length) {
        found.push(puzzle[posi][posj]);
      }
    });

    let join = found.join('');
    if (
      join.length === 4 &&
      (join === "MSSM" || join === "MMSS" || join === "SMMS" || join === "SSMM")
    ) {
      matches++;
    }
  });

  return matches;
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

  let searchCount = wordSearch(lines, "MAS");

  // Output the two arrays
  console.log('Count  :', searchCount);
});

// Resume stdin to start reading
stdin.resume();
