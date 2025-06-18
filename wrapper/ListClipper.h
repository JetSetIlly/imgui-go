#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct tagIggListClipperResults
{
   int DisplayStart;
   int DisplayEnd;
   float ItemsHeight;
} IggListClipperResults;

extern void iggListClipperAll(IggListClipperResults *results, int items_count, float items_height, uintptr_t draw);

#ifdef __cplusplus
}
#endif
