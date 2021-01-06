#include "program.h"

#include "stdlib.h"

programt* pgCreate()
{
        programt* pg     = malloc(sizeof(programt));
        pg->instructions = NULL;
        return pg;
}

void pgFree(programt* pg)
{
        vecFree(pg->instructions);
        free(pg);
}
