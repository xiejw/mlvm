#ifndef MLVM_RANDOM_SPLITTABLE_H_
#define MLVM_RANDOM_SPLITTABLE_H_

#include <cstdint>

namespace mlvm::Random {

class Splittable64 {
 public:
  using int64 = int64_t;

  Splittable64(int64 seed) : seed_{seed} { (void)seed_; };

 private:
  int64 seed_;
};

}  // namespace mlvm::Random

#endif
