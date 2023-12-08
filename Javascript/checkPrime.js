const readline = require("readline");

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

function checkPrime(num) {
    for (let i = 1; i <= 1000; i++) {
        if (isPrime(i) && num === i) {
            return true;
        }
    }
    return false;
}

function main() {
    const rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout
    });

    rl.question("Enter a number: ", (num) => {
        const parsedNum = parseInt(num, 10);
        if (!isNaN(parsedNum)) {
            if (checkPrime(parsedNum)) {
                console.log(`${parsedNum} is a prime number.`);
            } else {
                console.log(`${parsedNum} is not a prime number.`);
            }
            rl.close();
        } else {
            console.log("Invalid input. Please enter a valid number.");
            rl.close();
        }
    });

    rl.on("close", () => {
        process.exit(0);
    });
}

if (typeof module !== "undefined" && module.exports) {
    module.exports = { isPrime, checkPrime, main };
}
main();
