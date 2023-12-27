(require '[clojure.string :as str])
(def input (slurp "input/day20.txt"))

(def sample "broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a")

(def sample2 "broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output")

(defn conj-outputs [conj nodes]
  (let [ftr (fn [conj [_ [_ ts]]] (some #(= % conj) ts))]
    (->> (filter #(ftr conj %) nodes)
         (map first))))

(defn default-data [nodes name]
  (case (first (nodes name))
    \% false
    \& (into {} (map vector (conj-outputs name nodes) (repeat true)))))

(defn parse-line [line]
  (let [[name wires] (str/split line #"\s+->\s+")]
    {(apply str (rest name)) [(first name) (str/split wires #", ")]}))

(defn parse [text]
  (->> (str/split text #"\n")
       (map parse-line)
       (reduce into)))
       

(defn make-starting-data [nodes]
  (->> nodes
       (keys) 
       (remove #{"roadcaster"})
       (map #(hash-map % (default-data nodes %)))
       (into {})))
       

(defn apply-signal [source node signal data]
  (let [[prefix node-name] node] 
   (case prefix
     \b [data signal]
     nil [data nil]
     \% (if signal
          (let [curr-state (data node-name)]
            [(assoc data node-name (not curr-state)) curr-state])
          [data nil])
     \& (let [n-data (assoc-in data [node-name source] signal)] 
          [n-data (if (some identity (vals (n-data node-name))) false true)]))))

(defn process-signal [to-do nodes data lows highs nodes-of-interest, count]
  (if (nil? to-do) [lows highs data nodes-of-interest]
      (let [[l node signal] (last to-do)
            [p targets] (nodes node)
            [updated-data next-signal] (apply-signal l [p node] signal data)
            acch (if signal highs (inc highs))
            accl (if signal (inc lows) lows)
            next-todo (if (some? next-signal)
                        (into (or (butlast to-do) (list))
                              (map #(vector node % next-signal) targets))
                        (butlast to-do))
            ;P2
            convs (for [[node _] nodes-of-interest
                        :when (first (map second (data node)))]
                    [node count])
            noi (into nodes-of-interest convs)]
        
        (cond
          (> (+ acch accl) 10000000) "FUCK"
          :else (process-signal next-todo nodes updated-data accl acch noi, count)))))

(def button-presses 5000)

(defn push-go [nodes data counter nodes-of-interest]
  (loop [data data counter counter lows 0 highs 0 noi nodes-of-interest]
    (let [todo (- counter 1)
          [l h data noi] (process-signal (list [nil "roadcaster" true]) nodes data lows highs noi (- button-presses todo))] 

      (if (> todo 0)
        (recur data todo l h noi)
        [l h noi]))))

(defn conected-into-nodes [nodes tnode]
  (for [[node [p targets]] nodes 
        :when (some #(= % tnode) targets)] 
    node))

(comment
  (let [smp (parse input)
        starting-data (make-starting-data smp)
        nodes-of-interest (conected-into-nodes smp "bq")
        noi (into {} (for [node nodes-of-interest] [node -1]))]
    (->> (push-go smp starting-data button-presses noi)
         (last)
         (vals)
         (apply *)))
  
  ;P1
  (let [smp (parse input)
        starting-data (make-starting-data smp)]
    (->> (push-go smp starting-data 1000 [])
         (butlast)
         (apply *))))

  



