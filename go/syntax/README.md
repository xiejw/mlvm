User Code

```
(defn @batch  3)
(defn @hidden_size 2)

(defn rng (prng_create 456))

(defn a (
    prng_norm rng \[@batch  @hidden_size]))
(defn b (
    tensor
        \[@batch @hidden_size]
        [1.0 2.0 3.0 4.0 5.0 6.0]))

(+ a b)
```

Language Design
- Type scheme deduction for dimension.
- Functional.
