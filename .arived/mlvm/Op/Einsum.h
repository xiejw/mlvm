#ifndef MLVM_OP_EINSUM_H_
#define MLVM_OP_EINSUM_H_

#include <iostream>
#include <vector>

#include "mlvm/Foundation/Logging.h"

namespace mlvm::OP {

class EinsumHelper {
 public:
  void makePlan(std::vector<char> lhs, std::vector<char> rhs,
                std::vector<char> output);
};

}  // namespace mlvm::OP

#endif
