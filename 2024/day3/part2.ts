
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
  let enabled = true;
  let current: stage = stage.unmatch;
  const start = "mul(";
  const startdo = "do()";
  const startdont = "don't()";
  let counter = -1;
  let counterdo = -1;
  let counterdont = -1;

  let firstNumber: string[] = [];
  let secondNumber: string[] = [];
  let sum = 0;

  let ans = 0;
  function generate(arr: string[]): number {
    return parseInt(arr.join(''), 10);
  }

  for (const c of line) {
    console.log("c:", c);

    if (c === startdo[counterdo + 1]) {
      counterdo++
      if (counterdo === startdo.length - 1) {
        enabled = true
        console.log("abling");
      }
    } else {
      counterdo = -1;
    }


    if (c === startdont[counterdont + 1]) {
      counterdont++
      if (counterdont === startdont.length - 1) {
        enabled = false
        console.log("disabling");

      }
    } else {
      counterdont = -1;
    }

    if (!enabled) {
      continue
    }

    if (current === stage.unmatch) {
      if (c === start[counter + 1]) {
        counter++;
        if (counter === start.length - 1) {
          console.log("mul");
          current = stage.first;
        }
      } else {
        current = stage.unmatch;
        firstNumber = [];
        secondNumber = [];
      }
    } else if (current === stage.first) {
      if (c >= '0' && c <= '9') {
        firstNumber.push(c);
      } else if (c === ',') {
        if (firstNumber.length === 0) {
          counter = -1;
          current = stage.unmatch;
          firstNumber = [];
          secondNumber = [];
        } else {
          console.log("sec ,");
          current = stage.second;
        }
      } else {
        counter = -1;
        current = stage.unmatch;
        firstNumber = [];
        secondNumber = [];
      }
    } else if (current === stage.second) {
      if (c >= '0' && c <= '9') {
        secondNumber.push(c);
      } else if (c === ')') {
        if (secondNumber.length === 0) {
          counter = -1;
          current = stage.unmatch;
          firstNumber = [];
          secondNumber = [];
        } else {
          console.log("third ,");
          current = stage.third;
          console.log("1:", firstNumber, " 2:", secondNumber);

          ans = generate(firstNumber) * generate(secondNumber);
          sum += ans

          console.log("a:", ans);
          counter = -1;
          current = stage.unmatch;
          firstNumber = [];
          secondNumber = [];
        }
      } else {
        counter = -1;
        current = stage.unmatch;
        firstNumber = [];
        secondNumber = [];
      }
    }
  }

  return sum;
};

let input = '';

stdin.setEncoding('utf-8');

stdin.on('data', (chunk) => {
  input += chunk;
});

stdin.on('end', () => {
  const line = input;
  console.log("ans : ", processLine(line));
});

stdin.resume();
