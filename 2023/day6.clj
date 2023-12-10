(require '[clojure.string :as str])

(def input "Time:        59     79     65     75
Distance:   597   1234   1032   1328")

(def sample "Time:      7  15   30
Distance:  9  40  200")

(defn parse-input [s] (->> s
                           (#(str/split % #":"))
                           (rest)
                           (map #(re-seq #"\d+" %))
                           (map #(map read-string %))
                           (apply zipmap)))

(defn distance-travelled [time button-pressed] 
  (* (- time button-pressed) button-pressed))

(defn ways-to-win [time dist]
  (->> (map #(distance-travelled time %) (range 1 time))
       (filter #(< dist %))
       (count))
  )

(->> (parse-input input)
     (map #(apply ways-to-win %))
     (reduce *))

; P2
(defn parse-input [s] (->> s
                           (#(str/split % #":"))
                           (rest)
                           (map #(re-seq #"\d+" %))
                           (map #(apply str %))
                           (map read-string))
                           )

(->> (parse-input input)
     (apply ways-to-win))
