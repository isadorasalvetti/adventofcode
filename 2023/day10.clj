(require '[clojure.string :as str])
(require '[clojure.set :as st])

(def input (slurp "input/day10.txt"))

(def sample "...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........")

(defn apply-first-val [val col]
  (map #(concat [val] %) col))

(defn parse [string] 
  (->> string
       (#(str/split % #"\n"))
       (map #(map-indexed vector %))
       (map-indexed vector)
       (map #(apply apply-first-val %))
       (reduce concat)
       (map #(vector (take 2 %) (last %)))
       (into {}))
     )

(def north [-1 0])
(def south [1 0])
(def east [0 1])
(def west [0 -1])

(defn go-next [dict last-location location]
  (let [pipe-shape (dict location)
        next-steps (case pipe-shape
                     \- [(map + location east)
                         (map + location west)]
                     
                     \| [(map + location north)
                         (map + location south)]
                     
                     \L [(map + location north)
                         (map + location east)]
                     
                     \J [(map + location north)
                         (map + location west)]
                     
                     \7 [(map + location south)
                         (map + location west)]
                     
                     \F [(map + location south)
                         (map + location east)] 
                     "fuck")] 
    (->> next-steps
         (filter #(not= last-location %))
         first)) 
  )

(defn find-cycle [dict last-spot spot] 
  (loop [last-spot last-spot spot spot counter 0]
    (if (or (= (dict spot) \S) (< 50000 counter)) counter
      (let [next-step (go-next dict last-spot spot)]
        (recur spot next-step (inc counter))
        ))))

(defn start [dict]
  (first (keep #(when (= (val %) \S)
           (key %)) dict)))

(defn first-step [dict s]
  (if (some #(= % (dict (map + s east))) [\- \7 \J] ) (map + s east)
      (if (some #(= % (dict (map + s south))) [\| \L \J]) (map + s south)
          (if (some #(= % (dict (map + s west)))  [\- \L \F]) (map + s west)
              (map + s north)))))

(let [dict (parse input)
      s (start dict)
      f-step (first-step dict s)] 
  (-> (find-cycle dict s f-step)
      (quot 2)
      (+ 1)))

; P2 
(defn find-cycle [dict last-spot spot]
  (loop [last-spot last-spot spot spot acc []]
    (if (or (= (dict spot) \S) (< 50000 (count acc))) (conj acc spot)
        (let [next-step (go-next dict last-spot spot)]
          (recur spot next-step (conj acc spot))))))

(defn is-right-side [el]
  (some #(= % el) [\F \7 \| \S]))

(defn makegap [row n1 n2]
  (map vector (repeat (- n2 n1) row) (range n1 n2)))

(defn line-gaps [line]
  (->> (partition 2 line)
       (map #(makegap (first (first %)) 
                      (second (first %))
                      (second (second %))))))

(let [dict (parse input)
      s (start dict)
      f-step (first-step dict s)
      border (find-cycle dict s f-step)]
  (print s f-step "\n")
  (->>  border
        (filter #(is-right-side (dict %)))
        (group-by first)
        (vals)
        (map #(sort-by second %))
        (map line-gaps)
        (flatten)
        (partition 2)
        (set)
        (#(st/difference % (set border)))
        (count)
        )
  )




