(require '[clojure.string :as str]
         '[clojure.set :as st])

(def input (slurp "input/day16.txt"))
(def sample ".|...\\....
|.-.\\.....
.....|-...
........|.
..........
.........\\
..../.\\\\..
.-.-/..|..
.|....-|.\\
..//.|....")

(defn parse-to-map [string]
  (let [split-str (str/split string #"\n")]
    (->> (for [[y line] (map-indexed vector split-str)]
           (for [[x char] (map-indexed vector line)]
             [[y x] char]))
         (reduce concat)
         (into {}))))

(def symbol-map {[0 1] \> [0 -1] \< [1 0] \v [-1 0] \^ })

(defn print-map-step [visited size [hdir hpos]]
  (let [keys (->> visited
                  (#(map second %))
                  (set))]
    (->> (for [y (range size)
               x (range size)]
           (cond
             (= hpos [y x]) (symbol-map hdir)
             (contains? keys [y x]) \#
             :else \.))
         (partition size)
         (map (partial apply str))
         (str/join "\n")
         (str (newline)))))

(defn print-map [visited size]
  (let [keys (->> visited
                  (map reverse)
                  (map #(map (partial apply vector) %))
                  (map #(apply vector %))
                  (into {}))]
    (->> (for [y (range size)
               x (range size)]
           (cond
             (contains? keys [y x]) (symbol-map (keys [y x]))
             :else \.))
         (partition size)
         (map (partial apply str))
         (str/join "\n")
         (str (newline) (newline)))))

(defn rot-right [[dy dx]] [dx (- dy)])
(defn rot-left [[dy dx]] [(- dx) dy])

(defn apply-dir-pos [[dy dx] pos]
  [[dy dx] (map + pos [dy dx])])

(defn next-move [space [dy dx] pos]
  (let [curr-inst (space pos)
        moves (case curr-inst
                \. [(apply-dir-pos [dy dx] pos)]
                \\ (if (= dy 0)
                     [(apply-dir-pos (rot-right [dy dx]) pos)]
                     [(apply-dir-pos (rot-left [dy dx]) pos)])
                \/ (if (= dx 0)
                     [(apply-dir-pos (rot-right [dy dx]) pos)]
                     [(apply-dir-pos (rot-left [dy dx]) pos)])
                \| (if (= dx 0)
                     [(apply-dir-pos [dy dx] pos)]
                     [(apply-dir-pos [1 0] pos)
                      (apply-dir-pos [-1 0] pos)])
                \- (if (= dy 0)
                     [(apply-dir-pos [dy dx] pos)]
                     [(apply-dir-pos [0 1] pos)
                      (apply-dir-pos [0 -1] pos)]))]
    moves))

(defn run-lazer [space dir pos]
  (loop [to-go [[dir pos]]
         path (set [])]
    ;(println (print-map-step path 10 (first to-go)) (first to-go))
    (if (> (count path) 5000000) (print "FUCK")
        (if (empty? to-go) path
            (let [[ndir npos] (first to-go)]
              (if (and (contains? space npos)
                       (not (contains? path [ndir npos])))
                (recur (concat (rest to-go) (next-move space ndir npos))
                       (conj path [ndir npos]))
                (recur (rest to-go)
                       path)))))))

(-> (parse-to-map input)
    (run-lazer [0 1] '(0 0))
    ;(print-map 110)
    ;(print)
    (#(map second %))
    (set)
    (count)
    )

;P2
(defn candidates [size]
  (concat
   (->> (range 1 (- size 1))
        (map #(vector [0 1] [% 0])))
   (->> (range 1 (- size 1))
        (map #(vector [0 -1] [% (- size 1)])))
   (->> (range 1 (- size 1))
        (map #(vector [1 0] [0 %])))
   (->> (range 1 (- size 1))
        (map #(vector [-1 0] [(- size 1) %])))))

(defn count-energy [space [dir pos]]
  (-> space
      (run-lazer dir pos)
      (#(map second %))
      (set)
      (count)))

(count-energy (parse-to-map sample) [[0 1] '(0 0)])

(let [parsed (parse-to-map input)]
  (->> (map (partial count-energy parsed) (candidates 110))
  (reduce max)))

(comment

  (def asample ".-|..
..\\..") ;([0 0] [0 1] [0 2] [1 2] [1 3] [1 4])

  (def bsample "..\\..
../..") ;([0 0] [0 1] [0 2] [1 0] [1 1] [1 2])

  (def csample ".\\/..
.\\/..") ;([0 0] [0 1] [0 2] [0 3] [0 4] [1 1] [1 2])

  (def dsample ".|...
.\\|..
..-..") ;([0 0] [0 1] [0 2] [1 2] [2 0] [2 1] [2 2] [2 3] [2 4])"

  (-> asample
      (parse-to-map)
      (run-lazer [0 1] '(0 0))
    ;(print-map 5)
    ;(print)
      )

  (map #(-> %
            (parse-to-map)
            (run-lazer [0 1] '(0 0))
            (print-map 5)
            (print)) [asample bsample csample dsample])
  )
