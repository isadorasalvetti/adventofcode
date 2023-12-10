
(require '[clojure.string :as str])
(def input (slurp "input/day9.txt"))

(def sample "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45")

(defn parse [input] (->> input
                         (#(str/split % #"\n"))
                         (map #(re-seq #"-?\d+" %))
                         (map #(map read-string %))))


(defn pyramid-steps [prev-steps nums] 
  (let [new-nums (->> nums
                      (partition 2 1)
                      (map #(apply - (reverse %))))]
  (if (every? #(= 0 %) new-nums)
    (reverse prev-steps)
    (pyramid-steps (conj prev-steps new-nums) new-nums))))

(defn make-prediction [num next-nums]
  (if (empty? next-nums) 
    num
    (make-prediction (+ num (last (first next-nums))) 
                     (rest next-nums))) 
  )

(make-prediction 0 '((3 3 3 3 3) (0 3 6 9 12 15)))

(->> input
     (parse)
     (map #(pyramid-steps [%] %))
     (map #(make-prediction 0 %))
     (reduce +)
     )

; P2

(defn make-prediction-backward [num next-nums]
  (if (empty? next-nums)
    num
    (make-prediction-backward (- (first (first next-nums)) num)
                     (rest next-nums))))

(defn pyramid-steps-backward [prev-steps nums]
  (let [new-nums (->> nums
                      (partition 2 1)
                      (map #(apply - (reverse %))))]
    (if (every? #(= 0 %) new-nums)
      prev-steps
      (pyramid-steps (conj prev-steps new-nums) new-nums))))

(->> input
     (parse)
     (map #(pyramid-steps-backward [%] %))
     (map #(make-prediction-backward 0 %))
     (reduce +)
     )