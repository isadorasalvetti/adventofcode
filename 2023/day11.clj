(require '[clojure.string :as str])
(def input (->> (slurp "input/day11.txt") (#(str/split % #"\n"))))
(def sample (->> "...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#....."
                 (#(str/split % #"\n"))))

(defn space-size [space]
  [(count (first space)) (count space)]
  )

(defn apply-first-val [val col]
  (map #(concat [val] %) col))

(defn parse-to-map [string]
  (->> string 
       (map #(map-indexed vector %))
       (map-indexed vector)
       (map #(apply apply-first-val %))
       (reduce concat)
       (filter #(= (last %) \#))
       (map #(vector (take 2 %) (last %)))
       (into {})
       ))

(defn is-empty-y [y space-keys]
  (not (some #(= y (second %)) space-keys)))

(defn is-empty-x [x space-keys]
  (not (some #(= x (first %)) space-keys)))

(defn make-pairs [space]
  (->> space
       (keys)
       (#(for [[i a] (map-indexed vector %)
              [j b] (map-indexed vector %)
              :when (< i j)] 
           [a b]))
       ))

(defn get-all-empties [func keys a b] 
  (->> (sort [a b])
       (apply range) 
       (map-indexed #(vector %1 (func %2 keys)))
       (filter #(= true (second %)))
       (map first)
       ))

(defn get-distance [[xa ya] [xb yb] empty-cols empty-rows]
  (let [[xa xb] (sort [xa xb])
        [ya yb] (sort [ya yb])
        dist [(- xb xa) (- yb ya)]
        x-spaces (count (filter #(and (< % xb) (> % xa)) empty-rows))
        y-spaces (count (filter #(and (< % yb) (> % ya)) empty-cols))]
    (->> (map + [(* x-spaces (- 1000000 1)) (* y-spaces (- 1000000 1))] dist)
         (reduce +))))

(let [parsed-space (parse-to-map input)
      [sizex sizey] (space-size input)
      parsed-keys (keys parsed-space)
      pairs-to-compute (make-pairs parsed-space)
      empty-cols (get-all-empties is-empty-y parsed-keys 0 sizey)
      empty-rows (get-all-empties is-empty-x parsed-keys 0 sizex)]
  (->> pairs-to-compute
       (map #(get-distance (first %) (second %) empty-cols empty-rows))
       (reduce +))
  )


