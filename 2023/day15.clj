(require '[clojure.string :as str])

(def input (slurp "input/day15.txt"))
(def sample "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7")

(defn parse [data]
  (str/split data #","))

(defn hash [string acc]
  (if (empty? string) acc
      (hash (rest string)
            (->> acc
                 (+ (int (first string)))
                 (* 17)
                 (#(rem % 256))
                 ))
      ))

;(hash "HASH" 0) ; 52

(comment ;P1 
  (->> input
     (parse)
     (map #(hash % 0))
     (reduce +)
     ))

(defn separator [coll]
  (let [counter (count (take-while #(and (not= \- %) (not= \= %)) coll))]
    (if (= counter (count coll)) nil counter)))

(defn remove-f [el vec]
  (remove #(= el (first %)) vec))

(defn parse-instruction [string boxes]
  (let [[id op] (split-at (separator string) string)
        hashed-id (hash id 0)
        in-box (or (boxes hashed-id) {})]
    (if (= (first op) \=)
      (assoc boxes hashed-id
             (assoc in-box id (read-string (str (second op)))))
      (assoc boxes hashed-id
             (dissoc in-box id)))))

(defn parse-instructions [instructions boxes]
  (if (empty? instructions) boxes
      (let [updated-boxes (parse-instruction (first instructions) boxes)]
        (parse-instructions (rest instructions) updated-boxes))))

(defn power [id box]
  (if (empty? box) 0
      (->> box
           (map second)
           (map #(* (inc id) %))
           (map-indexed #(* (inc %1) %2))
           )))

(->> input
     (parse)
     (#(parse-instructions % {}))
     (map #(apply power %))
     (flatten)
     (reduce +)
     )


