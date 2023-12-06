(require '[clojure.string :as str])
(def input (slurp "input/day5.txt"))

(def sample "seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4")


(defn parse-input [s] (->> s
                           (#(str/split % #":"))
                           (rest)
                           (map #(re-seq #"\d+" %))
                           (map #(map read-string %))))


(defn num-to-map [nums]
  (let [[dest source lenght] nums]
    (zipmap (range source (+ source lenght))
            (range dest (+ dest lenght)))))

(defn source-sp-dest [seed nums]
  (let [[dest source lenght] nums]
    (if (and (>= seed source) (<= seed (+ source lenght)))
      (+ dest (- seed source))
      nil)))


(defn source-to-dest [source dest-map]
  (let [sp-dest (->> (map #(source-sp-dest source %) dest-map)
                     (filter some?)
                     (first))]
    (if sp-dest sp-dest source)))

(comment
  (let [test-dict '((50 98 2) (52 50 48))]
    (source-to-dest 79 test-dict)
    ;(source-sp-dest 14 (second test-dict))
    ))

(defn do-conversion [sources, dest-maps]
  (if (empty? dest-maps) sources
      (do-conversion (map #(source-to-dest % (first dest-maps)) sources) (rest dest-maps))))

(let [parsed-input (parse-input input)
      seeds (first parsed-input)
      maps (rest parsed-input)]
  (->> maps
       (map #(partition 3 %))
       (do-conversion seeds)
       (apply min)))

;P2

(defn amm-to-range [nums]
  [(first nums) (+ (first nums) (second nums) (- 1))])

(defn get-next-range [bound target-bounds replace-bounds new-bounds]
  (if (empty? target-bounds) (conj new-bounds bound) 
      (if (> (count new-bounds) 10) new-bounds
          (let [to-compare (first target-bounds)
                to-replace (replace-bounds to-compare)
                range-lenght (- (second bound) (first bound))
                dist-to-start (- (first bound) (first to-compare))
                dist-to-end (- (second to-compare) (second bound))]

            (print bound target-bounds to-replace range-lenght dist-to-start dist-to-end new-bounds "\n")

            (if (>= dist-to-start 0)
              (if (>= dist-to-end 0)
                (conj new-bounds [(+ (first to-replace) dist-to-start)
                                  (- (second to-replace) dist-to-end)])
                (if (< (+ range-lenght dist-to-end) 0)
                  (get-next-range bound (rest target-bounds) replace-bounds new-bounds)
                  
                  (let [cut-point (+ (last bound) dist-to-end)]
                    (get-next-range [(first bound) cut-point] target-bounds replace-bounds
                                    (conj new-bounds
                                          [(+ cut-point 1) (last bound)])))))
              
              (if (< (+ range-lenght dist-to-start) 0)
                (get-next-range bound (rest target-bounds) replace-bounds new-bounds)
                (let [cut-point (+ (first bound) dist-to-start)]
                  (get-next-range [(+ cut-point 1) (last bound)] target-bounds replace-bounds
                                  (conj new-bounds
                                        [(first bound) cut-point])))))))))



(get-next-range [79 93] (sort '([98 100] [50 98])) {[98 100] [50 52], [50 98] [52 100]} [])

(defn range-to-ranges [seed-bounds dest-map]
  (let [dest-bounds (map #(amm-to-range [(first %) (last %)]) dest-map)
        source-bounds (map #(amm-to-range [(second %) (last %)]) dest-map)
        target-bounds (zipmap source-bounds dest-bounds)]
    (print seed-bounds (sort source-bounds) target-bounds "-Taerget bounds \n")
    (get-next-range seed-bounds (sort source-bounds) target-bounds [])))

(defn process-ranges [seeds maps]
  (let [flat-seeds (reduce concat seeds)
        new-ranges (map #(range-to-ranges % (first maps)) flat-seeds)]
    (print flat-seeds  "\n")
    (if (empty? maps) seeds (process-ranges new-ranges (rest maps)))))

(let [parsed-input (parse-input sample)
      seeds (partition 2 (first parsed-input))
      seed-bounds (map amm-to-range seeds)
      maps (->> (rest parsed-input) (map #(partition 3 %)))] 
  (process-ranges [seed-bounds] maps))

