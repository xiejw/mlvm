#ifndef MLVM_FOUNDATION_LOGGING_H_
#define MLVM_FOUNDATION_LOGGING_H_

#include <iostream>
#include <type_traits>

// Checks with the `level` is on.
#define LOG_IS_ON(level) ((level) <= mlvm::logging::Logger::currentLevel())

// `level` must be int type (static checked).
#define LOG(level)                                          \
  static_assert(std::is_same<decltype(level), int>::value); \
  (!LOG_IS_ON(level)) ? (void)0                             \
                      : mlvm::logging::VoidType::instance&  \
                        mlvm::logging::Logger::currentLogger()

#define LOG_INFO() LOG(static_cast<int>(mlvm::logging::Level::Info))
#define LOG_DEBUG() LOG(static_cast<int>(mlvm::logging::Level::Debug))

// Unexposed namespace.
namespace mlvm::logging {

enum class Level : int {
  All = -99,
  Fatal = -2,
  Error = -1,
  Info = 0,
};

class Logger {
 public:
  static Logger& currentLogger() {
    static bool has_lines_emitted = false;
    if (has_lines_emitted) {
      // Append the end-of-line first.
      Logger::loggerInstance.log("\n");
    }
    has_lines_emitted = true;
    return Logger::loggerInstance;
  }

  // Returns the current debugging level.
  static int currentLevel() { return 0; }

 protected:
  static Logger loggerInstance;

 public:
  void flush() { std::cout << std::flush; }

  template <typename T>
  Logger& operator<<(T&& o) {
    log(o);
    return *this;
  }

 protected:
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

namespace mlvm {

// Should be created once, most likely in the `main()`.
class LoggerManager {
 public:
  LoggerManager(int argc, char** argv) {
    (void)argc;
    (void)argv;
  };

  ~LoggerManager() { logging::Logger::currentLogger().flush(); }
};

}  // namespace mlvm

#endif
