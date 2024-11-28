#ifndef STRUCTS_H
#define STRUCTS_H

#include <stdint.h>

//typedef struct {
//  BlockID block_id;
//  Block *Block;
//} BlockResult;

typedef int64_t GoRef;

enum ErrorCode {
    EVENTS_BYTES_SYNC_ERROR,
    EVENTS_BYTES_ASYNC_ERROR,
};

#endif