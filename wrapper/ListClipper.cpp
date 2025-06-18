#include "ConfiguredImGui.h"

#include "ListClipper.h"
#include "WrapperConverter.h"
#include "_cgo_export.h"

#include <cstdint>
#include <iostream>

void iggListClipperAll(IggListClipperResults *results, int items_count, float items_height, uintptr_t draw)
{
   ImGuiListClipper imguiClipper;
   imguiClipper.Begin(items_count, items_height);
   while (imguiClipper.Step()) {
	   for (int i = imguiClipper.DisplayStart; i < imguiClipper.DisplayEnd; ++i) {
		   listClipperDraw(draw, i);
	   }
   }
   results->DisplayStart = imguiClipper.DisplayStart;
   results->DisplayEnd = imguiClipper.DisplayEnd;
   results->ItemsHeight = imguiClipper.ItemsHeight;
}
