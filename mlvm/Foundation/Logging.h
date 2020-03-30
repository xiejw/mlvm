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
                        mlvm::logging::Logger::getCurrentLogger()

#define LOG_INFO() LOG(static_cast<int>(mlvm::logging::Level::Info))

#define LOG_FLUSH() mlvm::logging::Logger::getCurrentLogger().flush();

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
  static Logger& getCurrentLogger() {
    static bool logged = false;
    if (logged) {
      Logger::currenLogger.log("\n");
    }
    logged = true;
    return Logger::currenLogger;
  }

  static int currentLevel() { return 0; }

 protected:
  static Logger currenLogger;

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

#endif
