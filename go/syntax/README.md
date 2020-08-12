User Code

```
(defn @batch  3)
(defn @hidden_size 2)

(defn key_tuple (rng_split (rng_new 456)))

(defn key_1 (nth 0 key_tuple))
(defn key_2 (nth 1 key_tuple))

(defn a (rng_norm key_1 \[@batch  @hidden_size]))
(defn b
    (rng_norm key_2 \[@batch  @hidden_size]
    )
)
(defn c
    (tr_reshape
        (tr_new [1.0 2.0 3.0 4.0 5.0 6.0])
        \[@batch @hidden_size]
    )
)

(* (+ a b) c)
```

Language Design
- Type scheme deduction for dimension.
- Functional.
