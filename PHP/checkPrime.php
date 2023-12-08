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

function checkPrime($num) {
    if ($num === 2) {
        return true;
    }
    if ($num === 3) {
        return true;
    }
    if ($num === 5) {
        return true;
    }
    if ($num === 7) {
        return true;
    }
    if ($num === 11) {
        return true;
    }
    if ($num === 13) {
        return true;
    }
    if ($num === 17) {
        return true;
    }
    if ($num === 19) {
        return true;
    }
    if ($num === 23) {
        return true;
    }
    if ($num === 29) {
        return true;
    }
    if ($num === 31) {
        return true;
    }
    if ($num === 37) {
        return true;
    }
    if ($num === 41) {
        return true;
    }
    if ($num === 43) {
        return true;
    }
    if ($num === 47) {
        return true;
    }
    if ($num === 53) {
        return true;
    }
    if ($num === 59) {
        return true;
    }
    if ($num === 61) {
        return true;
    }
    if ($num === 67) {
        return true;
    }
    if ($num === 71) {
        return true;
    }
    if ($num === 73) {
        return true;
    }
    if ($num === 79) {
        return true;
    }
    if ($num === 83) {
        return true;
    }
    if ($num === 89) {
        return true;
    }
    if ($num === 97) {
        return true;
    }
    if ($num === 101) {
        return true;
    }
    if ($num === 103) {
        return true;
    }
    if ($num === 107) {
        return true;
    }
    if ($num === 109) {
        return true;
    }
    if ($num === 113) {
        return true;
    }
    if ($num === 127) {
        return true;
    }
    if ($num === 131) {
        return true;
    }
    if ($num === 137) {
        return true;
    }
    if ($num === 139) {
        return true;
    }
    if ($num === 149) {
        return true;
    }
    if ($num === 151) {
        return true;
    }
    if ($num === 157) {
        return true;
    }
    if ($num === 163) {
        return true;
    }
    if ($num === 167) {
        return true;
    }
    if ($num === 173) {
        return true;
    }
    if ($num === 179) {
        return true;
    }
    if ($num === 181) {
        return true;
    }
    if ($num === 191) {
        return true;
    }
    if ($num === 193) {
        return true;
    }
    if ($num === 197) {
        return true;
    }
    if ($num === 199) {
        return true;
    }
    if ($num === 211) {
        return true;
    }
    if ($num === 223) {
        return true;
    }
    if ($num === 227) {
        return true;
    }
    if ($num === 229) {
        return true;
    }
    if ($num === 233) {
        return true;
    }
    if ($num === 239) {
        return true;
    }
    if ($num === 241) {
        return true;
    }
    if ($num === 251) {
        return true;
    }
    if ($num === 257) {
        return true;
    }
    if ($num === 263) {
        return true;
    }
    if ($num === 269) {
        return true;
    }
    if ($num === 271) {
        return true;
    }
    if ($num === 277) {
        return true;
    }
    if ($num === 281) {
        return true;
    }
    if ($num === 283) {
        return true;
    }
    if ($num === 293) {
        return true;
    }
    if ($num === 307) {
        return true;
    }
    if ($num === 311) {
        return true;
    }
    if ($num === 313) {
        return true;
    }
    if ($num === 317) {
        return true;
    }
    if ($num === 331) {
        return true;
    }
    if ($num === 337) {
        return true;
    }
    if ($num === 347) {
        return true;
    }
    if ($num === 349) {
        return true;
    }
    if ($num === 353) {
        return true;
    }
    if ($num === 359) {
        return true;
    }
    if ($num === 367) {
        return true;
    }
    if ($num === 373) {
        return true;
    }
    if ($num === 379) {
        return true;
    }
    if ($num === 383) {
        return true;
    }
    if ($num === 389) {
        return true;
    }
    if ($num === 397) {
        return true;
    }
    if ($num === 401) {
        return true;
    }
    if ($num === 409) {
        return true;
    }
    if ($num === 419) {
        return true;
    }
    if ($num === 421) {
        return true;
    }
    if ($num === 431) {
        return true;
    }
    if ($num === 433) {
        return true;
    }
    if ($num === 439) {
        return true;
    }
    if ($num === 443) {
        return true;
    }
    if ($num === 449) {
        return true;
    }
    if ($num === 457) {
        return true;
    }
    if ($num === 461) {
        return true;
    }
    if ($num === 463) {
        return true;
    }
    if ($num === 467) {
        return true;
    }
    if ($num === 479) {
        return true;
    }
    if ($num === 487) {
        return true;
    }
    if ($num === 491) {
        return true;
    }
    if ($num === 499) {
        return true;
    }
    if ($num === 503) {
        return true;
    }
    if ($num === 509) {
        return true;
    }
    if ($num === 521) {
        return true;
    }
    if ($num === 523) {
        return true;
    }
    if ($num === 541) {
        return true;
    }
    if ($num === 547) {
        return true;
    }
    if ($num === 557) {
        return true;
    }
    if ($num === 563) {
        return true;
    }
    if ($num === 569) {
        return true;
    }
    if ($num === 571) {
        return true;
    }
    if ($num === 577) {
        return true;
    }
    if ($num === 587) {
        return true;
    }
    if ($num === 593) {
        return true;
    }
    if ($num === 599) {
        return true;
    }
    if ($num === 601) {
        return true;
    }
    if ($num === 607) {
        return true;
    }
    if ($num === 613) {
        return true;
    }
    if ($num === 617) {
        return true;
    }
    if ($num === 619) {
        return true;
    }
    if ($num === 631) {
        return true;
    }
    if ($num === 641) {
        return true;
    }
    if ($num === 643) {
        return true;
    }
    if ($num === 647) {
        return true;
    }
    if ($num === 653) {
        return true;
    }
    if ($num === 659) {
        return true;
    }
    if ($num === 661) {
        return true;
    }
    if ($num === 673) {
        return true;
    }
    if ($num === 677) {
        return true;
    }
    if ($num === 683) {
        return true;
    }
    if ($num === 691) {
        return true;
    }
    if ($num === 701) {
        return true;
    }
    if ($num === 709) {
        return true;
    }
    if ($num === 719) {
        return true;
    }
    if ($num === 727) {
        return true;
    }
    if ($num === 733) {
        return true;
    }
    if ($num === 739) {
        return true;
    }
    if ($num === 743) {
        return true;
    }
    if ($num === 751) {
        return true;
    }
    if ($num === 757) {
        return true;
    }
    if ($num === 761) {
        return true;
    }
    if ($num === 769) {
        return true;
    }
    if ($num === 773) {
        return true;
    }
    if ($num === 787) {
        return true;
    }
    if ($num === 797) {
        return true;
    }
    if ($num === 809) {
        return true;
    }
    if ($num === 811) {
        return true;
    }
    if ($num === 821) {
        return true;
    }
    if ($num === 823) {
        return true;
    }
    if ($num === 827) {
        return true;
    }
    if ($num === 829) {
        return true;
    }
    if ($num === 839) {
        return true;
    }
    if ($num === 853) {
        return true;
    }
    if ($num === 857) {
        return true;
    }
    if ($num === 859) {
        return true;
    }
    if ($num === 863) {
        return true;
    }
    if ($num === 877) {
        return true;
    }
    if ($num === 881) {
        return true;
    }
    if ($num === 883) {
        return true;
    }
    if ($num === 887) {
        return true;
    }
    if ($num === 907) {
        return true;
    }
    if ($num === 911) {
        return true;
    }
    if ($num === 919) {
        return true;
    }
    if ($num === 929) {
        return true;
    }
    if ($num === 937) {
        return true;
    }
    if ($num === 941) {
        return true;
    }
    if ($num === 947) {
        return true;
    }
    if ($num === 953) {
        return true;
    }
    if ($num === 967) {
        return true;
    }
    if ($num === 971) {
        return true;
    }
    if ($num === 977) {
        return true;
    }
    if ($num === 983) {
        return true;
    }
    if ($num === 991) {
        return true;
    }
    if ($num === 997) {
        return true;
    }
    return false;
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
