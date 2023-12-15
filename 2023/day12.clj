(require '[clojure.string :as str]
         '[clojure.math.combinatorics :as combo]
         '[taoensso.tufte :as tufte :refer (defnp p profile)])

(tufte/add-basic-println-handler! {:format-pstats-opts {:columns [:n-calls :max :mean :mad :clock :total]}})


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
         (filter #(is-valid % brokens)))))

(defn parse [sample]
  (->> sample
       (#(str/split % #"\n"))
       (map #(str/split % #" "))
       (map #(vector (first %)
                     (map read-string (str/split (second %) #","))))))

(comment 
  (time (->> input
           (parse)
           (map permutations-missing-parts)
           (map count)
           (reduce +)))
)

;p2 
(defn times-5 [line]
  (->> (map #(repeat 5 %) line)
       (#(vector (str/join "?" (first %))
                 (flatten (second %))))))

(defn make-str-poss [gap-size i]
             (apply str (concat (repeat i \.) (repeat gap-size \#) [\.])))

(defn validate-substr [substr curr-string]
             (->> (map vector substr (take (count substr) curr-string))
                  (every? #(or (= (first %) (second %)) (= \? (second %))))))

(def count-springs
  (memoize (fn [curr-string gaps]
             (if (empty? gaps) (if (every?
                                    #(or (= \. %) (= \? %))
                                    curr-string)
                                 1 0)
                 (let [curr-gap (first gaps)
                       rest-gaps (rest gaps)
                       next-size-needed (+ (reduce + rest-gaps) (- (count rest-gaps) 1))
                       rng (range (- (count curr-string) next-size-needed curr-gap))
                       possible-sp-arr (map #(make-str-poss curr-gap %) rng)]
                   (->> (filter #(validate-substr % curr-string) possible-sp-arr)
                        (map #(drop (count %) curr-string))
                        (pmap #(count-springs % rest-gaps))
                        (apply +)))))))

(comment (->> input
     (parse)
     (map times-5)
     (pmap #(count-springs (first %) (second %)))
     (reduce +)))

(comment
  (defnp count-one [] (->> '("?###????????" (3,2,1))
                           (times-5)
                           (#(count-springs (first %) (second %)))))
  (profile {} 
           (dotimes [_ 1] 
             (count-one)))

  (defn count-sample [] (->> input
                             (parse)
                             (take 20)
                             (map times-5)
                             (pmap #(count-springs (first %) (second %)))))
  (profile {}
           (dotimes [_ 1]
             (count-sample)))

  (permutations-missing-parts '("???.###" (1 1 3)))
  (is-valid "#.#.###" '(1 1 3))
  (parse sample))
