User Code

```
@batch = 32;
@hidden_size = 10;
@output = 1;


let prng = prng_create(seed: 456);

func create_parameter(prng: prng_t, shape: shape_t) tensor_t {
  return tensor_create(
      shape: shape,
      value: prng_norm(prng_split(prng), shape: shape))
}

let w = parameter_create(prng, shape: [@hidden_size, @output])
let b = parameter_create(prng, shape: [@output])

var i int;
for i = 0; i < 100; i++ {
}

```
