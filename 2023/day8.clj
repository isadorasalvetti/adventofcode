(require '[clojure.string :as str])
(def input (->> (slurp "input/day8.txt")
                (#(str/split % #"\n+"))))

(def sample (->> "LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)"
                 (#(str/split % #"\n+"))))

(defn parse-maps [sample](->> (rest sample)
                              (map #(str/split % #"\s+=\s+"))
                              (map #(vector (first %)
                                            (re-seq #"[A-Z0-9]+" (second %))))
                              (into {})))

(vector (take 5 (cycle (first sample))) (parse-maps sample))


(defn walk-map [steps paths curr-location]
  (loop [steps steps curr-location curr-location count 0]
    ;(print curr-location count "\n")
    (if (= curr-location "ZZZ")
      count
      (let [next-locations (paths curr-location)]
        (if (= (first steps) \L)
          (recur (rest steps) (first next-locations) (+ count 1))
          (recur (rest steps) (second next-locations) (+ count 1)))))))

(walk-map (cycle (first input)) (parse-maps input) "AAA")


;P2
(def sample2 (->> "LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)"
                  (#(str/split % #"\n+"))))

(defn walk-map2 [steps paths curr-location]
  (loop [steps steps curr-location curr-location count 0]
    (if (= (last curr-location) \Z)
      count
      (let [next-locations (paths curr-location)]
        (if (= (first steps) \L)
          (recur (rest steps) (first next-locations) (+ count 1))
          (recur (rest steps) (second next-locations) (+ count 1)))))))


( ->> ["KJA" "BGA" "AAA" "LTA" "NJA" "TTA"] 
 (map 
  #(walk-map2 (cycle (first input)) (parse-maps input) %)
  )
)