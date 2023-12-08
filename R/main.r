is_prime <- function(num) {
    if (num < 2) {
        return(FALSE)
    }

    return(all(num %% (2:sqrt(num)) != 0))
}

file_path <- "checkPrime.R"
cat("is_prime <- function(num) {\n", file = file_path)
cat("  if (num < 2) {\n", file = file_path, append = TRUE)
cat("    return(FALSE)\n", file = file_path, append = TRUE)
cat("  }\n\n", file = file_path, append = TRUE)
cat("  return(all(num %% (2:sqrt(num)) != 0))\n", file = file_path, append = TRUE)
cat("}\n\n", file = file_path, append = TRUE)

cat("check_prime <- function(num) {\n", file = file_path, append = TRUE)
for (i in 1:1000) {
    # if it's prime, say return true, else say return false
    if (is_prime(i)) {
        cat(paste("  if (num == ", i, ") {\n", sep = ""), file = file_path, append = TRUE)
        cat("    return(TRUE)\n", file = file_path, append = TRUE)
        cat("  }\n\n", file = file_path, append = TRUE)
    } else {
        cat(paste("  if (num == ", i, ") {\n", sep = ""), file = file_path, append = TRUE)
        cat("    return(FALSE)\n", file = file_path, append = TRUE)
        cat("  }\n\n", file = file_path, append = TRUE)
    }
}
cat("  return(FALSE)\n", file = file_path, append = TRUE)
cat("}\n\n", file = file_path, append = TRUE)

cat("main <- function() {\n", file = file_path, append = TRUE)
cat("  num <- as.integer(readline('Enter a number: '))\n", file = file_path, append = TRUE)
cat("\n", file = file_path, append = TRUE)
cat("  if (check_prime(num)) {\n", file = file_path, append = TRUE)
cat("    cat(paste(num, 'is a prime number.'), '\\n')\n", file = file_path, append = TRUE)
cat("  } else {\n", file = file_path, append = TRUE)
cat("    cat(paste(num, 'is not a prime number.'), '\\n')\n", file = file_path, append = TRUE)
cat("  }\n", file = file_path, append = TRUE)
cat("}\n\n", file = file_path, append = TRUE)

cat("if (interactive()) {\n", file = file_path, append = TRUE)
cat("  main()\n", file = file_path, append = TRUE)
cat("}\n", file = file_path, append = TRUE)

cat("File written successfully.\n")
