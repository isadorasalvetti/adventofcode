(require '[clojure.string :as str])
(def input (slurp "input/day25-no-edge.txt"))

(def sample "jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr")

(defn second< [x y]
  (let [c (compare (second x) (second y))]
    (if (not= c 0)
      c
      (compare x y))))

(defn parse-input [txt]
  (->> txt
       (str/split-lines)
       (map #(str/split % #":\s+"))
       (map #(vector (first %) (str/split (second %) #"\s+")))
       (into {})))

(defn assoc-reverse [values key dict]
  (if (empty? values) dict
      (assoc-reverse (rest values) key
                     (assoc dict (first values)
                            (conj (dict (first values)) key)))))

(defn complete-map [node-map complete]
  (if (empty? node-map)
    complete
    (let [[key vals] (first node-map)]
      (complete-map (dissoc node-map key) 
                    (assoc-reverse vals key complete)))))

(defn seen-nodes [to-see node-map seen-dist acc]
  (if (> acc 1000000) "FUCK" 
      (let [[node dist] (first to-see)
            node-parents (node-map node)
            n-to-see (->> node-parents
                          (filter #(not (contains? seen-dist %)))
                          (map #(vector % (inc dist))))]
        (if (empty? to-see) seen-dist
            (seen-nodes (into (disj to-see [node dist]) n-to-see)
                        node-map
                        (assoc seen-dist node dist)
                        (inc acc))))))

(defn solve [inpt]
  (let [parsed-map (parse-input inpt)
        complete-map (complete-map parsed-map parsed-map)
        seen-map (->> (keys complete-map)
                      (map #(seen-nodes (sorted-set-by second< [% 0]) complete-map {} 0)))]
    
    (->> (for [i seen-map]
           (let [sorted (sort-by second i)
                 max (second (last sorted))
                 last-entries (filter #(< (second %) (- max (/ max 1.5))) sorted)]
           (println sorted)
             last-entries))
         (reduce concat)
         (map first)
         (frequencies)
         (sort-by second)
         )))

;(solve sample)

(let [parsed-map (parse-input input)
      comp-map (complete-map parsed-map parsed-map)]

  (for [i ["ptn" "jbv"]]   
    (->> (seen-nodes (sorted-set-by second< [i 0]) comp-map {} 0)
         (count))))

(* 767 728)

;9 and 6
;hfx pzl
;bvb cmg
;nvd jqt

