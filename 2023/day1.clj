(require '[clojure.string :as str])
(def input (slurp "input/day1.txt"))


(comment "P1 -----------------")

(def sample 
  (str "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet" ))

(def get-nums (fn [s] (re-seq #"[0-9]" s)))
(def make-num (fn [s] (str (first s) (last s))))
(reduce +
        (map (comp read-string make-num get-nums) (str/split input #"\n")))


(comment "P2 -----------------")

(def sample 
  (str "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen"))

(def num-re "zero|one|two|three|four|five|six|seven|eight|nine")
(def r-forward (re-pattern (str "[0-9]|" num-re)))
(def r-backward (re-pattern (str "[0-9]|" (str/reverse num-re))))

(def get-num (fn [s] [(re-find r-forward s) (str/reverse (re-find r-backward (str/reverse s)))]))

(def word-to-num {"one" 1, "two" 2, "three" 3, "four" 4, "five" 5, "six" 6, "seven" 7, "eight" 8, "nine" 9})
(def conv-to-num (fn [s] (if (re-matches #"[0-9]" s) s (word-to-num s))))
(def make-f-l-num (fn [num] (reduce str (map conv-to-num num))))

(make-f-l-num ["one", "2"])

(re-find #"zero|one|two|three|four|five|six|seven|eight|nine|[0-9]" "twone")

(reduce + 
        (map (comp read-string make-f-l-num get-num) (str/split input #"\n")))


