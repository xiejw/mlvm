#ifndef MLVM_FOUNDATION_LOGGING_H_
#define MLVM_FOUNDATION_LOGGING_H_

#include <iostream>

#define LOG_INFO() mlvm::Logger::currenLogger

namespace mlvm {

class Logger {
 public:
  static Logger currenLogger;

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
