#include "opcode.h"
#include "sds.h"
#include "vec.h"

#include <stdio.h>

int main() {
  sds s = sdsNew("hello mlvm 123");
  printf("%s\n", s);
  sdsFree(s);
}
