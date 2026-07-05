#include "ConfiguredImGui.h"

#include "Dock.h"

void iggDockSpaceOverViewport()
{
   ImGui::DockSpaceOverViewport(0L, NULL, ImGuiDockNodeFlags_PassthruCentralNode, NULL);
}