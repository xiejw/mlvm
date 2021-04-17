#include <stdio.h>

#include "mnist.c"

static error_t readMnistData(unsigned char** images, unsigned char** labels);

int main()
{
        unsigned char* images = NULL;
        unsigned char* labels = NULL;
        error_t        err    = OK;

        err = readMnistData(&images, &labels);
        if (err) {
                goto clean;
        }

clean:
        if (images != NULL) free(images);
        if (labels != NULL) free(labels);
        return err;
}

// impl
error_t readMnistData(unsigned char** images, unsigned char** labels)
{
        error_t err = readMnistTrainingImages(images);
        if (err) {
                return err;
        }

        err = readMnistTrainingLabels(labels);
        if (err) {
                return err;
        }
        printf("sample label %d -- image:\n", (int)**labels);
        printMnistImage(*images);
        printf("smaple label %d -- image:\n", (int)*(*labels + 1));
        printMnistImage(*images + 28 * 28);
        return OK;
}
