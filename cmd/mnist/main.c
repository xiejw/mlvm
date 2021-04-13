#include <stdio.h>

#define MNIST_PATH "../mlvm_examples/files/"

const char* const train_images = MNIST_PATH "train-images-idx3-ubyte";
const char* const train_labels = MNIST_PATH "train-labels-idx1-ubyte";

int main()
{
        printf("hello mnist:\n    images: %s\n    labels: %s\n", train_images,
               train_labels);
}
