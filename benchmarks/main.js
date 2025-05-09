const { performance } = require('perf_hooks');

const start = performance.now();

let i = 1;

while (i < 100000) {
    console.log(i);
    i++;
}

const end = performance.now();
console.log(`Node took: ${(end - start).toFixed(3)} ms`);
