(require '[clojure.string :as str])

(def input (slurp "input/day13.txt"))
(def sample "#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#")

(defn parse [field] 
  (->> field
    (#(str/split % #"\n\n"))
    (map #(str/split % #"\n")))
)

(defn is-symetric? [line sep]
  (let [fst (take sep line)
        scnd (drop sep line)
        to-comp-ammount (min (count fst) (count scnd))]
    ;(print line sep (= (take to-comp-ammount scnd) (reverse (take-last to-comp-ammount fst))) "\n")
    (= (take to-comp-ammount scnd) (reverse (take-last to-comp-ammount fst)))
    ))

(defn find-symetry [block]
  (->> (for [i (range 1 (count (first block)))
             :when (every? identity (map #(is-symetric? % i) block))]
         i)
       (first))
)

(defn transpose [block]
   (apply mapv vector block))

(defn process-block [block]
  ;(print block "\n")
  (or (find-symetry block) (* 100 (find-symetry (transpose block)))))

(->> input
     (parse)
     (map process-block)
     (reduce +)
     )

; P2
(defn how-many-off [line sep] 
  (let [fst (take sep line)
        scnd (drop sep line)
        to-comp-ammount (min (count fst) (count scnd))]
    (->> (map = (take to-comp-ammount scnd) (reverse (take-last to-comp-ammount fst)))
         (map {false 1 true 0})
         (reduce +)
         )))

(defn find-one-off [block]
  (->> (for [i (range 1 (count (first block)))
             :when (= 1 (reduce + (map #(how-many-off % i) block)))]
         i)
       (first)))

(defn process-block-2 [block]
  ;(print block "\n")
  (or (find-one-off block) (* 100 (find-one-off (transpose block)))))

(->> input
     (parse)
     (map process-block-2)
     (reduce +))


(comment
  (first (parse sample))
  (find-one-off (first (parse sample)))

  (how-many-off "#.##..##." 5) ; false
  (how-many-off ".####...##." 3) ; false
  (how-many-off ".#.." 2)

  (transpose (second (parse sample)))
  (find-symetry (first (parse sample))) ; 5

  (is-symetric? [\# \# \. \# \# \. \#]  4)

  (is-symetric? "#.##..##." 5) ; true
  (is-symetric? ".####...##." 3) ; true
  (is-symetric? ".####...##." 4) ; false

  (".#.#....#.#.###
...#....#...#..
....####....#..
.#.######.#..##
###..##..###...
.#...##...#.###
..##....##.....
.#...##...#..##
..##....##..#..
.##########....
##.#.##.#.##...
....####.....##
##...##...##.##
##.#.##.#.##.##
...#....#...###
##...##...#####
....#..........")

  )