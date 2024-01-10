(require '[clojure.string :as str])
(def input (slurp "input/day23.txt"))

(def sample "#S#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#")

(defn is-path [garden pos]
  (contains? (set [\. \> \v \<]) (garden pos)))

(def dirs {\. [[0 1] [0 -1] [1 0] [-1 0]]
           \S [[0 1] [0 -1] [1 0] [-1 0]]
           \> [[0 1]]
           \< [[0 -1]]
           \v [[1 0]]})

(defn next-valid-steps [garden pos]
  (let [available-dirs (dirs (garden pos))
        possible (map #(map + pos %) available-dirs)]
    ;(println pos available-dirs possible)
    (filter (partial is-path garden) possible)))

(defn parse-to-map [string]
  (let [split-str (str/split string #"\n")]
    (->> (for [[y line] (map-indexed vector split-str)]
           (for [[x char] (map-indexed vector line)]
             [[y x] char]))
         (reduce concat)
         (into {}))))

(defn key-of [dict val]
  (first (for [[k v] dict :when (= v val)] k)))

(defn is-end? [garden pos]
  (not (contains? garden (map + [1 0] pos))))

(defn take-next-step [garden to-walk]
  (if (> (count to-walk) 1000000000000) "FUCK"
      (loop [to-walk to-walk
             max-steps 0]
        (if (empty? to-walk) max-steps
            (let [[curr path] (first to-walk)
                  rst (disj to-walk [curr path])
                  next-steps (->> (next-valid-steps garden curr)
                                  (filter #(not (contains? path %))))]
        ;(println curr path next-steps)
              (if (is-end? garden curr)
                (recur rst
                       (max (count path) max-steps))
                (recur (into rst (map #(vector % (conj path curr)) next-steps))
                       max-steps)))))))

(defn walk-steps [text]
  (let [garden (parse-to-map text)
        start (key-of garden \S)]
    (->> (take-next-step garden (set [[start (set [])]])))))

; (walk-steps input) ;P1  

(defn next-valid-steps [garden pos]
  (let [available-dirs (dirs \.)
        possible (map #(map + pos %) available-dirs)]
    ;(println pos available-dirs possible)
    (filter (partial is-path garden) possible)))

(defn build-graph [garden start]
  (loop [to-do [[start start start 0]] graph {}]
    (if (empty? to-do) graph
        (let [[last-node curr no-return dist] (first to-do)
              next-steps (->> (next-valid-steps garden curr)
                              (filter #(not (= no-return %))))]
          ;(println last-node curr no-return next-steps)
          (cond
            (contains? (graph last-node) [curr dist])
            (recur (rest to-do) graph)

            (or (> (count next-steps) 1) (is-end? garden curr))
            (let [new-nodes (map #(vector curr % curr 1) next-steps)
                  n-to-do (into (rest to-do) new-nodes)
                  n-graph-entries {curr
                                   (conj (or (graph curr) (set []))
                                         [last-node dist])
                                   last-node
                                   (conj (or (graph last-node) (set []))
                                         [curr dist])}]
              (recur n-to-do (merge graph n-graph-entries)))
            
            (= (count next-steps) 1)
            (let [n-to-do (conj (rest to-do) [last-node (first next-steps)
                                              curr (inc dist)])]
              (recur n-to-do graph))
            
            :else
            (recur (rest to-do) graph)

              ;(print "Adding" n-graph-entries)
            )))))

(defn index-of [item coll]
  (let [counter (count (take-while #(not= item %) coll))]
    (if (= counter (count coll)) nil counter)))

(defn walk-graph [graph garden start]
  (loop [to-walk (set [[start [] 0]])
         visited {}
         max-steps 0]
    (if (empty? to-walk) max-steps
        (let [[curr path dist-traveled] (first to-walk)
              rst (disj to-walk [curr path dist-traveled])
              next-steps (->> (graph curr)
                              (filter #(not (index-of (first %) path))))
              to-skip (> (visited (conj path curr) 0) dist-traveled)]
          
          (cond
            to-skip (do ;(println (visited path) dist-traveled path)
                        (recur rst visited max-steps))

            (is-end? garden curr)
            (do
              (if (> dist-traveled max-steps) (println dist-traveled max-steps)) 
              (recur rst 
                     (assoc visited (conj path curr) dist-traveled)
                     (max dist-traveled max-steps))) 

            :else
            (let [splitted-paths (map #(vector (first %)
                                               (conj path curr)
                                               (+ dist-traveled (second %)))
                                      next-steps)]
              (recur (into rst splitted-paths)
                     (assoc visited path dist-traveled) 
                     max-steps)))))))

(defn walk-steps-2 [text]
  (let [garden (parse-to-map text)
        start (key-of garden \S)
        graph (build-graph garden start)]
    (println graph)
    (walk-graph graph garden start)
    ))

;(walk-steps-2 input) ;P2
; Run for some 10min, pray it reaches the max answer soon.