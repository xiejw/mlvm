#include <fcntl.h>  // open
#include <stdio.h>
#include <stdlib.h>  // malloc
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

#define ASSER_NO_ERR(e)                      \
        if ((e)) {                           \
                printf("unexpected error."); \
                return -1;                   \
        }

int main()
{
        printf("hello mnist:\n    images: %s\n    labels: %s\n", train_images,
               train_labels);

        int fd = open(train_images, O_RDONLY);
        if (fd == -1) {
                printf("failed to open file: %s", train_images);
                return OK;
        }

        int v, n, w, h;
        ASSER_NO_ERR(readInt32(fd, &v));
        ASSER_NO_ERR(readInt32(fd, &n));
        ASSER_NO_ERR(readInt32(fd, &w));
        ASSER_NO_ERR(readInt32(fd, &h));

        printf("magic number: %08x\n", v);
        printf("n %d w %d h %d\n", n, w, h);

        printf("reading two samples.\n");
        n                  = 2;
        size_t         s   = n * w * h;
        unsigned char* buf = malloc(sizeof(unsigned char) * s);
        ssize_t        r_s = read(fd, (void*)buf, s);
        ASSER_NO_ERR(r_s != s);

        int line = 0;
        for (int i = 0; i < w; i++) {
                for (int j = 0; j < h; j++) {
                        unsigned char c = buf[j + line];
                        if (c == 0) {
                                printf(" ");
                        } else {
                                printf("%X", c / 16);
                        }
                }
                printf("\n");
                line += h;
        }

        for (int i = 0; i < w; i++) {
                for (int j = 0; j < h; j++) {
                        unsigned char c = buf[j + line];
                        if (c == 0) {
                                printf(" ");
                        } else {
                                printf("%X", c / 16);
                        }
                }
                printf("\n");
                line += h;
        }

        free(buf);
        close(fd);

        fd = open(train_labels, O_RDONLY);
        if (fd == -1) {
                printf("failed to open file: %s", train_labels);
                return OK;
        }

        ASSER_NO_ERR(readInt32(fd, &v));
        ASSER_NO_ERR(readInt32(fd, &n));
        printf("magic number: %08x\n", v);
        printf("n %d\n", n);
        unsigned char l;
        r_s = read(fd, (void*)&l, 1);
        ASSER_NO_ERR(r_s != 1);
        printf("label %d\n", (int)l);

        r_s = read(fd, (void*)&l, 1);
        ASSER_NO_ERR(r_s != 1);
        printf("label %d\n", (int)l);

        close(fd);
        return OK;
}
