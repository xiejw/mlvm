#ifndef MLVM_FOUNDATION_LOGGING_H_
#define MLVM_FOUNDATION_LOGGING_H_

#include <iostream>
#include <type_traits>

#include "mlvm/Foundation/Logger.h"

// Checks with the `level` is on.
#define LOG_IS_ON(level) ((level) <= mlvm::LoggerManager::currentLevel())

// `level` must be int type (static checked).
#define LOG(level)                                          \
  static_assert(std::is_same<decltype(level), int>::value); \
  (!LOG_IS_ON(level)) ? (void)0                             \
                      : mlvm::logging::VoidType::instance&  \
                        mlvm::LoggerManager::currentLogger()

#define LOG_INFO() LOG(static_cast<int>(mlvm::logging::Level::Info))
#define LOG_DEBUG() LOG(static_cast<int>(mlvm::logging::Level::Debug))

namespace mlvm {

// Should be created once, most likely in the `main()`.
//
// Deconstruction will flush the logger.
class LoggerManager {
 public:
  LoggerManager(int argc, char** argv) {
    (void)argc;
    (void)argv;
  };

  ~LoggerManager() { currentLogger().flush(); }

 public:
  static logging::Logger& currentLogger() {
    static bool has_lines_emitted = false;
    if (has_lines_emitted) {
      // Append the end-of-line first.
      loggerInstance.log("\n");
    }
    has_lines_emitted = true;
    return loggerInstance;
  }

  // Returns the current debugging level.
  static int currentLevel() { return 0; }

 protected:
  static logging::Logger loggerInstance;
};

}  // namespace mlvm

#endif
