#ifndef MLVM_ARRAY_SHAPE
#define MLVM_ARRAY_SHAPE

#include <memory>
#include <string>
#include <vector>

namespace mlvm {
namespace array {

// Immutable structure holding the shape information.
//
// Copy is recommended to share the Shape.
class Shape {
 public:
  explicit Shape(std::initializer_list<int> dim)
      : dims_{new std::vector<int>{dim}} {}

  const std::vector<int>& dims() const { return *dims_; };

 public:
  std::string DebugString() const;

 private:
  friend std::ostream& operator<<(std::ostream& os, const Shape& s);

 private:
  std::shared_ptr<std::vector<int>> dims_;
};

}  // namespace array
}  // namespace mlvm

#endif
