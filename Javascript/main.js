const fs = require('fs');

function isPrime(num) {
    if (num < 2) {
        return false;
    }

    for (let i = 2; i <= Math.sqrt(num); i++) {
        if (num % i === 0) {
            return false;
        }
    }

    return true;
}

const writeStream = fs.createWriteStream('checkPrime.js');

writeStream.write('function isPrime(num) {\n');
writeStream.write('    if (num < 2) {\n');
writeStream.write('        return false;\n');
writeStream.write('    }\n\n');
writeStream.write('    for (let i = 2; i <= Math.sqrt(num); i++) {\n');
writeStream.write('        if (num % i === 0) {\n');
writeStream.write('            return false;\n');
writeStream.write('        }\n');
writeStream.write('    }\n\n');
writeStream.write('    return true;\n');
writeStream.write('}\n\n');

writeStream.write('function checkPrime(num) {\n');
for (let i = 1; i <= 1000; i++) {
    if (isPrime(i)) {
        writeStream.write(`    if (num === ${i}) {\n`);
        writeStream.write('        return true;\n');
        writeStream.write('    }\n');
    }
}
writeStream.write('    return false;\n');
writeStream.write('}\n\n');

writeStream.write('function main() {\n');
writeStream.write('    const num = parseInt(prompt("Enter a number: "), 10);\n');
writeStream.write('    if (checkPrime(num)) {\n');
writeStream.write('        console.log(`${num} is a prime number.`);\n');
writeStream.write('    } else {\n');
writeStream.write('        console.log(`${num} is not a prime number.`);\n');
writeStream.write('    }\n');
writeStream.write('}\n\n');

writeStream.write('if (typeof module !== "undefined" && module.exports) {\n');
writeStream.write('    module.exports = { isPrime, checkPrime, main };\n');
writeStream.write('}\n');

writeStream.end();

console.log('File written successfully.');
