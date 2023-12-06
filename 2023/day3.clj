(require '[clojure.string :as str])
(require '[clojure.set :as st])
(def input (str/split (slurp "input/day3.txt") #"\n"))

(def sample
  (str/split (str "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..") #"\n"))

(defn re-pos [re s]
  (loop [m (re-matcher re s)
         res {}]
    (if (.find m)
      (recur m (assoc res (.start m) (.group m)))
      res)))

(defn process-index [a]
  (let [line (first a) matches (second a)]
     (map (fn [match] [[(first match) line] (second match)]) matches)))

(def parse-digit-pos
  (fn [input]
    (->> (for [[x line] (map-indexed vector input) 
               :when (re-find #"\d" line)]
           [x (re-pos #"\d+" line)])
         (map process-index)
         (reduce concat))))

(def parse-symbol-pos
  (fn [input]
    (->> (for [[y line] (map-indexed vector input)
               [x item] (map-indexed vector line)
               :when (re-find #"[^0-9\.]" (str item))]
           [[x y] (str item)])
         (reduce concat)
         (apply array-map))))

(defn adjacent-pos [pos] (let [x (first pos) y (last pos)]
                           (set [[(+ x 1) y][(- x 1) y]
                                [x (+ y 1)][x (- y 1)]
                                [(+ x 1) (+ y 1)][(- x 1) (- y 1)]
                                [(+ x 1) (- y 1)][(- x 1) (+ y 1)]])))

(defn adjacent-to-num [key num] 
  (->> (for [i (range (count num))] 
         (adjacent-pos [(+ (first key) i) (last key)]))
       (reduce st/union)))

(defn is-adjacent-num [entry] 
  (let [pos (first entry) num (last entry)]
    (->> (adjacent-to-num pos num) 
         (some (fn [arg] (contains? symbols arg)))
         )))

(def symbols (parse-symbol-pos sample))
(def numbers (parse-digit-pos sample))

; P1
(->> (filter is-adjacent-num numbers)
     (map (comp read-string last))
     (reduce +)
     )

; -----

(def symbols
    (->> (for [[y line] (map-indexed vector input)
               [x item] (map-indexed vector line)
               :when (re-find #"\*" (str item))]
           [[x y] (str item)])
         (reduce concat)
         (apply array-map)))

(defn adjacent-gear [entry]
  (let [pos (first entry) num (last entry)]
    (->> (adjacent-to-num pos num)
         (filter (fn [arg] (contains? symbols arg)))
         (map (fn[gear] [gear (read-string num)]))
         )
    ))

(def numbers (parse-digit-pos input))

; P2
(->> (map adjacent-gear numbers)
     (filter not-empty)
     (reduce concat)
     (group-by first)
     (map second)
     (map #(map second %))
     (filter #(== (count %) 2))
     (map #(* (first %) (second %)))
     (reduce +)
     )

(adjacent-gear [[0 0] "467"])

