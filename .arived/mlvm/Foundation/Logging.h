#ifndef MLVM_FOUNDATION_LOGGING_H_
#define MLVM_FOUNDATION_LOGGING_H_

#include <libgen.h>    // For basename
#include <sys/time.h>  // For time related parts.
#include <iostream>
#include <type_traits>  // For std::is_same

#include "absl/flags/parse.h"

#include "mlvm/Foundation/Logger.h"

// Checks with the `level` is on.
#define LOG_IS_ON(level) ((level) <= mlvm::LoggerManager::currentLevel())

// `level` must be int type (static checked).
#define LOG(level)                                                             \
  static_assert(std::is_same<decltype(level), int>::value);                    \
  (!LOG_IS_ON(level))                                                          \
      ? (void)0                                                                \
      : mlvm::logging::VoidType::instance& mlvm::LoggerManager::currentLogger( \
            __FILE__)

// Helpers
#define LOG_INFO() LOG(static_cast<int>(mlvm::logging::Level::Info))
#define LOG_DEBUG() LOG(static_cast<int>(mlvm::logging::Level::Debug))

namespace mlvm {

// Should be created once, most likely in the `main()`.
//
// Deconstruction will flush the logger.
class LoggerManager {
 public:
  LoggerManager(int argc, char** argv, bool parse_command_line) {
    if (parse_command_line) absl::ParseCommandLine(argc, argv);
  };

  ~LoggerManager() { currentLogger().flush(); }

 public:
  static logging::Logger& currentLogger(const char* path = nullptr) {
    static bool has_lines_emitted = false;
    if (has_lines_emitted) {
      // Append the end-of-line first.
      loggerInstance.log("\n");
    }
    has_lines_emitted = true;

    if (path != nullptr) {
      // Prints the
      timeval curTime;
      gettimeofday(&curTime, NULL);
      int milli = curTime.tv_usec / 1000;

      char buffer[80];
      strftime(buffer, 80, "%Y-%m-%d %H:%M:%S", localtime(&curTime.tv_sec));

      char currentTime[84] = "";
      sprintf(currentTime, "%s:%03d", buffer, milli);
      loggerInstance << currentTime << " " << basename((char*)path) << " ";
    }
    return loggerInstance;
  }

  // Returns the current debugging level.
  static int currentLevel();

 protected:
  static logging::Logger loggerInstance;
};

}  // namespace mlvm

#endif
