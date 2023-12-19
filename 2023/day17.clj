(require '[clojure.string :as str])

(def input (slurp "input/day17.txt"))
(def sample "2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533")

(defn parse-to-map [string]
  (let [split-str (str/split string #"\n")]
    (->> (for [[y line] (map-indexed vector split-str)]
           (for [[x char] (map-indexed vector line)]
             [[y x] (read-string (str char))]))
         (reduce concat)
         (into {}))))

(defn rot-right [[dy dx]] [dx (- dy)])
(defn rot-left [[dy dx]] [(- dx) dy])
(defn cost-path [space path] (reduce + (map space path)))

(cost-path (parse-to-map sample) [[0 0] '(1 0) '(2 0) '(2 1) '(3 1) '(3 0) '(4 0) '(5 0)])
(cost-path (parse-to-map sample) [[0 0] '(0 1) '(1 1) '(2 1) '(3 1) '(3 0) '(4 0) '(5 0)])

(defn is-valid-move [[dir pos acc _] space visited]
  (and (< acc 3)
       (contains? space pos)
       (not (contains? visited [dir pos acc]))))

(defn step-options [[last-dir pos consecutive cost] space visited]
  (let [pr (apply vector (map + pos (rot-right last-dir)))
        pl (apply vector (map + pos (rot-left last-dir)))
        pf (apply vector (map + pos last-dir))
        options [[(rot-left last-dir) pl 0 (+ cost (or (space pl) -1))]
                 [(rot-right last-dir) pr 0 (+ cost (or (space pr) -1))]
                 [last-dir pf (inc consecutive) (+ cost (or (space pf) -1))]]]
    (filter #(is-valid-move % space visited) options)))

(defn comp [[ad ap aacc acost] [cd cp cacc ccost]]
  ;(println [ad ap aacc acost] [cd cp cacc ccost])
  (compare [acost ad ap aacc] [ccost cd cp cacc]))

(-> (sorted-set-by comp [0 0 1 2] [0 0 1 0]) 
    (conj [0 0 1 3])
    (into [[0 0 1 2] [0 0 1 12] [0 0 2 2]])
    )

(defn travel [space st end]
  (loop [to-go (sorted-set-by comp [[0 1] st -1 0] [[1 0] st -1 0]) visited (set [])]
    (let [curr (first to-go)
          [cd cp cacc ccost] curr
          updated-visited (conj visited [cd cp cacc])]
      ;(print curr)
      (if (contains? visited [cd cp cacc])
        (recur (disj to-go curr) visited)
        (if (> ccost 100000) (let [] (print "Fuck" (count to-go) " ") ccost)
            (if (= cp end) ccost
                (recur (into (disj to-go curr)
                             (step-options curr space updated-visited))
                       updated-visited)))))))

;P1
(-> (parse-to-map input)
    (travel [0 0] [140 140])
    (time))

;P2
(defn step-options [[last-dir pos consecutive cost] space visited]
  (let [pr (apply vector (map + pos (rot-right last-dir)))
        pl (apply vector (map + pos (rot-left last-dir)))
        pf (apply vector (map + pos last-dir))
        rots [[(rot-left last-dir) pl 0 (+ cost (or (space pl) -1))]
              [(rot-right last-dir) pr 0 (+ cost (or (space pr) -1))]]
        forward [last-dir pf (inc consecutive) (+ cost (or (space pf) -1))]
        options (if (>= consecutive 3)
                  (into rots [forward])
                  [forward])] 
    (filter #(is-valid-move % space visited) options)))

(defn is-valid-move [[dir pos acc _] space visited]
  (and (< acc 10)
       (contains? space pos)
       (not (contains? visited [dir pos acc]))))

(-> (parse-to-map input)
    (travel [0 0] [140 140])
    (time))
