(require '[clojure.string :as str])
(require '[clojure.math.combinatorics :as combo])

(def input (slurp "input/day12.txt"))
(def sample "???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1")

(defn rep-with [str-line seq-to-replace]
  (if (empty? seq-to-replace)
      str-line
      (-> str-line
          (str/replace-first \? (first seq-to-replace))
          (rep-with (rest seq-to-replace)))))

(defn is-valid [line brokens]
  (->> line
       (re-seq #"#+")
       (map count)
       (= brokens)))

(defn permutations-missing-parts [[line-text brokens]]
  (let [total_broken (reduce + brokens)
        curr_broken (count (re-seq #"#" line-text))
        empty-spaces (count (re-seq #"\?" line-text))
        broken-to-add (- total_broken curr_broken)
        dots-to-add (- empty-spaces broken-to-add)
        permutation-base (concat (repeat broken-to-add \#) 
                               (repeat dots-to-add \.))]
    (->> permutation-base
         (combo/permutations)
         (map #(rep-with line-text %))
         (filter #(is-valid % brokens))
         )
    ))

(defn parse [sample]
  (->> sample
       (#(str/split % #"\n"))
       (map #(str/split % #" "))
       (map #(vector (first %)
                     (map read-string (str/split (second %) #","))))))

(->> input
     (parse)
     (map permutations-missing-parts)
     (map count)
     (reduce +))

(comment
  (permutations-missing-parts '("???.###" (1 1 3)))
  (is-valid "#.#.###" '(1 1 3)) 
  (parse sample)
  )