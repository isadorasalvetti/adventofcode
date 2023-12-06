(require '[clojure.string :as str])
(def input (slurp "input/day2.txt"))

(def sample (str "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"))

(def r-blue #"(\d+) blue")
(def r-green #"(\d+) green")
(def r-red #"(\d+) red")

(def line-to-num (fn [line r-color] 
                   (->> line 
                       (re-seq r-color ,,,) 
                       (map last ,,,) 
                       (map read-string ,,,))))

; Max 12 red, 13 green, 14 blue
(def possible-game? (fn [index line]
                      [index, (and 
                      (<= (reduce max (line-to-num line r-red)) 12)
                      (<= (reduce max (line-to-num line r-green)) 13)
                      (<= (reduce max (line-to-num line r-blue)) 14))]))

(def min-cubes (fn [line] [(reduce max (line-to-num line r-red))
                           (reduce max (line-to-num line r-green))
                           (reduce max (line-to-num line r-blue))]))


(def count-possible (fn [text] 
                      (->> (str/split text #"\n")
                           (map-indexed possible-game?)
                           (filter last)
                           (map #(+ (first %) 1))
                           (reduce +))))
;P1
(count-possible input)

(comment
  (possible-game? 1 "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
  (possible-game? 1 "Game 1: 3 blue, 4 red; 1 red, 2 green, 16 blue; 2 green")
  (line-to-num "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green" r-blue)
  (min-cubes "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")
  )