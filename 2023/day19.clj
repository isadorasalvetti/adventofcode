(require '[clojure.string :as str])
(def input (slurp "input/day19.txt"))

(def sample "px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}")

(def start "in")

(def <> {"<" < ">" >})

(defn parse-rule [rule]
  (let [[r d] (str/split rule #":")
        [right left] (str/split r #"[<>]")
        sign (first (re-seq #"[<>]" rule))]
    [d right sign (read-string left)]))

(defn parse-instruct [line]
  (let [[name rst _] (str/split line #"\{")
        rules (str/split rst #",")
        p-rules (conj
                 (apply vector (map parse-rule (butlast rules)))
                 (vector (apply str (butlast (last rules)))))]
    [name p-rules]))

(defn parse-gear [line]
  (->> (str/replace line #"[\{\}]" "")
       (#(str/split % #","))
       (map #(str/split % #"="))
       (map #(vector (first %) (read-string (second %))))
       (into {})))

(defn parse-full [in]
  (let [[inst g] (str/split in #"\n\n")
        instructions (map
                      parse-instruct
                      (str/split inst #"\n"))
        gears (map
               parse-gear
               (str/split g #"\n"))]
    [{"in" gears} (into {} instructions)]))

(some? nil)

(defn process-gear [gear instructions]
  (let [[dest g op num] (first instructions)]
    ;(println "Processing" gear instructions)
    (if (or (not (some? g))
            ((<> op) (gear g) num)) dest
        (process-gear gear (rest instructions)))))

(defn assoc-join [mp k v]
  (let [curr (mp k)]
    (if (some? curr)
      (assoc mp k (conj curr v))
      (assoc mp k [v]))))

(defn assoc-join-col [mp k v]
  (let [curr (mp k)]
    (if (some? curr)
      (assoc mp k (into curr v))
      (assoc mp k v))))

(defn do-bucket [gears inst-line acc]
  (if (empty? gears)
    acc
    (do-bucket (rest gears)
               inst-line
               (assoc-join acc
                           (process-gear (first gears) inst-line)
                           (first gears)))))

(defn merge-buckets [new-buckets gears complete]
  (if (empty? new-buckets) [gears complete]
      (let [[cb cont] (first new-buckets)]
        (if (or (= cb "A") (= cb "R"))
          (merge-buckets (dissoc new-buckets cb) gears
                         (assoc-join-col complete cb cont))
          (merge-buckets (dissoc new-buckets cb)
                         (assoc-join-col gears cb cont)
                         complete)))))

(defn do-all [gears instructions complete]
  (if (empty? gears) complete
      (let [[bucket gs] (first gears)
            new-buckets (do-bucket gs (instructions bucket) {})
            [merged-gears merged-complete]
            (merge-buckets new-buckets (dissoc gears bucket) complete)]
        (do-all merged-gears instructions merged-complete))))

(let [[gears instructions] (parse-full input)]
  (->> ((do-all gears instructions {}) "A")
       (map vals)
       (flatten)
       (reduce +)))

(comment
  (assoc-join {} 1 1)
  (process-gear
   {"x" 787, "m" 2655, "a" 1222, "s" 2876}
   [["qkq" "a" "<" 1006] ["A" "m" ">" 3090] ["rfg"]])
  (parse-full sample)
  (parse-rule "x<1416:A")
  (parse-instruct "qkq{x<1416:A,crn}")
  (parse-gear "{x=787,m=2655,a=1222,s=2876}"))


; P2

(defn make-split-range [gear-ranges g op num]
  (let [[a b] (gear-ranges g)
        [pass fail] (if (= op "<")
                      [[a (- num 1)] [num b]]
                      [[num b] [a (+ num 1)]])]
    [(assoc gear-ranges g pass) (assoc gear-ranges g fail)]))

(def starting-gear {(str "x") [1 4000], (str "m") [1 4000], (str "a") [1 4000], (str "s") [1 4000]})
(defn assoc-range [gear-ranges instructions]
  (let [[dest g op num] (first instructions)]
    ;(println "Processing" gear-ranges instructions)
    (if (not (some? g)) [{dest [gear-ranges]}]
        (let [[s l] (gear-ranges g)]
          (cond
            (or (and (= op ">") (> s num))  ; all passes
                (and (= op "<") (< l num))) [{dest [gear-ranges]}]
            (or (and (= op ">") (< l num))
                (and (= op "<") (> s num))) (assoc-range gear-ranges (rest instructions))
            :else (let [[succ fail] (make-split-range gear-ranges g op num)]
                    (into {dest [succ]} (assoc-range fail (rest instructions)))))))))

(reduce into [{1 {2 3}} {2 {2 3}}])

(defn do-all-ranges [gear-buckets instructions complete]
  ;(println gear-buckets)
  (if (empty? gear-buckets) complete
      (let [[bucket gs] (first gear-buckets)
            new-buckets (reduce into
                                (map #(assoc-range % (instructions bucket)) gs))
            ;a (println new-buckets "Found")
            [merged-gears merged-complete]
            (merge-buckets new-buckets (dissoc gear-buckets bucket) complete)
            ;a (println merged-gears "Merged")
            ]
        (do-all-ranges merged-gears instructions merged-complete))))

(let [[_ instructions] (parse-full input)]
  (->> ((do-all-ranges {"in" [starting-gear]} instructions {}) "A")
       (map vals)
       (map #(map (partial apply -) %))
       (map (partial reduce *))
       (reduce +)
       ))


(comment
  (first (assoc-range starting-gear [["A" "x" "<" 3000] ["R"]])))



