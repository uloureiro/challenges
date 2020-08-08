import { createRequire } from 'module';
const require = createRequire(import.meta.url);
const yargs = require("yargs");

const options = yargs
 .usage("Usage: -s <clouds>")
 .option("c", { alias: "clouds", describe: "A comma separated list of the sequence of thunderheads (1) or cumulus (0) clouds", type: "string", demandOption: true })
 .argv;

console.log(jumpingOnClouds(options.c.split(',')))

function jumpingOnClouds(c) {
  let currentPosition = 0
  let steps = 0

  let isValidPosition = (next, clouds) => {
    return (next <= clouds.length - 1) && clouds[next] == 0
  }
  let calculateJump = (current, clouds) => {
    if(isValidPosition(current + 2, clouds)) {
      return 2
    }

    if(isValidPosition(current + 1, clouds)) {
      return 1
    }

    return 0
  }
  while(currentPosition < c.length - 1) {
    let nextPosition = currentPosition + calculateJump(currentPosition, c)
    if(nextPosition > currentPosition) {
      steps++
      currentPosition = nextPosition
    } else { break }
  }

  return steps
}
