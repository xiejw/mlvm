#include "mlvm/Computation/Tuple.h"

#include <sstream>

#include "mlvm/Foundation/Macros.h"

namespace mlvm::computation {

using namespace array;
using namespace foundation;

Status Tuple::Add(ShapeLike shape_like) {
  MLVM_ASSIGN_OR_RETURN(shape, shape_like.get());

  auto item = new Item{std::move(shape)};
  items_.push_back(std::unique_ptr<Item>{item});
  return Status::OK;
}

}  // namespace mlvm::computation
