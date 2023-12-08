def is_prime(num)
    return false if num < 2

    (2..Math.sqrt(num).to_i).none? { |i| (num % i).zero? }
end

File.open("checkPrime.rb", "w") do |file|
    file.write("def is_prime(num)\n")
    file.write("    return false if num < 2\n\n")
    file.write("    (2..Math.sqrt(num).to_i).none? { |i| (num % i).zero? }\n")
    file.write("end\n\n")

    file.write("def check_prime(num)\n")
    (1..1000).each do |i|
        if is_prime(i)
            file.write("    return true if num == #{i}\n")
        else
            file.write("    return false if num == #{i}\n")
        end
    end
    file.write("    return false\n")
    file.write("end\n\n")

    file.write("def main\n")
    file.write("    print 'Enter a number: '\n")
    file.write("    num = gets.chomp.to_i\n")
    file.write("    if check_prime(num)\n")
    file.write("        puts \"\#{num} is a prime number.\"\n")
    file.write("    else\n")
    file.write("        puts \"\#{num} is not a prime number.\"\n")
    file.write("    end\n")
    file.write("end\n\n")

    file.write("if __FILE__ == $0\n")
    file.write("    main\n")
    file.write("end\n")
end

puts "File written successfully."
