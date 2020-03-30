#ifndef MLVM_FOUNDATION_LOGGING_H_
#define MLVM_FOUNDATION_LOGGING_H_

#include <iostream>

#define LOG_IS_ON(level) \
  (static_cast<int>(level) <= mlvm::logging::Logger::currentLevel())

#define LOG_INFO()                                 \
  (!LOG_IS_ON(mlvm::logging::Logger::Level::Info)) \
      ? (void)0                                    \
      : mlvm::logging::VoidType::instance&         \
        mlvm::logging::Logger::getCurrentLogger()

#define LOG_FLUSH() mlvm::logging::Logger::getCurrentLogger().flush();

namespace mlvm::logging {

class Logger {
 public:
  enum class Level : int {
    All = -99,
    Fatal = -2,
    Error = -1,
    Info = 0,
  };

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

class VoidType {
 public:
  static VoidType instance;

  inline void operator&(Logger& logger) { (void)logger; };
};

}  // namespace mlvm::logging

#endif
