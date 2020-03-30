#ifndef MLVM_FOUNDATION_LOGGER_H_
#define MLVM_FOUNDATION_LOGGER_H_

// Unexposed namespace.
namespace mlvm::logging {

enum class Level : int {
  All = -99,
  Fatal = -2,
  Error = -1,
  Info = 0,
  Debug = 1,
};

class Logger {
 public:
  void flush() { std::cout << std::flush; }

  template <typename T>
  Logger& operator<<(T&& o) {
    log(o);
    return *this;
  }

  template <typename T>
  void log(T& o) {
    std::cout << o;
  }
};

// Help class which converts the `Logger` to `(void)` so the LOG() macro can
// work.
class VoidType {
 public:
  static VoidType instance;

 public:
  inline void operator&(Logger& logger) { (void)logger; };
};

}  // namespace mlvm::logging

#endif
