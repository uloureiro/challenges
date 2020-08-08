import { createRequire } from 'module';
const require = createRequire(import.meta.url);
const yargs = require("yargs");

const options = yargs
 .usage("Usage: -s <string prefix> -n <size>")
 .option("s", { alias: "string prefix", describe: "The prefix of the infinite string", type: "string", demandOption: true })
 .option("n", { alias: "size", describe: "The size of the block of the string to analyze", type: "int", demandOption: true })
 .argv;

console.log(repeatedString(options.s, options.n))

// Complete the repeatedString function below.
function repeatedString(s, n) {
  let calculateTailSize = (prefix, length) => {
    let rawSize = length / prefix.length
    let decimalPart = rawSize - Math.floor(rawSize)
    return Math.round(prefix.length * decimalPart)
  }
  let calculateBodySize = (prefix, length) => {
    return Math.floor(length / prefix.length)
  }

  let regExp = new RegExp(/[^aA]/, 'g')
  let tailSize = calculateTailSize(s, n)
  let bodySize = calculateBodySize(s, n)
  let prefixLength = s.replace(regExp, '').length
  let tailLength = s.slice(0, tailSize).replace(regExp, '').length

  return (prefixLength * bodySize) + tailLength
}
