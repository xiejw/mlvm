#ifndef MLVM_COMPUTATION_TUPLE_
#define MLVM_COMPUTATION_TUPLE_

#include <optional>
#include <string>
#include <memory>
#include <vector>

#include "mlvm/Array/Shape.h"
#include "mlvm/Foundation/Status.h"

namespace mlvm::computation {

struct Item {

  //array::
  std::optional<std::string> name;
};


class Tuple {
  public:

    // foundation::Status Add(std::initialize_list<unsigned int> shape_l) {
    //   ASSIGN_OR_RETURN(shape, array:Shape:

    // }

  private:
    std::vector<std::unique_ptr<Item>> items_;
};


}  // namespace mlvm::computation

#endif
