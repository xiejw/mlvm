#ifndef SUPPORT_ERROR
#define SUPPORT_ERROR

namespace mlvm {
namespace support {

[[noreturn]] void FatalError(const char *fmt, ...);

}  // namespace support
}  // namespace mlvm

#endif
