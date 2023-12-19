(require '[clojure.string :as str])
(def input (slurp "input/day18.txt"))

(def sample "R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)")

(def l-to-dir {"R" [1 0] "L" [-1 0] "D" [0 1] "U" [0 -1]})

(defn parse-to-map [string]
  (->> string
       (str/split-lines)
       (map #(str/split % #"\s+"))
       (map (fn [[d n c]] [(l-to-dir d) (read-string n) c]))))

(defn mult-dir [[x y] amm] [(* amm x) (* amm y)])

(defn move-in-dir [pos dir amm]
  (apply vector (map + pos (mult-dir dir amm))))

(move-in-dir [0 1] [1 0] 10)

(defn make-dig-side [pos pool-corners corner]
  (if (empty? corner) pool-corners
      (let [[dir amm _] (first corner)
            next-pos (move-in-dir pos dir amm)]
        (make-dig-side next-pos (conj pool-corners next-pos) (rest corner)))))

(defn area [acc corners]
  (if (< (count corners) 2) (long acc)
      (let [[x1 y1] (first corners)
            [x2 y2] (second corners)]
        (area (+' acc (/ (- (*' x1 y2) (*' y1 x2)) 2))
              (rest corners)))))

(defn perimeter [acc corners] 
  (if (< (count corners) 2) (int acc)
       (let [[x1 y1] (first corners)
             [x2 y2] (second corners)]
         (perimeter (+' (abs (- x2 x1)) (abs (- y2 y1)) acc) 
                    (rest corners)))))

;P2
(def hex-value {:0 0 :1 1 :2 2 :3 3 :4 4 :5 5 :6 6 :7 7 :8 8 :9 9 :a 10 :b 11 :c 12 :d 13 :e 14 :f 15})

(defn hex-to-reverse-keyword-vector [hex]
  (->> (str/lower-case hex)
       reverse
       (map str)
       (map keyword)
       vec))

(defn hexify [hex]
  (->> (for [x (map-indexed vector (hex-to-reverse-keyword-vector hex))]
         (->> (* ((second x) hex-value)
                 (Math/pow 16 (first x)))))
       (reduce +)
       (long)))

(def c-to-dir {\0 [1 0] \2 [-1 0] \1 [0 1] \3 [0 -1]})
(defn parse-to-map [string]
  (->> string
       (str/split-lines)
       (map #(str/split % #"\s+"))
       (map #(butlast (drop 2 (last %))))
       (map #(vector (c-to-dir (last %)) (hexify (apply str (butlast %))) 0))
       ))

(parse-to-map sample)

(let [corners (->> input
                   (parse-to-map)
                   ((partial make-dig-side [0 0] []))
                   (#(conj % (first %))))
      cc (- (count corners) 1)
      p (perimeter 0 corners)
      a (area 0 corners)
      ]
  (+' a (/ (- p cc) 2) (/ cc 2) 1)
  )


