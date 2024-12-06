import { stdin, stdout } from 'node:process';

// Initialize arrays to hold the first and second columns
const firstColumn: number[] = [];
const secondColumn: number[] = [];

const compareLists = (firstColumn: [number], secondColumn: [number]) => {

  const countList1 = makeCount(firstColumn)
  const countList2 = makeCount(secondColumn)

  console.log("count 1 : ", countList1);
  console.log("count 2 : ", countList2);

  let sum = 0
  for (const num in countList1) {
    if (countList2[num]) {
      sum += countList1[num] * countList2[num] * Number(num)
    }
  }
  return sum
}

const makeCount = (list: number[]) => {
  const count = {}

  list.forEach(
    element => {
      if (count[element]) {
        count[element]++
      } else {
        count[element] = 1
      }
    }
  );
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
