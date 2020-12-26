/* This is a simplified version for sds, heavily based on Salvatore's sds.
 *
 * SDSLib 2.0 -- A C dynamic strings library
 *
 * Copyright (c) 2020, Jianwei Xie
 * Copyright (c) 2006-2015, Salvatore Sanfilippo <antirez at gmail dot com>
 * Copyright (c) 2015, Oran Agra
 * Copyright (c) 2015, Redis Labs, Inc
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *   * Redistributions of source code must retain the above copyright notice,
 *     this list of conditions and the following disclaimer.
 *   * Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in the
 *     documentation and/or other materials provided with the distribution.
 *   * Neither the name of Redis nor the names of its contributors may be used
 *     to endorse or promote products derived from this software without
 *     specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

#ifndef SDS_H_
#define SDS_H_

#include <stdarg.h>  // va_list
#include <stdlib.h>

#include "mlvm.h"

typedef char *sds;

extern const char *SDS_NOINIT;

typedef struct {
  size_t len;
  size_t alloc;
} sdshdr;

#define SDS_HDR(s) ((sdshdr *)((s) - (sizeof(sdshdr))))

static inline size_t sdsLen(const sds s) { return SDS_HDR(s)->len; }
static inline size_t sdsCap(const sds s) { return SDS_HDR(s)->alloc; }
static inline size_t sdsAvail(const sds s) {
  sdshdr *p = SDS_HDR(s);
  return p->alloc - p->len;
}
static inline void sdsSetLen(const sds s, size_t newlen) {
  SDS_HDR(s)->len = newlen;
}
static inline void sdsIncLen(const sds s, size_t inc) {
  SDS_HDR(s)->len += inc;
}
static inline void sdsSetCap(const sds s, size_t newcap) {
  SDS_HDR(s)->alloc = newcap;
}
static inline void sdsClear(sds s) {
  sdsSetLen(s, 0);
  s[0] = '\0';
}

sds  sdsNewLen(const void *init, size_t initlen);
sds  sdsNew(const char *init);
sds  sdsEmpty(void);
sds  sdsDup(const sds s);
void sdsFree(sds s);

void sdsCatLen(_mut_ sds *s, const void *t, size_t len);
void sdsCat(_mut_ sds *s, const char *t);
void sdsCatSds(_mut_ sds *s, const sds t);
void sdsCatVprintf(_mut_ sds *s, const char *fmt, va_list ap);
void sdsCatPrintf(_mut_ sds *s, const char *fmt, ...);

// Low level functions exposed to the user API.
void sdsMakeRoomFor(_mut_ sds *s, size_t addlen);

// unimplemented.

sds sdsCpylen(sds s, const char *t, size_t len);
sds sdsCpy(sds s, const char *t);

sds sdscatfmt(sds s, char const *fmt, ...);
int sdscmp(const sds s1, const sds s2);

#endif
