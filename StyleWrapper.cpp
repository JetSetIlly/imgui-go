#include "imguiWrappedHeader.h"
#include "StyleWrapper.h"
#include "WrapperConverter.h"

void iggStyleSetColor(IggGuiStyle handle, int colorID, IggVec4 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   if ((colorID >= 0) && (colorID < ImGuiCol_COUNT))
   {
      importValue(style->Colors[colorID], *value);
   }
}

void iggStyleScaleAllSizes(IggGuiStyle handle, float scale)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->ScaleAllSizes(scale);
}
