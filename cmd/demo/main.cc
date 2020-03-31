#include <memory>

#include "mlvm/Foundation/Logging.h"
#include "mlvm/Foundation/Macros.h"
#include "mlvm/Foundation/StatusOr.h"
#include "mlvm/IR/Function.h"
#include "mlvm/Runtime/Evaluator.h"

// Builds a IR::Function.
mlvm::StatusOr<std::unique_ptr<mlvm::IR::Function>> buildFunction() {
  auto fn = std::make_unique<mlvm::IR::Function>("main");
  auto c0 = fn->newConst(mlvm::IR::ConstTensor{mlvm::IR::Array{"c0"}});
  auto ins = fn->newInstruction(mlvm::IR::OpType::Add, std::vector{c0, c0});
  auto o0 = ins->outputAt(0);
  ins = fn->newInstruction(mlvm::IR::OpType::Add, std::vector{o0, o0});

  MLVM_RETURN_IF_ERROR(fn->setOutput(ins->outputAt(0)));
  return fn;
}

int main(int argc, char** argv) {
  mlvm::LoggerManager mgr{argc, argv, /*parse_command_line=*/true};

  LOG_INFO() << "Build function.";
  MLVM_ASSIGN_OR_FATAL(auto fn, buildFunction());

  mlvm::RT::Evaluator eval{};
  LOG_INFO() << "Run function.";
  MLVM_FATAL_IF_ERROR(eval.run(*fn));
}
