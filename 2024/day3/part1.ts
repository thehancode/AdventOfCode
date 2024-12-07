
import { stdin, stdout } from 'node:process';

const generate = (number: string[]) => {
  console.log("gen: ", number);

  let ans = 0
  let current = 1
  for (const d of number.reverse()) {
    ans += Number(d) * current
    current *= 10
  }
  console.log("ans: ", ans);
  return ans
}

const processLine = (line: string) => {

  enum stage { unmatch, first, second, third };
  let current: stage = stage.unmatch
  let start = "mul("
  let counter = -1
  let firstNumber: string[] = []
  let secondNumber: string[] = []
  let sum = 0
  for (const c of line) {
    console.log(c);

    if (current === stage.unmatch) {

      if (c === start[counter + 1]) {
        console.log("matching",);
        counter++
        if (counter === start.length - 1) {
          current = stage.first
          console.log("matched",);
        }
      } else {
        counter = -1
        current = stage.unmatch
        firstNumber = []
        secondNumber = []
      }
    }

    if (current === stage.first) {
      if (!Number.isNaN(c)) {
        console.log("number");
        firstNumber.push(c)
      } else if (c === ',') {
        current = stage.second
        console.log("coma");
        continue
      } else {
        counter = -1
        current = stage.unmatch
        firstNumber = []
        secondNumber = []
      }
    }
    if (current === stage.second) {
      if (!Number.isNaN(c)) {
        console.log("second",);
        secondNumber.push(c)
      } else if (c === ')') {
        current = stage.third
      } else {
        counter = -1
        current = stage.unmatch
        firstNumber = []
        secondNumber = []
      }
    }
    if (current === stage.third) {
      console.log("generate");
      sum += generate(firstNumber) * generate(secondNumber)
      counter = -1
      current = stage.unmatch
      firstNumber = []
      secondNumber = []
    }
  }
  return sum
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
  const line = input;

  console.log("ans : ", processLine(line));
  // Output the two arrays
});

// Resume stdin to start reading
stdin.resume();
