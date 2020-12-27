#include "stdio.h"

#include "adt/sds.h"

int main() {
  sds_t s = sdsNew("hello mlvm\n");
  printf("%s\n", s);
  sdsFree(s);
}
