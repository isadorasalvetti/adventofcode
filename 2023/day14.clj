(require '[clojure.string :as str])

(def input (slurp "input/day14.txt"))
(def sample "O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....")

(def max-steps 4000000000)
(def size 100)

(defn parse-to-map [string]
  (let [split-str (str/split string #"\n")]
    (->> (for [[y line] (map-indexed vector split-str)]
           (for [[x char] (map-indexed vector line)]
             [[y x] char]))
         (reduce concat)
         (into {}))))

(defn in-bound [i]
  (and (< i size) (>= i 0)))

(defn get-next-pos-ns [[y x] dir]
  (if (in-bound x) [y (+ x 1)]
      (if (in-bound y) [(- y dir) 0] nil)))

(defn get-next-pos-ew [[y x] dir]
  (if (in-bound y) [(+ y 1) x]
      (if (in-bound x) [0 (- x dir)] nil)))

(defn get-next-pos [[y x] [dy dx]]
  (if (not= dy 0)
    (get-next-pos-ns [y x] dy)
    (get-next-pos-ew [y x] dx)))

(defn get-next-rock [rock-map [y x] dir]
  (if (= (rock-map [y x]) \O)
    [y x]
    (let [next-pos (get-next-pos [y x] dir)]
      (if (nil? next-pos) nil
          (get-next-rock rock-map next-pos dir)))))

(defn slide-rock [rock-map [y x] [diry dirx]]
  (let [slide-into [(+ y diry) (+ x dirx)]]
    (if (or (not (contains? rock-map slide-into))
            (= (rock-map slide-into) \#)
            (= (rock-map slide-into) \O))
      (assoc rock-map [y x] \O)
      (slide-rock rock-map slide-into [diry dirx]))))

(defn move-all-rocks [rock-map [y x] dir]
  (if (nil? x) rock-map
      (let [to-move (get-next-rock rock-map [y x] dir)]
        (if (nil? to-move)
          rock-map
          (-> rock-map
              (assoc to-move \.)
              (slide-rock to-move dir)
              (move-all-rocks (get-next-pos [y x] dir) dir))))))

(defn stress [keys]
  (->> keys
       (map #(- size (first %)))
       (reduce +)))

(comment ;P1 
  (->> sample
       (parse-to-map)
       (#(move-all-rocks % [0 0] [-1 0]))
       (filter #(= (second %) \O))
       (keys)
       (stress)))

(def start-dirs '([[0 0] [-1 0]]
                  [[0 0] [0 -1]]
                  [[99 0] [1 0]]
                  [[0 99] [0 1]]))

(defn index-of [item coll]
  (let [counter (count (take-while #(not= item %) coll))]
    (if (= counter (count coll)) nil counter)))

(defn cicle-dirs [map dirs rocks-hit]
  ; slower cause dumb and didnt concider a full CYCLE of 4 moves ):
  (if (> (count rocks-hit) 1000) "fuck"
      (let [[s d] (first dirs)
            moved (move-all-rocks map s d)
            rocks (->> moved
                       (filter #(= (second %) \O))
                       (keys)
                       (set))
            f-match (index-of rocks rocks-hit)] 
        (if (some? f-match) [f-match (drop f-match rocks-hit)]
            (cicle-dirs moved (rest dirs) (conj rocks-hit rocks))))))

(comment ;P2
  (let [parsed-input (parse-to-map input)
        [offset cycle] (cicle-dirs parsed-input (cycle start-dirs) [])
        tg (mod (- max-steps offset) (count cycle))
        ]
    (stress (nth cycle (- tg 1)))
    )
)