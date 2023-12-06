(require '[clojure.math :as math])
(require '[clojure.string :as str])
(def input (slurp "input/day4.txt"))

(def sample (str "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"))

(defn parse-line [line]
   (->> line
        (re-seq #"(Card +\d+: +)(.*$)")
        ((comp last last))
        (#(str/split % #" +\| +"))
        (map #(str/split % #" +"))
        (map #(map read-string %))
       )
)

(defn parse-input [input]
  (->> input
       (#(str/split % #"\n"))
       (map parse-line)
       )
  )

(defn is-winning [card] 
  (let [win-n (first card) nums (second card)] (map #(contains? win-n %) nums))
  )

(defn point-card-win [card]
 (->> card
      (map set)
      (is-winning)
      (filter true?)
      (count)
      (#(- % 1))
      (math/pow 2)
      (math/floor) 
      )
)

(defn count-card-win [card]
  (->> card
       (map set)
       (is-winning)
       (filter true?)
       (count)))



;P1
( ->> input
 (parse-input)
 (map point-card-win)
 (reduce +)
)


;P2
(defn recurse-cards [cards to-process counter] ; Yay for stack overflow! :(
  (print "ran!")
  (let [curr-card (first to-process)
        curr-card-id (first curr-card)
        curr-wins (second curr-card)
        to-add (subvec cards (+ curr-card-id 1) (+ curr-card-id curr-wins 1))
        next-to-process (concat to-add (rest to-process))]
    (print [curr-card-id curr-wins counter])
    (if (not-empty next-to-process)
      (recurse-cards cards next-to-process (+ counter 1))
      (+ counter 1)))
  )

(defn cards-added [cards card]
  (let [ curr-card-id (first card)
         curr-wins (second card)
         to-add (subvec cards (+ curr-card-id 1) (+ curr-card-id curr-wins 1))
       ]
    (+ curr-wins (reduce + (map #(cards-added cards %) to-add)))
    )
  )


(let [indexed-cards (->> input
                         (parse-input)
                         (map count-card-win)
                         (map-indexed vector))]
  (->> (map #(cards-added (into (vector) indexed-cards) %) indexed-cards)
       (reduce +)
       (+ (count indexed-cards))
       )
  )


