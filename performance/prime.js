function isPrime(n) {
    if (n <= 1) {
        return false;
    }
    if (n <= 3) {
        return true;
    }
    if (n % 2 === 0 || n % 3 === 0) {
        return false;
    }
    let i = 5;
    while (i * i <= n) {
        if (n % i === 0 || n % (i + 2) === 0) {
            return false;
        }
        i += 6;
    }
    return true;
}

function jsNthPrime(n) {
    let count = 0;
    let num = 2;
    while (true) {
        if (isPrime(num)) {
            count++;
            if (count === n) {
                return num;
            }
        }
        num++;
    }
}


function testGo(n) {
    console.log("running in Go")
    console.time("Took");
    console.log(nthPrime(n))
    console.timeEnd("Took");
    console.log("completed in Go")
}


function testJS(n) {
    console.log("running in JS")
    console.time("Took");
    console.log(jsNthPrime(n))
    console.timeEnd("Took");
    console.log("completed in JS")
}

function compare(n) {
    console.log("Find the " + n + "th prime number")
    testGo(n)
    console.log("========================")
    testJS(n)
}
