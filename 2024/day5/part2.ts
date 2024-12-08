import { stdin, stdout } from 'node:process';

const checkRules = (rules: { [myKey: string]: number[] }, orders: number[][]) => {
  let sum: number = 0;
  let central: number = 0;

  for (let o = 0; o < orders.length; o++) {
    const order = orders[o];
    let lap = 0;
    let followRules = false;
    let followRulesAtFirst = true;
    while (!followRules) {
      followRules = true;
      lap++;
      for (let i = 0; i < order.length; i++) {
        const n = order[i];
        const previous = order.slice(0, i);

        if (!rules[String(n)]) {
          continue;
        }

        const afterRules = rules[String(n)];

        for (let a = 0; a < afterRules.length; a++) {
          const after = afterRules[a];

          if (previous.includes(after)) {

            console.log("Rule broken ", n, " ", after, " order :", order);

            const afterIndex = order.findIndex(e => e === after);

            if (afterIndex !== -1) {
              order[i] = after;
              order[afterIndex] = n;
              followRules = false;
              if (lap === 1) {
                followRulesAtFirst = false;
              }
              break;
            }
          }
        }
      }

    }
    if (!followRulesAtFirst) {
      central = order[Math.trunc(order.length / 2)]
      console.log("not follows :", order)
      console.log("central: ", central)
      sum += central
    }
  }

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
