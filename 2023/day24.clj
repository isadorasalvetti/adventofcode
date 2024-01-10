(require '[clojure.string :as str] 
         '[clojure.math.combinatorics :as combo])
(def input (slurp "input/day24.txt"))

(def sample "19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3")

(defn parse-txt [txt]
  (->> (str/split txt #"\n")
       (mapcat #(str/split % #"\s+@\s+"))
       (map #(str/split % #",\s+"))
       (map #(map read-string %))
       (map (partial apply vector))
       (partition 2)))
(parse-txt sample)

(defn det [a b] (- (* (first a) (second b)) (* (second a) (first b))))

;p1 + t1*d1 = p2 + t2*d2

;t1*dx1 - t2*dx2 + x1-x2 = 0
;t1*dy1 - t2*dy2 + y1-y2 = 0


(defn has-intersection? [min max l1 l2]
  (let [[[x1 y1 _] [dx1 dy1 _]] l1
        [[x2 y2 _] [dx2 dy2 _]] l2
        dx3 (- x1 x2)
        dy3 (- y1 y2)

        det0 (det [dx1 dx2] [dy1 dy2])
        det1 (det [dx1 dx3] [dy1 dy3])
        det2 (det [dx2 dx3] [dy2 dy3])] 
    
    (if (= det0 0) false
        (let [s (/ det1 det0)
              t (/ det2 det0)
              ix (float (+ x1 (* t dx1)))
              iy (float (+ y1 (* t dy1)))] 
          ;(println l1 l2)
          ;(println ix iy (and (>= t 0) (<= min ix max) (<= min iy max)))
          (and (>= t 0) (>= s 0) (<= min ix max) (<= min iy max))
          ))))

(defn solver_coefs [p1 p2 p3 p4]
  (let [[[x1 y1 z1] [dx1 dy1 dz1]] p1
        [[x2 y2 z2] [dx2 dy2 dz2]] p2
        [[x3 y3 z3] [dx3 dy3 dz3]] p3
        [[x4 y4 z4] [dx4 dy4 dz4]] p4] 
    
    (for [[x1 y1 dx1 dy1] [[x1 y1 dx1 dy1] [x2 y2 dx2 dy2]]
            [x2 y2 dx2 dy2] [[x3 y3 dx3 dy3] [x4 y4 dx4 dy4]]]
        (println (+ y1 (- y2)) ;"A" ;DX
                 (+ (- x1) x2) ;"B" ;DY
                 0
                 (+ (- dy1) dy2) ;"X"
                 (+ dx1 (- dx2)) ;"Y"
                 0
                 (+ (* (- dy1) x1) (* y1 dx1) (* dy2 x2) (* (- y2) dx2))))

      (for [[x1 z1 dx1 dz1] [[x1 z1 dx1 dz1] [x2 z2 dx2 dz2]]
            [x2 z2 dx2 dz2] [[x3 z3 dx3 dz3] [x4 z4 dx4 dz4]]]
        (println (+ z1 (- z2)) ;"A" ;DX
                 0
                 (+ (- x1) x2) ;"B" ;DZ
                 (+ (- dz1) dz2) ;"X"
                 0
                 (+ dx1 (- dx2)) ;"Z"
                 (+ (* (- dz1) x1) (* z1 dx1) (* dz2 x2) (* (- z2) dx2))))
    ))
  ; 24, 13 / -3, 1

  ;DX*y1 - x1*DY - X*dy1 + dx1*Y + dy1*x1 - y1*dx1
  ;(-) DX*y2 - x2*DY - X*dy2 + dx2*Y + dy2*x2 - y2*dx2
  
(comment
  ;P2
  (->> (parse-txt input)
       (take 4)
       (apply solver_coefs))

  (+ 446533732372768 293892176908833 180204909018503)

  ;P1
  (->> (parse-txt input)
       (#(combo/combinations % 2))
       (map #(has-intersection? min-area max-area (first %) (second %)))
       (filter identity)
       (count)))


;p1 + t1*d1 = p2 + t1*d2
;t1*d2 - t1*d1 = p1 - p2
;d2 - d1 = p1-p2 / t1
;t1 = p1-p2 / d2-d1

;dx*t1 + x - dx1*t1 - x1 =0
;dy*t1 + y - dy1*t1 - y1 =0
;dz*t1 + z - dz1*t1 - z1 =0

;(dx-dx1) t1 = x1 - x
;t1 = (x1 - x) / (dx - dx1)

;(x1 - X) / (DX - dx1) = (y - y1) / (DY - dy1)
;(DX - dx1)*(-Y + y1) = (x1 - X)*(DY - dy1)
;- DX*Y + DX*y1 + dx1*Y - y1*dx1 = -X*DY +X*dy1 +x1*DY - dy1*x1
;X*DY - DX*Y = DX*y1 + dx1*Y - X*dy1 - y1*dx1 - x1*DY + dy1*x1 

;DX*Y - X*DY = DX*y1 - DY*x1 - X*dy1 + Y*dx1 + y1*dx1 - dy1*x1
;DX*Y - X*DY = DX*y2 - DY*x2 - X*dy2 + Y*dx2 + y2*dx2 - dy2*x2

;DX*(y1-y2) - DY*(x1-x2) - X*(dy1-dy2) + Y*(dx1-dx2) + y1*dx1 - dy1*x1 = 0


;p1 + t1*d1 = px + t1*dx
;p2 + t2*d2 = px + t2*dx
;p3 + t3*d3 = px + t3*dx
;...

;p1 + t1*d1 = px + t1*dx

(def min-area 200000000000000)
(def max-area 400000000000000)
