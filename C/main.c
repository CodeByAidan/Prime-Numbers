#include <stdio.h>
#include <stdbool.h>

bool isPrime(int num)
{
    if (num < 2)
    {
        return false;
    }

    for (int i = 2; i * i <= num; i++)
    {
        if (num % i == 0)
        {
            return false;
        }
    }
    return true;
}

int main()
{
    FILE *file = fopen("checkPrime.c", "w");

    if (file == NULL)
    {
        fprintf(stderr, "Error opening file for writing.\n");
        return 1;
    }

    fprintf(file, "#ifndef CHECK_PRIME_H\n");
    fprintf(file, "#define CHECK_PRIME_H\n\n");

    fprintf(file, "#include <stdio.h>\n");
    fprintf(file, "#include <stdbool.h>\n\n");

    fprintf(file, "bool checkPrime(int num) {\n");
    for (int i = 1; i <= 1000; i++)
    {
        if (isPrime(i))
        {
            fprintf(file, "    if (num == %d) {\n", i);
            fprintf(file, "        return true;\n");
            fprintf(file, "    }\n");
        }
    }
    fprintf(file, "    return false;\n}\n\n");

    fprintf(file, "int main() {\n");
    fprintf(file, "    int num;\n");
    fprintf(file, "    printf(\"Enter a number: \");\n");
    fprintf(file, "    scanf(\"%%d\", &num);\n\n");

    fprintf(file, "    if (checkPrime(num)) {\n");
    fprintf(file, "        printf(\"%%d is a prime number.\\n\", num);\n");
    fprintf(file, "    } else {\n");
    fprintf(file, "        printf(\"%%d is not a prime number.\\n\", num);\n");
    fprintf(file, "    }\n\n");

    fprintf(file, "    return 0;\n");
    fprintf(file, "}\n");

    fprintf(file, "#endif // CHECK_PRIME_H\n");

    fclose(file);
    printf("File written successfully.\n");
    return 0;
}