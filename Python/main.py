def is_prime(num):
    if num < 2:
        return False

    return all(num % i != 0 for i in range(2, int(num**0.5) + 1))

with open("checkPrime.py", "w") as file:
    file.write("def is_prime(num):\n")
    file.write("    if num < 2:\n")
    file.write("        return False\n\n")
    file.write("    for i in range(2, int(num**0.5) + 1):\n")
    file.write("        if num % i == 0:\n")
    file.write("            return False\n")
    file.write("    return True\n\n")

    file.write("def check_prime(num):\n")
    for i in range(1, 1001):
        if is_prime(i):
            file.write(f"    if num == {i}:\n")
            file.write("        return True\n")
    file.write("    return False\n\n")

    file.write("def main():\n")
    file.write("    num = int(input('Enter a number: '))\n")
    file.write("    if check_prime(num):\n")
    file.write("        print(f'{num} is a prime number.')\n")
    file.write("    else:\n")
    file.write("        print(f'{num} is not a prime number.')\n\n")

    file.write("if __name__ == '__main__':\n")
    file.write("    main()\n")

print("File written successfully.")
