(require '[clojure.string :as str])
(def input (slurp "input/day7.txt"))
(def tiger-input (slurp "input/tiger-day7.txt"))

(def sample "32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483")

(defn score [hand](->> hand
                       (frequencies)
                       (map second)
                       (sort >)
                       (#(concat % (repeat 0)))
                       (take 2)
                       (apply vector)))

(defn card-to-num [card]
  (let [face-cards {\T 10 \J 11 \Q 12 \K 13 \A 14}]
    (if (contains? face-cards card)
      (face-cards card)
      (read-string (str card)))))

(card-to-num \8)

(defn comp-cards [cards1 cards2]
  (if (= (first cards1) (first cards2))
    (comp-cards (rest cards1) (rest cards2))
    (< (card-to-num (first cards1))
       (card-to-num (first cards2)))))

(defn sorter [hand1 hand2]
  (let [score1 (first hand1)
        score2 (first hand2)] 
  (if (= score1 score2)
    (comp-cards (second hand1) (second hand2))
    (< (compare score1 score2) 0)
    )))


(score "32T3K")

;P1
(->> input
     (#(str/split % #"\n"))
     (map #(str/split % #" "))
     (map #(vector (score (first %)) (first %) (second %)))
     (sort sorter)
     ;(map second)
     (map #(read-string (last %)))
     (map-indexed #(* (+ %1 1) %2))
     (reduce +)
     )


;P2
(defn card-to-num2 [card]
  (let [face-cards {\T 10 \J 0 \Q 12 \K 13 \A 14}]
    (if (contains? face-cards card)
      (face-cards card)
      (read-string (str card)))))

(defn comp-cards2 [cards1 cards2]
  (if (= (first cards1) (first cards2))
    (comp-cards2 (rest cards1) (rest cards2))
    (< (card-to-num2 (first cards1))
       (card-to-num2 (first cards2)))))

(defn sorter2 [hand1 hand2]
  (let [score1 (first hand1)
        score2 (first hand2)]
    (if (= score1 score2)
      (comp-cards2 (second hand1) (second hand2))
      (< (compare score1 score2) 0))))

(defn add-js [freqs]
  (if (and (contains? freqs \J) (> (count freqs) 1))
    (let [most-common (key (apply max-key val (dissoc freqs \J)))]
      (-> (update freqs most-common
                        #(+ % (freqs \J)))
        (dissoc \J))) 
    freqs))

(defn score2 [hand] (->> hand
                         (frequencies)
                         (add-js) 
                         (map second)
                         (sort >)
                         (#(concat % (repeat 0))) 
                         (take 2)
                         (apply vector)))

(score "KTJJT")

(->> input
     (#(str/split % #"\n"))
     (map #(str/split % #" "))
     (map #(vector (score2 (first %)) (first %) (second %)))
     (sort sorter2)
     ;(map second)
     (map #(read-string (last %)))
     (map-indexed #(* (+ %1 1) %2))
     (reduce +)
     )

