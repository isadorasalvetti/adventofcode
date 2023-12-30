(require '[clojure.string :as str])
(def input (slurp "input/day21.txt"))

(def sample "...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........")

(def dim 131)

(defn key-of [dict val]
  (first (for [[k v] dict :when (= v val)] k)))

(defn parse-to-map [string]
  (let [split-str (str/split string #"\n")]
    (->> (for [[y line] (map-indexed vector split-str)]
           (for [[x char] (map-indexed vector line)]
             [[y x] char]))
         (reduce concat)
         (into {}))))

(def dirs [[0 1] [0 -1] [1 0] [-1 0]])

(defn is-path [garden pos]
  (or (= (garden (map #(mod % dim) pos)) \.)
      (= (garden (map #(mod % dim) pos)) \S)))

(defn next-valid-steps [garden pos]
  (let [possible (map #(map + pos %) dirs)]
    (filter (partial is-path garden) possible)))

(defn take-next-step [garden to-walk max-steps]
  (loop [to-walk to-walk visited (set nil) end (set nil)]
    (if (empty? to-walk) end
        (let [[counter pos] (apply min-key first to-walk)
              rst (disj to-walk [counter pos])
              updated-visited (into visited [pos])
              updated-end (if (even? (- max-steps counter)) (into end [pos]) end)]
          (if
           (or (>= counter max-steps)
               (contains? visited pos)) (recur rst updated-visited updated-end)
           (let [next-pos (next-valid-steps garden pos)
                 next-from-pos (map #(vector (+ counter 1) %) next-pos)
                 next-to-walk (into rst next-from-pos)]
             ;(println pos counter updated-end next-from-pos)
             (recur next-to-walk updated-visited updated-end)))))))

(defn walk-steps [text steps-num]
  (let [garden (parse-to-map text)
        start (key-of garden \S)]
    (->> (take-next-step garden (set [[0 start]]) steps-num )
         (count)
         )))


(comment
  (walk-steps input 64) ;P1

  (->> (map #(+ (* % 131) 65) (range 3))
       (map #(walk-steps input %)))
  ;(3867 34253 94909) <- Fit on those
  ;(65 196 327)

  ; P2
  ; 26501365 - Steps
  ; 616990239256449 - Result

  )