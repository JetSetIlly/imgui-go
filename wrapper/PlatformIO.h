#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggPlatformIO iggGetCurrentPlatformIO(void);

extern void iggPlatformIoRegisterClipboardFunctions(IggPlatformIO handle);
extern void iggPlatformIoClearClipboardFunctions(IggPlatformIO handle);

#ifdef __cplusplus
}
#endif
