import { stdin, stdout } from 'node:process';

// Initialize arrays to hold the first and second columns
const allLevel: number[][] = [];

const checkIfSafe = (levels: number[]) => {
  for (let i = 0; i < levels.length - 1; i++) {
    if (Math.abs(levels[i] - levels[i + 1]) > 3) {
      return false
    }
  }
  return true
}
const checkIfOrdered = (levels: number[]) => {
  let everyAsc = true;
  let everyDesc = true;
  for (let i = 0; i < levels.length - 1; i++) {

    if (everyAsc && levels[i] >= levels[i + 1]) {
      everyAsc = false
    }

    if (everyDesc && levels[i] <= levels[i + 1]) {
      everyDesc = false
    }
  }
  return everyDesc || everyAsc
}

const checkAllLevels = (allLevels: number[][]) => {

  let count = 0
  allLevels.forEach(
    (levels) => {
      console.log(levels);


      if (checkIfSafe(levels) && checkIfOrdered(levels)) {
        count++
        console.log("safe");

      }
    }
  )
  return count
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

  for (const line of lines) {
    // Split each line by any whitespace (spaces, tabs, etc.)
    const levels = line.trim().split(/\s+/).map(Number);
    allLevel.push(levels)
  }

  // Output the two arrays
  console.log('Data :', allLevel);
  console.log('checkAllLevels :', checkAllLevels(allLevel));
});

// Resume stdin to start reading
stdin.resume();
