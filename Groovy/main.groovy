def is_prime(num) {
    if (num < 2) {
        return false
    }

    (2..Math.sqrt(num).toInteger()).each { i ->
        if (num % i == 0) {
            return false
        }
    }
    return true
}

def writeToFile() {
    def fileContent = """
        def is_prime(num) {
            if (num < 2) {
                return false
            }

            (2..Math.sqrt(num).toInteger()).each { i ->
                if (num % i == 0) {
                    return false
                }
            }
            return true
        }

        def check_prime(num) {
            if (is_prime(num)) {
                return true
            }
            return false
        }

        def main() {
            print "Enter a number: "
            def num = System.console().readLine().trim().toInteger()
            if (check_prime(num)) {
                println "\$num is a prime number."
            } else {
                println "\$num is not a prime number."
            }
        """
        // then write literally using a for-loop all the numbers from 1-1001, if num == {} return true, return false:
        (1..1001).each { if (is_prime(it)) {
            fileContent += """
                if (num == $it) {
                    return true
                }
            """
            }
        }

    fileContent +=
    """
            return false
        }
        if (this.class.name == 'Script1') {
            main()
        }
        main()
    """

    new File("checkPrime.groovy").text = fileContent
}

writeToFile()
println "File written successfully."
