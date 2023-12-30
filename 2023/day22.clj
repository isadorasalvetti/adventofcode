(require '[clojure.string :as str])
(def input (slurp "input/day22.txt"))

(def sample "1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9")

(defn parse-line [line]
  (->> (str/split line #"~")
       (map #(str/split % #","))
       (map #(map read-string %))
       (apply (partial map vector))))

(defn parse-input [txt]
  (->> (str/split txt #"\n")
       (map parse-line)))

(defn overlaps [[min1 max1] [min2 max2]] 
  (<= (max min1 min2) (min max1 max2))) 

(defn overlaps-cube [[x1 y1 _] [x2 y2 _]]
  (and (overlaps x1 x2) (overlaps y1 y2)))

(defn max-z [max-range] (second (last max-range)))
(defn min-z [max-range] (first (last max-range)))

(defn supporting-next [blocks]
  (let [block (first blocks)
        overlapping-above (filter (partial overlaps-cube block) (rest blocks)) 
        supporting-height (+ (max-z block) 1)]
    [block (filter #(= supporting-height (min-z %)) overlapping-above)]
     ))

(defn fall-blocks [blocks]
  (loop [to-process blocks 
         fallen []]
    (if (empty? to-process) fallen
        (let [block (first to-process)
              [xrange yrange [minz maxz]] block
              overlapping-bellow (filter (partial overlaps-cube block) fallen)
              supporting-height (apply max (conj (map max-z overlapping-bellow) -1))
              nz (+ 1 supporting-height)]
          (recur (rest to-process) 
                 (conj fallen [xrange yrange [nz (+ nz (- maxz minz))]]))))))

(defn to-remove? [[block supporting] block-supporting]
  (let [supported-elsewhere (vals (dissoc block-supporting block))
        se-set (reduce into (set []) supported-elsewhere)]
    (if (empty? supporting) (do 
                              ;(println block "Supports nothing") 
                              true)
        (if (every? #(contains? se-set %) supporting)
          (do 
            ;(println block "Supported elsewere") 
            true)
          (do 
            ;(println block "Only support") 
            false)))))

(let [parsed-input (sort-by #(first (last %)) (parse-input input))
      fallen (fall-blocks parsed-input)
      supporting (loop [to-process fallen acc {}]
                   (if (empty? to-process) acc
                       (recur (rest to-process)
                              (into acc [(supporting-next to-process)]))))]
  ;(println supporting)
  (->> (filter #(to-remove? % supporting) supporting)
       (count)
       ))

(comment
  (overlaps-cube [[1 1] [1 1] [8 9]] [[1 1] [1 1] [8 9]])
  (overlaps-cube [[2 2] [1 1] [8 9]] [[1 1] [1 1] [8 9]]))