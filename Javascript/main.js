const fs = require('fs');
const readline = require('readline');

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

writeStream.write('const readline = require("readline");\n\n');

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
writeStream.write('    for (let i = 1; i <= 1000; i++) {\n');
writeStream.write('        if (isPrime(i) && num === i) {\n');
writeStream.write('            return true;\n');
writeStream.write('        }\n');
writeStream.write('    }\n');
writeStream.write('    return false;\n');
writeStream.write('}\n\n');

writeStream.write('function main() {\n');
writeStream.write('    const rl = readline.createInterface({\n');
writeStream.write('        input: process.stdin,\n');
writeStream.write('        output: process.stdout\n');
writeStream.write('    });\n\n');
writeStream.write('    rl.question("Enter a number: ", (num) => {\n');
writeStream.write('        const parsedNum = parseInt(num, 10);\n');
writeStream.write('        if (!isNaN(parsedNum)) {\n');
writeStream.write('            if (checkPrime(parsedNum)) {\n');
writeStream.write('                console.log(`${parsedNum} is a prime number.`);\n');
writeStream.write('            } else {\n');
writeStream.write('                console.log(`${parsedNum} is not a prime number.`);\n');
writeStream.write('            }\n');
writeStream.write('            rl.close();\n');
writeStream.write('        } else {\n');
writeStream.write('            console.log("Invalid input. Please enter a valid number.");\n');
writeStream.write('            rl.close();\n');
writeStream.write('        }\n');
writeStream.write('    });\n\n');
writeStream.write('    rl.on("close", () => {\n');
writeStream.write('        process.exit(0);\n');
writeStream.write('    });\n');
writeStream.write('}\n\n');

writeStream.write('if (typeof module !== "undefined" && module.exports) {\n');
writeStream.write('    module.exports = { isPrime, checkPrime, main };\n');
writeStream.write('}\n');
writeStream.write('main();\n');

writeStream.end();

console.log('File written successfully.');
