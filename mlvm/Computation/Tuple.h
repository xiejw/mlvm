#ifndef MLVM_COMPUTATION_TUPLE_
#define MLVM_COMPUTATION_TUPLE_

#include <memory>
#include <optional>
#include <string>
#include <vector>

#include "mlvm/Array/ShapeLike.h"
#include "mlvm/Foundation/Status.h"

namespace mlvm::computation {

struct Item {
  array::Shape shape;
  std::optional<std::string> name;
};

class Tuple {
 public:
  foundation::Status Add(array::ShapeLike shape_like);

 private:
  std::vector<std::unique_ptr<Item>> items_;
};

}  // namespace mlvm::computation

#endif
