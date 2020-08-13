### User Code 1

```
(defn @batch  3)
(defn @hidden_size 2)
(defn s [@batch, @hidden_size])

(defn key_tuple (rng_split (rng_new 456)))

(defn key_1 (nth 0 key_tuple))
(defn key_2 (nth 1 key_tuple))

(defn a (rng_norm key_1 s))
(defn b (rng_norm key_2 s))
(defn c
    (tr_reshape
        (tr_new [1.0 2.0 3.0 4.0 5.0 6.0])
        [@batch @hidden_size]
    )
)

(* (+ a b) c)
```

Language Design
- Type scheme deduction for dimension.
- Functional.

Basic Types
- List<a>
- Dim

### User Code 2

```
(= @batch  3)
(= @hidden_size 2)
(= s [@batch, @hidden_size])

(= key_tuple (rng_split (rng_new 456)))

(= key_1 (nth 0 key_tuple))
(= key_2 (nth 1 key_tuple))

(= a (rng_norm key_1 s))
(= b (rng_norm key_2 s))

(= r (+ a b))

(= t_list (IO read "weights")) // [Object]

(= w::Tensor<@hidden_size> (nth 0 t_list)) // Tensor<@hidden_size>

(IO write "weight" (list::[Object] (+ a w)))

```

Type checking at runtime.
Type in parser and inference.
Pure vs inpure.

