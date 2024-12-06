import { stdin, stdout } from 'node:process';

// Initialize arrays to hold the first and second columns
const firstColumn: number[] = [];
const secondColumn: number[] = [];

const compareLists = (firstColumn: [number], secondColumn: [number]) => {

  const sortedFirstColumn = firstColumn.slice().sort()
  const sortedSecondColumn = secondColumn.slice().sort()
  const diferencesSum = Array.from({ length: firstColumn.length }, (_, i) => Math.abs(sortedFirstColumn[i] - sortedSecondColumn[i])).reduce((acc, curr) => acc + curr, 0)

  return diferencesSum
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
    const [first, second] = line.trim().split(/\s+/);

    const num1 = Number(first);
    const num2 = Number(second);

    // Validate that both parts are valid numbers
    if (!isNaN(num1) && !isNaN(num2)) {
      firstColumn.push(num1);
      secondColumn.push(num2);
    } else {
      console.error(`Invalid numbers in line: "${line}"`);
    }
  }

  // Output the two arrays
  console.log('First Column:', firstColumn);
  console.log('Second Column:', secondColumn);
  console.log('sum', compareLists(firstColumn, secondColumn));
});

// Resume stdin to start reading



stdin.resume();
