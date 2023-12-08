<?php

function isPrime($num) {
    if ($num < 2) {
        return false;
    }

    for ($i = 2; $i <= sqrt($num); $i++) {
        if ($num % $i === 0) {
            return false;
        }
    }

    return true;
}

$fileContent = '<?php

function isPrime($num) {
    if ($num < 2) {
        return false;
    }

    for ($i = 2; $i <= sqrt($num); $i++) {
        if ($num % $i === 0) {
            return false;
        }
    }

    return true;
}

function checkPrime($num) {
';

for ($i = 1; $i <= 1000; $i++) {
    if (isPrime($i)) {
        $fileContent .= "    if (\$num === $i) {\n";
        $fileContent .= "        return true;\n";
        $fileContent .= "    }\n";
    }
}

$fileContent .= '    return false;
}

function main() {
    $num = (int) readline("Enter a number: ");
    if (checkPrime($num)) {
        echo "$num is a prime number.\n";
    } else {
        echo "$num is not a prime number.\n";
    }
}

if (php_sapi_name() === "cli") {
    main();
}
';

file_put_contents("checkPrime.php", $fileContent);

echo "File written successfully.\n";
?>
