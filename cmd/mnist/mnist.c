// include for one-off
#include <fcntl.h>   // open
#include <stdlib.h>  // malloc
#include <unistd.h>  // close

#include "base/error.h"

// http://yann.lecun.com/exdb/mnist/
// gzip -d -k file_name.gz
#define MNIST_PATH "../mlvm_examples/files/"

static const char* const train_images = MNIST_PATH "train-images-idx3-ubyte";
static const char* const train_labels = MNIST_PATH "train-labels-idx1-ubyte";
static const char* const test_images  = MNIST_PATH "t10k-images-idx3-ubyte";
static const char* const test_labels  = MNIST_PATH "t10k-labels-idx1-ubyte";

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
static const int labelsFileMagic = 0x00000801;
static const int imagesFileMagic = 0x00000803;

// internal: Read 4 bytes and convert to big endian integer
static error_t
readInt32(int fd, int* v_int)
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

#define NO_ERR(e)                                                            \
        if ((e)) {                                                           \
                errEmitNote("unexpected error during parsing mnist files."); \
        }

error_t
readMnistTrainingImages(unsigned char** data)
{
        error_t        err = OK;
        unsigned char* buf = NULL;
        int            v, n, w, h;
        int            fd = open(train_images, O_RDONLY);

        if (fd == -1) {
                return errNew("failed to open file: %s", train_images);
        }

        NO_ERR(readInt32(fd, &v));
        if (v != imagesFileMagic) {
                err = errNew("image file magic number not match.");
                goto clean;
        }

        NO_ERR(readInt32(fd, &n));
        NO_ERR(readInt32(fd, &w));
        NO_ERR(readInt32(fd, &h));

        size_t s = n * w * h;

        buf         = malloc(sizeof(unsigned char) * s);
        ssize_t r_s = read(fd, (void*)buf, s);
        if (r_s != s) {
                err = errNew("file is incomplete.");
                goto clean;
        }
        *data = buf;

clean:
        close(fd);
        if (err != OK && buf != NULL) {
                free(buf);
        }
        return err;
}

error_t
readMnistTrainingLabels(unsigned char** data)
{
        error_t        err = OK;
        unsigned char* buf = NULL;
        int            v, n;
        int            fd = open(train_labels, O_RDONLY);

        if (fd == -1) {
                return errNew("failed to open file: %s", train_labels);
        }

        NO_ERR(readInt32(fd, &v));
        if (v != labelsFileMagic) {
                err = errNew("label file magic number not match.");
                goto clean;
        }

        NO_ERR(readInt32(fd, &n));

        buf       = malloc(sizeof(unsigned char) * n);
        ssize_t s = read(fd, (void*)buf, n);
        if (n != s) {
                err = errNew("file is incomplete.");
                goto clean;
        }
        *data = buf;

clean:
        close(fd);
        if (err != OK && buf != NULL) {
                free(buf);
        }
        return err;
}

// prints the mnist image at buf.
void
printMnistImage(unsigned char* buf)
{
        int line = 0;
        for (int i = 0; i < 28; i++) {
                for (int j = 0; j < 28; j++) {
                        unsigned char c = buf[j + line];
                        if (c == 0) {
                                printf(" ");
                        } else {
                                printf("%X", c / 16);
                        }
                }
                printf("\n");
                line += 28;
        }
}

#undef NO_ERR
