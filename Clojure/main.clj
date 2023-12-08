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