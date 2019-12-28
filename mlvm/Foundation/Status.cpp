#include "mlvm/Foundation/Status.h"

namespace mlvm::foundation {

const Status Status::OK = Status({}, {});
const Status Status::InvalidArguments =
    Status(ErrorCode::INVALID_ARGUMENTS, {});

}  // namespace mlvm::foundation
