import { stdin, stdout } from 'node:process';


const checkRules = (rules: { [myKey: string]: number[] }, orders: number[][]) => {
  let sum = 0
  orders.forEach(
    order => {
      let followRules = true
      order.forEach(
        (n, i) => {
          if (!followRules) return
          let previous = order.slice(0, i)
          if (!rules[String(n)]) return
          rules[String(n)].forEach(
            after => {
              if (previous.includes(after)) {
                followRules = false
                return
              }
            }
          )
        }
      )
      if (followRules) {
        console.log("follows :", order)
        let central = order[Math.trunc(order.length / 2)]
        console.log("central: ", central)
        sum += central
      }
    }
  )
  return sum
}

let input = '';

stdin.setEncoding('utf-8');

stdin.on('data', (chunk) => {
  input += chunk;
});


function parseInput(input: string) {
  const lines = input.trim().split("\n");
  const groups: number[][] = [];

  let rules: { [myKey: string]: number[] } = {}
  let isGroupSection = false;

  for (const line of lines) {
    if (line.includes("|")) {
      const pair = line.split("|").map(Number);
      if (!rules[pair[0]]) {
        rules[pair[0]] = [pair[1]]
      } else {
        rules[pair[0]].push(pair[1])
      }
    } else if (line.includes(",")) {
      isGroupSection = true;
      const group = line.split(",").map(Number);
      groups.push(group);
    }
  }

  return [rules, groups]
};
// When input ends, process it
stdin.on('end', () => {
  let [rules, orders] = parseInput(input)
  console.log(rules)
  console.log(orders)
  console.log('sum  :', checkRules(rules, orders));
});

// Resume stdin to start reading
stdin.resume();
