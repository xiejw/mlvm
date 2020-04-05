#include "mlvm/Op/Einsum.h"

#include "absl/container/flat_hash_map.h"

namespace mlvm::OP {

void EinsumHelper::makePlan(std::vector<char> lhs, std::vector<char> rhs,
                            std::vector<char> output) {
  // Stage: Collecting statistic.
  absl::flat_hash_map<char, int> dim_count{};
  absl::flat_hash_map<char, int> lhs_dim_count{};
  absl::flat_hash_map<char, int> rhs_dim_count{};
  absl::flat_hash_map<char, int> output_dim_count{};

  std::cout << "LHS: ";
  for (auto& dim : lhs) {
    std::cout << dim << " ";
    lhs_dim_count[dim] += 1;
    dim_count[dim] += 1;
  }
  std::cout << "\n";

  std::cout << "RHS: ";
  for (auto& dim : rhs) {
    std::cout << dim << " ";
    rhs_dim_count[dim] += 1;
    dim_count[dim] += 1;
  }
  std::cout << "\n";

  std::cout << "OUPPUT: ";
  for (auto& dim : output) {
    std::cout << dim << " ";
    output_dim_count[dim] += 1;
  }
  std::cout << "\n";

  for (auto& re : lhs_dim_count) {
    std::cout << re.first << " " << re.second << "\n";
  }
  std::cout << "\n";

  for (auto& re : rhs_dim_count) {
    std::cout << re.first << " " << re.second << "\n";
  }
  std::cout << "\n";

  for (auto& re : output_dim_count) {
    std::cout << re.first << " " << re.second << "\n";
  }
  std::cout << "\n";

  // int num_inputs = 2;
  for (auto& re : dim_count) {
    auto& dim = re.first;
    bool disapper = output_dim_count[dim] == 0;
    bool unique = (lhs_dim_count[dim] == 0 || rhs_dim_count[dim] == 0);

    if (!unique and !disapper) {
      std::cout << dim << " batch \n";
    } else if (!unique) {  // disapper
      std::cout << dim << " contract\n";
    } else if (!disapper) {  // unique
      std::cout << dim << " free\n";
    } else {
      std::cout << dim << " reduce\n";
    }
  }
};

}  // namespace mlvm::OP
