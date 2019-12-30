#ifndef MLVM_COMPUTATION_PROGRAM_
#define MLVM_COMPUTATION_PROGRAM_

#include "mlvm/Computation/Function.h"
#include "mlvm/Foundation/Status.h"
#include "mlvm/Foundation/StatusOr.h"

#include <string>
#include <unordered_map>

namespace mlvm::computation {

class Program {
 public:
  Program(std::string name) : name_{name} {};

 public:
  Function* makeFunc(std::string fn_name) {
    auto fn = new Function{fn_name};
    fns_[fn_name] = std::unique_ptr<Function>{fn};
    // Check key dost noe exist.
    return fns_[fn_name].get();
  }

 private:
  std::unordered_map<std::string, std::unique_ptr<Function>> fns_;
  std::string name_;
};

}  // namespace mlvm::computation

#endif
