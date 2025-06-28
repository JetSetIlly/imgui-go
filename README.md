# Dear ImGui for Go

This library is a Go wrapper for [Dear ImGui](https://github.com/ocornut/imgui)

This specific instance of the library is a continuation of the [library started by Christian Haas](https://github.com/inkyblackness/imgui-go), who ceased development some time ago. Despite the original project not being updated in quite some time, I still find it a good solution for my projects. As such, I have decided to update the project so that it uses more recent versions of the base Dear Imgui project.

The current version of Dear Imgui being used is the docking branch of v1.91.9b

Not all features of v1.91.9b are available but in principle they can be added using the existing methods.

Docking can be enabled but viewports are currently unimplemented.

### Changes from version 4 of imgui-go

If you're moving from a previous version of imgui-go you should be aware that some things have changed, due to changes in Dear Imgui itself.

__Keyboard__

`KeyPress()` and `KeyRelease()` have been removed and replaced with `AddKeyEvent()`.
 
See the SDL and GLFW platform examples in the updated [imgui-go-examples](https://github.com/JetSetIlly/imgui-go-examples) repository for how to use the new function.

__Cursor Positioning__

`SetCursorPos()` and `SetScreenCursorPos()` requires a call to `Dummy()` if the position change causes the parent area to be extended. Dear Imgui will panic if it can't validate the new extent.

__Content Regions__

`ContentRegionMax()` has been removed. `ContentRegionAvail()` is a good substitute in many cases, but remember to take the current cursor position into account.

__Disabling Widgets__

`ItemFlagsDisabled` has been removed. Use `BeginDisabled()` and `EndDisabled()` instead.

__ChildFlags__

`BeginChildV()` now takes flags of the type `ChildFlags` and not `WindowFlags`.
 
Some WindowFlags have no ChildFlags equivalent: `NoMove`, `AlwaysAutoResize`, `NoScrollbar` and `AlwaysVerticalScrollbar`.

All missing functionality can be achieved in other ways. For example, `NoMove` can be replicated by detecting the hover status of the Child:

		imgui.BeginChild()
		...
		imgui.EndChild()
		
		state.hover := imgui.IsHovered()
	
And then next frame, changing the window flags for the containing window:

		var flgs imgui.Flags
		if state.hover {
			flags = imgui.WindowFlagsNoMove
		} else  {
			flags = imgui.WindowFlagsNone
		}

__ListClipper__

`ListClipperAll()` replaces all other functions of the ListClipper type `Begin()`, `Step()` and `End()`. The ListClipper type is now returned by the new function and contains only the final results.

		results := imgui.ListClipperAll(len(entries), func(i int) {
			imgui.Text(entries[i])
		})

This new function was added because I was having problems maintaining state when returning from the three canonical functions. Calling Go as a callback from C solves the problem and makes for a convenient function. The downside is that you can't change the ListClipper state between the Step() and DisplayStart-to-DisplayEnd loop - but having tried doing that kind of thing in the past, being prevented from doing it is probably a feature.
