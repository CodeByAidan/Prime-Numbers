;; Original C++ Code
;; #include <iostream>
;; #include <fstream>

;; bool isPrime(int num) {
;;     if (num < 2) {
;;         return false;
;;     }

;;     for (int i = 2; i * i <= num; i++) {
;;         if (num % i == 0) {
;;             return false;
;;         }
;;     }
;;     return true;
;; }

;; int main() {
;;     std::ofstream file("checkPrime.cpp");

;;     if (!file.is_open()) {
;;         std::cerr << "Error opening file for writing." << std::endl;
;;         return 1;
;;     }

;;     file << "#include <iostream>\n";
;;     file << "#include <fstream>\n";
;;     file << "\n";
;;     file << "bool checkPrime(int num) {\n";
;;     for (int i = 1; i <= 1000; i++) {
;;         if (isPrime(i)) {
;;             file << "    if (num == " << i << ") {\n";
;;             file << "        return true;\n";
;;             file << "    }\n";
;;         }
;;     }
;;     file << "    return false;\n";
;;     file << "}\n";
;;     file << "\n";
;;     file << "int main() {\n";
;;     file << "    int num;\n";
;;     file << "    std::cout << \"Enter a number: \";\n";
;;     file << "    std::cin >> num;\n";
;;     file << "\n";
;;     file << "    if (checkPrime(num)) {\n";
;;     file << "        std::cout << num << \" is a prime number.\\n\";\n";
;;     file << "    } else {\n";
;;     file << "        std::cout << num << \" is not a prime number.\\n\";\n";
;;     file << "    }\n";
;;     file << "\n";
;;     file << "    return 0;\n";
;;     file << "}\n";

;;     file.close();
;;     std::cout << "File written successfully." << std::endl;

;;     return 0;
;; }

;; Clojure Code
(defn is-prime? [num]
  (if (< num 2)
    false
    (loop [i 2]
      (if (<= (* i i) num)
        (if (zero? (mod num i))
          false
          (recur (inc i)))
        true))))

(defn generate-checkprime-file []
  (with-open [file (java.io.FileWriter. "checkPrime.clj")]
    (binding [*out* file]
      (println "(defn is-prime? [num]"
               "  (if (< num 2)"
               "    false"
               "    (loop [i 2]"
               "      (if (<= (* i i) num)"
               "        (if (zero? (mod num i))"
               "          false"
               "          (recur (inc i)))"
               "        true))))")
      (println)
      (println "(defn check-prime [num]"
               "  (cond"))
      (doseq [i (range 1 1001)]
        (println "    (= num " i ") (is-prime? " i ")"))
      (println "    :else false))")
      (println)
      (println "(defn -main []"
               "  (print \"Enter a number: \")"
               "  (flush)"
               "  (let [num (read-line)"
               "        num (Integer. num)]"
               "    (if (check-prime num)"
               "      (println (str num \" is a prime number.\"))"
               "      (println (str num \" is not a prime number.\")))))")
      (println)
      (println "(-main)")))

(generate-checkprime-file)
