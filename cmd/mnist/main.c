#include <fcntl.h>  // open
#include <stdio.h>
#include <unistd.h>  // close

#define error_t int
#define OK      0

// http://yann.lecun.com/exdb/mnist/
// gzip -d -k file_name.gz
#define MNIST_PATH "../mlvm_examples/files/"
const char* const train_images = MNIST_PATH "train-images-idx3-ubyte";
const char* const train_labels = MNIST_PATH "train-labels-idx1-ubyte";
const char* const test_images  = MNIST_PATH "t10k-images-idx3-ubyte";
const char* const test_labels  = MNIST_PATH "t10k-labels-idx1-ubyte";

//
// The basic format is
//
//   magic number
//   size in dimension 0
//   size in dimension 1
//   size in dimension 2
//   .....
//   size in dimension N
//   data
//
// # magic number
// The magic number is an integer (MSB first). The first 2 bytes are always 0.
// The third byte codes the type of the data:
//    0x08: unsigned byte
//    0x09: signed byte
//    0x0B: short (2 bytes)
//    0x0C: int (4 bytes)
//    0x0D: float (4 bytes)
//    0x0E: double (8 bytes)
// The 4-th byte codes the number of dimensions of the vector/matrix: 1 for
// vectors, 2 for matrices....
//
// # dimension size
// The sizes in each dimension are 4-byte integers (MSB first, high endian, like
// in most non-Intel processors).
//
// # data
// The data is stored like in a C array, i.e. the index in the last dimension
// changes the fastest.
const int labelsFileMagic = 0x00000801;
const int imagesFileMagic = 0x00000803;

// internal: Read 4 bytes and convert to big endian integer
error_t readInt32(int fd, int* v_int)
{
        // ssize_t read(int fd, void *buf, size_t count);
        unsigned char buf[4];
        ssize_t       n = read(fd, (void*)buf, 4);
        if (n != 4) {
                return n;
        }
        // manually unroll.
        int v = (int)(buf[0]);
        v <<= 8;
        v += (int)(buf[1]);
        v <<= 8;
        v += (int)(buf[2]);
        v <<= 8;
        v += (int)(buf[3]);
        *v_int = v;
        return OK;
}

int main()
{
        printf("hello mnist:\n    images: %s\n    labels: %s\n", train_images,
               train_labels);

        int fd = open(train_images, O_RDONLY);
        if (fd == -1) {
                printf("failed to open file: %s", train_images);
        } else {
                int v;
                readInt32(fd, &v);
                printf("magic number: %08x\n", v);
                printf("closed file handle: %d\n", fd);
                close(fd);
        }
}
