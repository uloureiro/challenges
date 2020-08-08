import { createRequire } from 'module';
const require = createRequire(import.meta.url);
const yargs = require("yargs");

const options = yargs
  .usage("Usage: -s <string prefix> -n <size>")
  .option("ar", { alias: "colors", describe: "A comma separetad values list containing the colors of each sock", type: "string", demandOption: true })
  .argv;

console.log(sockMerchant(options.ar.split(',')))

function sockMerchant(ar) {
  let drawer = []

  for(let i=0;i<ar.length;i++) {
    if(drawer[ar[i]])
      drawer[ar[i]]++
    else
      drawer[ar[i]] = 1
  }

  let counter = 0
  for(let i=0;i<drawer.length;i++) {
    if(drawer[i] > 1) {
      counter += Math.floor(drawer[i] / 2)
    }
  }

  return counter;
}
