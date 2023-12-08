using System;
using System.IO;

class PrimeScriptGenerator
{
    static bool IsPrime(int num)
    {
        if (num < 2)
        {
            return false;
        }
        for (int i = 2; i <= Math.Sqrt(num); i++)
        {
            if (num % i == 0)
            {
                return false;
            }
        }
        return true;
    }

    static void Main()
    {
        string scriptFileName = "CheckPrime.cs";

        using (StreamWriter file = new StreamWriter(scriptFileName))
        {
            file.WriteLine("using System;");
            file.WriteLine("class PrimeChecker");
            file.WriteLine("{");
            file.WriteLine("    static bool IsPrime(int num)");
            file.WriteLine("    {");
            file.WriteLine("        if (num < 2)");
            file.WriteLine("        {");
            file.WriteLine("            return false;");
            file.WriteLine("        }");
            file.WriteLine("        for (int i = 2; i <= Math.Sqrt(num); i++)");
            file.WriteLine("        {");
            file.WriteLine("            if (num % i == 0)");
            file.WriteLine("            {");
            file.WriteLine("                return false;");
            file.WriteLine("            }");
            file.WriteLine("        }");
            file.WriteLine("        return true;");
            file.WriteLine("    }");
            file.WriteLine("");
            file.WriteLine("    static bool CheckPrime(int num)");
            file.WriteLine("    {");

            for (int i = 2; i <= 1000; i++)
            {
                if (IsPrime(i))
                {
                    file.WriteLine("        if (num == " + i + ")");
                    file.WriteLine("        {");
                    file.WriteLine("            return true;");
                    file.WriteLine("        }");
                }
            }

            file.WriteLine("        return false;");
            file.WriteLine("    }");
            file.WriteLine("");
            file.WriteLine("    static void Main()");
            file.WriteLine("    {");
            file.WriteLine("        Console.Write(\"Enter a number: \");");
            file.WriteLine("        int num = Convert.ToInt32(Console.ReadLine());");
            file.WriteLine("");
            file.WriteLine("        if (CheckPrime(num))");
            file.WriteLine("        {");
            file.WriteLine("            Console.WriteLine(num + \" is a prime number.\");");
            file.WriteLine("        }");
            file.WriteLine("        else");
            file.WriteLine("        {");
            file.WriteLine("            Console.WriteLine(num + \" is not a prime number.\");");
            file.WriteLine("        }");
            file.WriteLine("    }");
            file.WriteLine("}");
        }

        Console.WriteLine("Script file " + scriptFileName + " created.");
    }
}
