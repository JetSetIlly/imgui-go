#include "ConfiguredImGui.h"

#include "PlatformIO.h"
#include "WrapperConverter.h"

#include <string>

IggPlatformIO iggGetCurrentPlatformIO()
{
   return reinterpret_cast<IggPlatformIO>(&ImGui::GetPlatformIO());
}

extern "C" void iggPlatformIoSetClipboardText(IggPlatformIO handle, char *text);
extern "C" char *iggPlatformIoGetClipboardText(IggPlatformIO handle);

static void iggPlatformIoSetClipboardTextWrapper(ImGuiContext *ctx, char const *text)
{
   iggPlatformIoSetClipboardText(ctx, const_cast<char *>(text));
}

static char const *iggPlatformIoGetClipboardTextWrapper(ImGuiContext *ctx)
{
   return iggPlatformIoGetClipboardText(ctx);
}

void iggPlatformIoRegisterClipboardFunctions(IggPlatformIO handle)
{
   ImGuiPlatformIO &io = *reinterpret_cast<ImGuiPlatformIO *>(handle);
   io.Platform_ClipboardUserData = handle;
   io.Platform_GetClipboardTextFn = iggPlatformIoGetClipboardTextWrapper;
   io.Platform_SetClipboardTextFn = iggPlatformIoSetClipboardTextWrapper;
}

void iggPlatformIoClearClipboardFunctions(IggPlatformIO handle)
{
   ImGuiPlatformIO &io = *reinterpret_cast<ImGuiPlatformIO *>(handle);
   io.Platform_GetClipboardTextFn = nullptr;
   io.Platform_SetClipboardTextFn = nullptr;
   io.Platform_ClipboardUserData = nullptr;
}
