#ifndef MLVM_FOUNDATION_LOGGING_H_
#define MLVM_FOUNDATION_LOGGING_H_

#include <iostream>

#define LOG_INFO() mlvm::Logger::getCurrentLogger()
#define LOG_FLUSH() mlvm::Logger::getCurrentLogger().flush();

namespace mlvm {

class Logger {
 public:
  static Logger currenLogger;

  static Logger& getCurrentLogger() {
    static bool logged = false;
    if (logged) {
      Logger::currenLogger.Log("\n");
    }
    logged = true;
    return Logger::currenLogger;
  }

  void flush() { std::cout << std::flush; }

  template <typename T>
  void Log(T& o) {
    std::cout << o;
  }
};

template <typename T>
Logger& operator<<(Logger& logger, T&& o) {
  logger.Log(o);
  return logger;
}

}  // namespace mlvm

#endif
