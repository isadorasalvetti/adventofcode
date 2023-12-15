(require '[clojure.string :as str])

(def sample "???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1")
(def input (slurp "input/day12.txt"))


(defn parse [sample]
  (for [line (str/split-lines sample)]
    (let [[left right] (str/split line #" ")]
    [left 
     (->> (re-seq #"\d+" right)
          (map read-string))])))

(comment
  (parse sample)
  )

(defn contains [a coll] (some #(= a %) coll))

(def matches 
  (memoize 
   (fn [left right]
     ;(print left right "\n")
     (if (empty? right) (if (contains \# left) 0 1)
         (+
         ; skip
          (if (or
               (empty? left)
               (= \# (first left))) 0
              (matches (rest left) right))
         ; match
          (let [f (first right)
                prefix (take f left)]
            (if (or (not= (count prefix) f)
                    (contains \. prefix)
                    (= \# (nth left f \.)))
              0
              (matches (drop (+ f 1) left) (rest right)))))))))

(comment
  (matches "....." '(1))
  (matches "##" '(1))
  (matches "?#..." '(2))
  (matches "" '(1))
  (matches ".?" '(1))
  (matches "???.###" '(1 1 3))
  (times-5 ["???.###" '(1 1 3)])
  )

(defn p1 [input]
  (let [rows (parse input)]
    (->> (map #(apply matches %) rows)
         (reduce +))))

(defn times-5 [[gears nums]]
  [(str/join \? (repeat 5 gears)) 
   (flatten (repeat 5 nums))])

(defn p2 [input]
     (->> input
         (parse)
         (map #(apply matches (times-5 %)))
         (reduce +)))

(p2 input)
