#ifndef ERROR_H
#define ERROR_H

// Design:
//
// - All error codes are negative numbers. 0 is success.
// - Limited error codes are pre-defined here. If possible matching to these
//   cannonical codes or just -1 (EUNSPECIFIED)
//
// - A special bit is used to indicate whether a thread local error message is
//   set.  This error message is stored in a global space. So, lifetime is
//   static.
//
//   With the bit stored in error code, it is easy to, based on this signal,
//   reset or concat the new error message.

typedef int errort;

#define OK 0

#define EUNSPECIFIED -1
#define ENOT_FOUND   -2

#endif
