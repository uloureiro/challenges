import { createRequire } from 'module';
const require = createRequire(import.meta.url);
const yargs = require("yargs");

const options = yargs
  .usage("Usage: -s <string prefix> -n <size>")
  .option("s", { alias: "path", describe: "A single string of characters that describe Gary's his path", type: "string", demandOption: true })
  .argv;

console.log(countingValleys(options.s))

function countingValleys(s) {
  let stepsArray = s.split('')
  let previousAltitude = 0
  let currentAltitude = 0
  let valleys = 0

  for (let i = 0; i < stepsArray.length; i++) {
    if (stepsArray[i] == "U") {
      currentAltitude++
    } else {
      currentAltitude--
    }

    if (currentAltitude < 0 && previousAltitude >= 0)
      valleys++

    previousAltitude = currentAltitude
  }

  return valleys
}
