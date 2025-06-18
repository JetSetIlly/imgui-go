#pragma once

#include <stdint.h>

// As recommended in imgui/imconfig.h, the IMGUI_USER_CONFIG must be specified
// whenever imgui is used. As such, the wrapper code should never include imgui.h directly.
#define IMGUI_USER_CONFIG "wrapper/ConfigOverride.h"

#define IMGUI_DEFINE_MATH_OPERATORS 1
#define ImTextureID uintptr_t

// IM_OFFSET_OF was obsoleted in dearimgui 1.90. we redefine it here
#define IM_OFFSETOF(_TYPE,_MEMBER)  offsetof(_TYPE, _MEMBER)

#include "imgui.h"
