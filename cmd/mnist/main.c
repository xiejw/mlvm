#include <stdio.h>

#include "mnist.c"

int main()
{
        unsigned char* images = NULL;
        unsigned char* labels = NULL;
        error_t        err    = OK;

        err = readMnistTrainingImages(&images);
        if (err) {
                goto clean;
        }

        printMnistImage(images);
        printMnistImage(images + 28 * 28);

        err = readMnistTrainingLabels(&labels);
        if (err) {
                goto clean;
        }
        printf("label %d\n", (int)*labels);
        printf("label %d\n", (int)*(labels + 1));

clean:
        if (images != NULL) free(images);
        if (labels != NULL) free(labels);
        return err;
}
