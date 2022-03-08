#include "tjpgd.h"
#include "tjpgd-go.h"

unsigned char decodeWork[5120];
const int szWork = 5120;

typedef struct {
    unsigned int dataRemain;  // the number of data remaining in the imageBuffer.
    uint8_t *imageBuffer;   // pointer to JPEG image buffer.
} IODEV;

unsigned int in_func (
        JDEC* jdec,       /* Pointer to the decompression object */
        uint8_t* buff,    /* Pointer to buffer to store the read data */
        unsigned int ndata     /* Number of bytes to read */
        );

int out_func (
        JDEC* jdec,    /* Pointer to the decompression object */
        void* bitmap,  /* RGB bitmap to be output */
        JRECT* rect    /* Rectangular region to output */
        );

size_t inFunc(JDEC *jd, uint8_t *buff, size_t nbyte) {
    IODEV *dev = (IODEV *)jd->device;

    if (nbyte > dev->dataRemain) {
        nbyte = dev->dataRemain;
    }
    if (buff) {
        memcpy(buff, dev->imageBuffer, nbyte);
    }
    dev->dataRemain -= nbyte;
    dev->imageBuffer += nbyte;

    return nbyte;
}

int outFunc(JDEC *jd, void *bitmap, JRECT *rect) {
    uint8_t *src;

    src = (uint8_t *)bitmap;
    unsigned short sz = (rect->right - rect->left + 1) * (rect->bottom - rect->top + 1) * 2;

    callbackFromTjpgd(rect->left, rect->top, rect->right, rect->bottom, bitmap, sz);
    return 1;
}

int decodeFromBytes(unsigned char *b, int length, int scale) {
    JDEC jdec;
    IODEV devid;

    devid.dataRemain = length;
    devid.imageBuffer = &b[0];

    JRESULT ret;
    ret = jd_prepare(&jdec, inFunc, (void *)decodeWork, szWork, &devid);
    if (ret == JDR_OK) {
        //printf("Image size is %u sz %u. %u bytes of work ares is used.\n", jdec.width, jdec.height, szWork - jdec.sz_pool);

        ret = jd_decomp(&jdec, outFunc, scale);
        if (ret == JDR_OK) {
            //printf("\rDecompression succeeded.\n");
        }
    }
    return (int)ret;
}
