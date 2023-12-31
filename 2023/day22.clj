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
    (if (empty? supporting) true
        (every? #(contains? se-set %) supporting))))

(defn is-single-support [[block supporting] block-supporting]
  (let [supported-elsewhere (vals (dissoc block-supporting block))
        se-set (reduce into (set []) supported-elsewhere)]
    [block (filter #(not (contains? se-set %)) supporting)]    
    ))

(defn assoc-reverse [values key dict]
  (if (empty? values) dict
      (assoc-reverse (rest values) key
                     (assoc dict (first values) 
                            (conj (dict (first values)) key)))))

(defn supported-by [supporting]
  (loop [to-do supporting 
         acc {}]
    (if (empty? to-do) acc
        (let [[block sup] (first to-do)]
          (recur (rest to-do)
                 (assoc-reverse sup block acc))))))

(defn will-fall? [candidate fallen supported]
  ;(println candidate (supported candidate) fallen)
  (every? #(contains? fallen %) (supported candidate)))

(defn would-fall [to-fall supporting supported fallen acc]
  (if (empty? to-fall) acc
      (let [falling (first to-fall)
            u-fallen (conj fallen falling)
            candidates-to-fall (supporting falling)
            actual-fallen (filter #(will-fall? % u-fallen supported)
                                  candidates-to-fall)]
        ;(println acc falling candidates-to-fall actual-fallen) 
        (would-fall (into (rest to-fall) actual-fallen)
                    supporting supported u-fallen (inc acc)))))

(defn p1 [supporting] (->> (filter #(to-remove? % supporting) supporting)
                           (count)))

(defn p2 [fallen supporting supported] 
  (->> (map 
        #(would-fall [%] supporting supported (set {}) -1) 
        fallen)
       (reduce +)))

(let [parsed-input (sort-by #(first (last %)) (parse-input input))
      fallen (fall-blocks parsed-input)
      supporting (loop [to-process fallen acc {}]
                   (if (empty? to-process) acc
                       (recur (rest to-process)
                              (into acc [(supporting-next to-process)]))))
      supported (supported-by supporting)] 
  
  (println "P1" (p1 supporting))
  (println "P2" (p2 fallen supporting supported))
  )