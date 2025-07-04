package imgui

// #include "wrapper/Tables.h"
import "C"

// TableFlags for BeginTableV().
type TableFlags int

// Tables
// [BETA API] API may evolve slightly! If you use this, please update to the next version when it comes out!
// - Full-featured replacement for old Columns API.
// - See Demo->Tables for demo code.
// - See top of imgui_tables.cpp for general commentary.
// - See TableFlags_ and TableColumnFlags_ enums for a description of available flags.
// The typical call flow is:
// - 1. Call BeginTable().
// - 2. Optionally call TableSetupColumn() to submit column name/flags/defaults.
// - 3. Optionally call TableSetupScrollFreeze() to request scroll freezing of columns/rows.
// - 4. Optionally call TableHeadersRow() to submit a header row. Names are pulled from TableSetupColumn() data.
// - 5. Populate contents:
//   - In most situations you can use TableNextRow() + TableSetColumnIndex(N) to start appending into a column.
//   - If you are using tables as a sort of grid, where every columns is holding the same type of contents,
//     you may prefer using TableNextColumn() instead of TableNextRow() + TableSetColumnIndex().
//     TableNextColumn() will automatically wrap-around into the next row if needed.
//   - IMPORTANT: Comparatively to the old Columns() API, we need to call TableNextColumn() for the first column!
//   - Summary of possible call flow:
//     ----------------------------------------------------------------------------------------------------------
//     TableNextRow() -> TableSetColumnIndex(0) -> Text("Hello 0") -> TableSetColumnIndex(1) -> Text("Hello 1")  // OK
//     TableNextRow() -> TableNextColumn()      -> Text("Hello 0") -> TableNextColumn()      -> Text("Hello 1")  // OK
//     TableNextColumn()      -> Text("Hello 0") -> TableNextColumn()      -> Text("Hello 1")  // OK: TableNextColumn() automatically gets to next row!
//     TableNextRow()                           -> Text("Hello 0")                                               // Not OK! Missing TableSetColumnIndex() or TableNextColumn()! Text will not appear!
//     ----------------------------------------------------------------------------------------------------------
//
// - 5. Call EndTable()
//
// Flags for BeginTable()
// [BETA API] API may evolve slightly! If you use this, please update to the next version when it comes out!
//   - Important! Sizing policies have complex and subtle side effects, more so than you would expect.
//     Read comments/demos carefully + experiment with live demos to get acquainted with them.
//   - The DEFAULT sizing policies are:
//   - Default to TableFlagsSizingFixedFit    if ScrollX is on, or if host window has WindowFlagsAlwaysAutoResize.
//   - Default to TableFlagsSizingStretchSame if ScrollX is off.
//   - When ScrollX is off:
//   - Table defaults to TableFlagsSizingStretchSame -> all Columns defaults to TableColumnFlagsWidthStretch with same weight.
//   - Columns sizing policy allowed: Stretch (default), Fixed/Auto.
//   - Fixed Columns will generally obtain their requested width (unless the table cannot fit them all).
//   - Stretch Columns will share the remaining width.
//   - Mixed Fixed/Stretch columns is possible but has various side-effects on resizing behaviors.
//     The typical use of mixing sizing policies is: any number of LEADING Fixed columns, followed by one or two TRAILING Stretch columns.
//     (this is because the visible order of columns have subtle but necessary effects on how they react to manual resizing).
//   - When ScrollX is on:
//   - Table defaults to TableFlagsSizingFixedFit -> all Columns defaults to TableColumnFlagsWidthFixed
//   - Columns sizing policy allowed: Fixed/Auto mostly.
//   - Fixed Columns can be enlarged as needed. Table will show an horizontal scrollbar if needed.
//   - When using auto-resizing (non-resizable) fixed columns, querying the content width to use item right-alignment e.g. SetNextItemWidth(-FLT_MIN) doesn't make sense, would create a feedback loop.
//   - Using Stretch columns OFTEN DOES NOT MAKE SENSE if ScrollX is on, UNLESS you have specified a value for 'inner_width' in BeginTable().
//     If you specify a value for 'inner_width' then effectively the scrolling space is known and Stretch or mixed Fixed/Stretch columns become meaningful again.
//   - Read on documentation at the top of imgui_tables.cpp for details.
const (
	// Features
	TableFlagsNone              TableFlags = 0
	TableFlagsResizable         TableFlags = 1 << 0 // Enable resizing columns.
	TableFlagsReorderable       TableFlags = 1 << 1 // Enable reordering columns in header row (need calling TableSetupColumn() + TableHeadersRow() to display headers)
	TableFlagsHideable          TableFlags = 1 << 2 // Enable hiding/disabling columns in context menu.
	TableFlagsSortable          TableFlags = 1 << 3 // Enable sorting. Call TableGetSortSpecs() to obtain sort specs. Also see TableFlagsSortMulti and TableFlagsSortTristate.
	TableFlagsNoSavedSettings   TableFlags = 1 << 4 // Disable persisting columns order, width and sort settings in the .ini file.
	TableFlagsContextMenuInBody TableFlags = 1 << 5 // Right-click on columns body/contents will display table context menu. By default it is available in TableHeadersRow().
	// Decorations
	TableFlagsRowBg                      TableFlags = 1 << 6                                            // Set each RowBg color with ImGuiCol_TableRowBg or ImGuiCol_TableRowBgAlt (equivalent of calling TableSetBgColor with ImGuiTableBgFlags_RowBg0 on each row manually)
	TableFlagsBordersInnerH              TableFlags = 1 << 7                                            // Draw horizontal borders between rows.
	TableFlagsBordersOuterH              TableFlags = 1 << 8                                            // Draw horizontal borders at the top and bottom.
	TableFlagsBordersInnerV              TableFlags = 1 << 9                                            // Draw vertical borders between columns.
	TableFlagsBordersOuterV              TableFlags = 1 << 10                                           // Draw vertical borders on the left and right sides.
	TableFlagsBordersH                   TableFlags = TableFlagsBordersInnerH | TableFlagsBordersOuterH // Draw horizontal borders.
	TableFlagsBordersV                   TableFlags = TableFlagsBordersInnerV | TableFlagsBordersOuterV // Draw vertical borders.
	TableFlagsBordersInner               TableFlags = TableFlagsBordersInnerV | TableFlagsBordersInnerH // Draw inner borders.
	TableFlagsBordersOuter               TableFlags = TableFlagsBordersOuterV | TableFlagsBordersOuterH // Draw outer borders.
	TableFlagsBorders                    TableFlags = TableFlagsBordersInner | TableFlagsBordersOuter   // Draw all borders.
	TableFlagsNoBordersInBody            TableFlags = 1 << 11                                           // [ALPHA] Disable vertical borders in columns Body (borders will always appear in Headers). -> May move to style
	TableFlagsNoBordersInBodyUntilResize TableFlags = 1 << 12                                           // [ALPHA] Disable vertical borders in columns Body until hovered for resize (borders will always appear in Headers). -> May move to style
	// Sizing Policy (read above for defaults)
	TableFlagsSizingFixedFit    TableFlags = 1 << 13 // Columns default to _WidthFixed or _WidthAuto (if resizable or not resizable), matching contents width.
	TableFlagsSizingFixedSame   TableFlags = 2 << 13 // Columns default to _WidthFixed or _WidthAuto (if resizable or not resizable), matching the maximum contents width of all columns. Implicitly enable TableFlagsNoKeepColumnsVisible.
	TableFlagsSizingStretchProp TableFlags = 3 << 13 // Columns default to _WidthStretch with default weights proportional to each columns contents widths.
	TableFlagsSizingStretchSame TableFlags = 4 << 13 // Columns default to _WidthStretch with default weights all equal, unless overridden by TableSetupColumn().
	// Sizing Extra Options
	TableFlagsNoHostExtendX        TableFlags = 1 << 16 // Make outer width auto-fit to columns, overriding outer_size.x value. Only available when ScrollX/ScrollY are disabled and Stretch columns are not used.
	TableFlagsNoHostExtendY        TableFlags = 1 << 17 // Make outer height stop exactly at outer_size.y (prevent auto-extending table past the limit). Only available when ScrollX/ScrollY are disabled. Data below the limit will be clipped and not visible.
	TableFlagsNoKeepColumnsVisible TableFlags = 1 << 18 // Disable keeping column always minimally visible when ScrollX is off and table gets too small. Not recommended if columns are resizable.
	TableFlagsPreciseWidths        TableFlags = 1 << 19 // Disable distributing remainder width to stretched columns (width allocation on a 100-wide table with 3 columns: Without this flag: 33,33,34. With this flag: 33,33,33). With larger number of columns, resizing will appear to be less smooth.
	// Clipping
	TableFlagsNoClip TableFlags = 1 << 20 // Disable clipping rectangle for every individual columns (reduce draw command count, items will be able to overflow into other columns). Generally incompatible with TableSetupScrollFreeze().
	// Padding
	TableFlagsPadOuterX   TableFlags = 1 << 21 // Default if BordersOuterV is on. Enable outermost padding. Generally desirable if you have headers.
	TableFlagsNoPadOuterX TableFlags = 1 << 22 // Default if BordersOuterV is off. Disable outermost padding.
	TableFlagsNoPadInnerX TableFlags = 1 << 23 // Disable inner padding between columns (double inner padding if BordersOuterV is on, single inner padding if BordersOuterV is off).
	// Scrolling
	TableFlagsScrollX TableFlags = 1 << 24 // Enable horizontal scrolling. Require 'outer_size' parameter of BeginTable() to specify the container size. Changes default sizing policy. Because this creates a child window, ScrollY is currently generally recommended when using ScrollX.
	TableFlagsScrollY TableFlags = 1 << 25 // Enable vertical scrolling. Require 'outer_size' parameter of BeginTable() to specify the container size.
	// Sorting
	TableFlagsSortMulti    TableFlags = 1 << 26 // Hold shift when clicking headers to sort on multiple column. TableGetSortSpecs() may return specs where (SpecsCount > 1).
	TableFlagsSortTristate TableFlags = 1 << 27 // Allow no sorting, disable default sorting. TableGetSortSpecs() may return specs where (SpecsCount == 0).
	// Miscellaneous
	TableFlagsHighlightHoveredColumn TableFlags = 1 << 28 // Highlight column headers when hovered (may evolve into a fuller highlight)
)

// TableColumnFlags for TableSetupColumnV().
type TableColumnFlags int

const (
	// Input configuration flags
	TableColumnFlagsNone                 TableColumnFlags = 0
	TableColumnFlagsDisabled             TableColumnFlags = 1 << 0  // Overriding/master disable flag: hide column, won't show in context menu (unlike calling TableSetColumnEnabled() which manipulates the user accessible state)
	TableColumnFlagsDefaultHide          TableColumnFlags = 1 << 1  // Default as a hidden/disabled column.
	TableColumnFlagsDefaultSort          TableColumnFlags = 1 << 2  // Default as a sorting column.
	TableColumnFlagsWidthStretch         TableColumnFlags = 1 << 3  // Column will stretch. Preferable with horizontal scrolling disabled (default if table sizing policy is _SizingStretchSame or _SizingStretchProp).
	TableColumnFlagsWidthFixed           TableColumnFlags = 1 << 4  // Column will not stretch. Preferable with horizontal scrolling enabled (default if table sizing policy is _SizingFixedFit and table is resizable).
	TableColumnFlagsNoResize             TableColumnFlags = 1 << 5  // Disable manual resizing.
	TableColumnFlagsNoReorder            TableColumnFlags = 1 << 6  // Disable manual reordering this column, this will also prevent other columns from crossing over this column.
	TableColumnFlagsNoHide               TableColumnFlags = 1 << 7  // Disable ability to hide/disable this column.
	TableColumnFlagsNoClip               TableColumnFlags = 1 << 8  // Disable clipping for this column (all NoClip columns will render in a same draw command).
	TableColumnFlagsNoSort               TableColumnFlags = 1 << 9  // Disable ability to sort on this field (even if ImGuiTableFlags_Sortable is set on the table).
	TableColumnFlagsNoSortAscending      TableColumnFlags = 1 << 10 // Disable ability to sort in the ascending direction.
	TableColumnFlagsNoSortDescending     TableColumnFlags = 1 << 11 // Disable ability to sort in the descending direction.
	TableColumnFlagsNoHeaderLabel        TableColumnFlags = 1 << 12 // TableHeadersRow() will submit an empty label for this column. Convenient for some small columns. Name will still appear in context menu or in angled headers. You may append into this cell by calling TableSetColumnIndex() right after the TableHeadersRow() call.
	TableColumnFlagsNoHeaderWidth        TableColumnFlags = 1 << 13 // Disable header text width contribution to automatic column width.
	TableColumnFlagsPreferSortAscending  TableColumnFlags = 1 << 14 // Make the initial sort direction Ascending when first sorting on this column (default).
	TableColumnFlagsPreferSortDescending TableColumnFlags = 1 << 15 // Make the initial sort direction Descending when first sorting on this column.
	TableColumnFlagsIndentEnable         TableColumnFlags = 1 << 16 // Use current Indent value when entering cell (default for column 0).
	TableColumnFlagsIndentDisable        TableColumnFlags = 1 << 17 // Ignore current Indent value when entering cell (default for columns > 0). Indentation changes _within_ the cell will still be honored.
	TableColumnFlagsAngledHeader         TableColumnFlags = 1 << 18 // TableHeadersRow() will submit an angled header row for this column. Note this will add an extra row.

	// Output status flags read-only via TableGetColumnFlags()
	TableColumnFlagsIsEnabledTableColumnFlags                  = 1 << 24 // Status: is enabled == not hidden by user/api (referred to as "Hide" in _DefaultHide and _NoHide) flags.
	TableColumnFlagsIsVisibleTableColumnFlags                  = 1 << 25 // Status: is visible == is enabled AND not clipped by scrolling.
	TableColumnFlagsIsSorted                  TableColumnFlags = 1 << 26 // Status: is currently part of the sort specs
	TableColumnFlagsIsHoveredTableColumnFlags                  = 1 << 27 // Status: is hovered by mouse
)

// TableRowFlags for TableNextRowV().
type TableRowFlags int

const (
	// TableRowFlagsNone default = 0.
	TableRowFlagsNone TableRowFlags = 0
	// TableRowFlagsHeaders identify header row (set default background color + width of its contents accounted different for auto column width).
	TableRowFlagsHeaders TableRowFlags = 1 << 0
)

// TableBgTarget for TableSetBgColor
//
// Background colors are rendering in 3 layers:
//   - Layer 0: draw with RowBg0 color if set, otherwise draw with ColumnBg0 if set.
//   - Layer 1: draw with RowBg1 color if set, otherwise draw with ColumnBg1 if set.
//   - Layer 2: draw with CellBg color if set.
//
// The purpose of the two row/columns layers is to let you decide if a
// background color changes should override or blend with the existing color.
// When using TableFlagsRowBg on the table, each row has the RowBg0 color automatically set for odd/even rows.
// If you set the color of RowBg0 target, your color will override the existing RowBg0 color.
// If you set the color of RowBg1 or ColumnBg1 target, your color will blend over the RowBg0 color.
type TableBgTarget int

const (
	// TableBgTargetNone default = 0.
	TableBgTargetNone TableBgTarget = 0
	// TableBgTargetRowBg0 sets row background color 0 (generally used for background, automatically set when TableFlagsRowBg is used).
	TableBgTargetRowBg0 TableBgTarget = 1
	// TableBgTargetRowBg1 sets row background color 1 (generally used for selection marking).
	TableBgTargetRowBg1 TableBgTarget = 2
	// TableBgTargetCellBg sets cell background color (top-most color).
	TableBgTargetCellBg TableBgTarget = 3
)

// SortDirection used in TableColumnSortSpecs, etc.
type SortDirection int

const (
	// SortDirectionNone no sort.
	SortDirectionNone SortDirection = 0
	// SortDirectionAscending sorts Ascending = 0->9, A->Z etc.
	SortDirectionAscending SortDirection = 1
	// SortDirectionDescending sorts Descending = 9->0, Z->A etc.
	SortDirectionDescending SortDirection = 2
)

// BeginTableV creates a table
// (Read carefully because this is subtle but it does make sense!)
// About 'outerSize', its meaning needs to differ slightly depending of if we are using ScrollX/ScrollY flags:
//
//	X:
//	- outerSize.x < 0.0f  ->  right align from window/work-rect maximum x edge.
//	- outerSize.x = 0.0f  ->  auto enlarge, use all available space.
//	- outerSize.x > 0.0f  ->  fixed width
//	Y with ScrollX/ScrollY: using a child window for scrolling:
//	- outerSize.y < 0.0f  ->  bottom align
//	- outerSize.y = 0.0f  ->  bottom align, consistent with BeginChild(). not recommended unless table is last item in parent window.
//	- outerSize.y > 0.0f  ->  fixed child height. recommended when using Scrolling on any axis.
//	Y without scrolling, we output table directly in parent window:
//	- outerSize.y < 0.0f  ->  bottom align (will auto extend, unless NoHostExtendV is set)
//	- outerSize.y = 0.0f  ->  zero minimum height (will auto extend, unless NoHostExtendV is set)
//	- outerSize.y > 0.0f  ->  minimum height (will auto extend, unless NoHostExtendV is set)
//
// About 'innerWidth':
//
//	With ScrollX:
//	- innerWidth  < 0.0f  ->  *illegal* fit in known width (right align from outerSize.x) <-- weird
//	- innerWidth  = 0.0f  ->  fit in outer_width: Fixed size columns will take space they need (if avail, otherwise shrink down), Stretch columns becomes Fixed columns.
//	- innerWidth  > 0.0f  ->  override scrolling width, generally to be larger than outerSize.x. Fixed column take space they need (if avail, otherwise shrink down), Stretch columns share remaining space!
//	Without ScrollX:
//	- innerWidth          ->  *ignored*
//
// Details:
//   - If you want to use Stretch columns with ScrollX, you generally need to specify 'innerWidth' otherwise the concept
//     of "available space" doesn't make sense.
//   - Even if not really useful, we allow 'innerWidth < outerSize.x' for consistency and to facilitate understanding
//     of what the value does.
func BeginTableV(id string, columnsCount int, flags TableFlags, outerSize Vec2, innerWidth float32) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	outerSizeArg, _ := outerSize.wrapped()
	return C.iggBeginTable(idArg, C.int(columnsCount), C.int(flags), outerSizeArg, C.float(innerWidth)) != 0
}

// BeginTable calls BeginTableV(id, columnsCount, 0, imgui.Vec2{}, 0.0).
func BeginTable(id string, columnsCount int) bool {
	return BeginTableV(id, columnsCount, 0, Vec2{}, 0.0)
}

// EndTable closes the scope for the previously opened table.
// only call EndTable() if BeginTable() returns true!
func EndTable() {
	C.iggEndTable()
}

// TableNextRowV appends into the first cell of a new row.
func TableNextRowV(flags TableRowFlags, minRowHeight float32) {
	C.iggTableNextRow(C.int(flags), C.float(minRowHeight))
}

// TableNextRow calls TableNextRowV(0, 0.0).
func TableNextRow() {
	TableNextRowV(0, 0.0)
}

// TableNextColumn appends into the next column (or first column of next row if currently in last column)
// Return true when column is visible.
func TableNextColumn() bool {
	return C.iggTableNextColumn() != 0
}

// TableSetColumnIndex appends into the specified column. Return true when column is visible.
func TableSetColumnIndex(columnN int) bool {
	return C.iggTableSetColumnIndex(C.int(columnN)) != 0
}

// TableSetupColumnV specify label, resizing policy, default width/weight, id, various other flags etc.
//   - Use TableSetupColumn() to specify label, resizing policy, default width/weight, id, various other flags etc.
//   - Use TableHeadersRow() to create a header row and automatically submit a TableHeader() for each column.
//     Headers are required to perform: reordering, sorting, and opening the context menu.
//     The context menu can also be made available in columns body using TableFlags_ContextMenuInBody.
//   - You may manually submit headers using TableNextRow() + TableHeader() calls, but this is only useful in
//     some advanced use cases (e.g. adding custom widgets in header row).
//   - Use TableSetupScrollFreeze() to lock columns/rows so they stay visible when scrolled.
func TableSetupColumnV(label string, flags TableColumnFlags, initWidthOrHeight float32, userID uint) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	C.iggTableSetupColumn(labelArg, C.int(flags), C.float(initWidthOrHeight), C.uint(userID))
}

// TableSetupColumn calls TableSetupColumnV(label, 0, 0.0, 0).
func TableSetupColumn(label string) {
	TableSetupColumnV(label, 0, 0.0, 0)
}

// TableSetupScrollFreeze locks columns/rows so they stay visible when scrolled.
func TableSetupScrollFreeze(cols int, rows int) {
	C.iggTableSetupScrollFreeze(C.int(cols), C.int(rows))
}

// TableHeadersRow submits all headers cells based on data provided to TableSetupColumn() + submit context menu.
func TableHeadersRow() {
	C.iggTableHeadersRow()
}

// TableHeader submits one header cell manually (rarely used).
func TableHeader(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	C.iggTableHeader(labelArg)
}

// TableGetColumnCount returns number of columns (value passed to BeginTable).
func TableGetColumnCount() int {
	return int(C.iggTableGetColumnCount())
}

// TableGetColumnIndex return current column index.
func TableGetColumnIndex() int {
	return int(C.iggTableGetColumnIndex())
}

// TableGetRowIndex return current row index.
func TableGetRowIndex() int {
	return int(C.iggTableGetRowIndex())
}

// TableGetColumnNameV returns "" if column didn't have a name declared by TableSetupColumn(). Pass -1 to use current column.
func TableGetColumnNameV(columnN int) string {
	return C.GoString(C.iggTableGetColumnName(C.int(columnN)))
}

// TableGetColumnName calls TableGetColumnNameV(-1).
func TableGetColumnName() string {
	return TableGetColumnNameV(-1)
}

// TableGetColumnFlagsV return column flags so you can query their Enabled/Visible/Sorted/Hovered status flags. Pass -1 to use current column.
func TableGetColumnFlagsV(columnN int) TableColumnFlags {
	return TableColumnFlags(C.iggTableGetColumnFlags(C.int(columnN)))
}

// TableGetColumnFlags calls TableGetColumnFlagsV(-1).
func TableGetColumnFlags() TableColumnFlags {
	return TableGetColumnFlagsV(-1)
}

// TableSetBgColorV changes the color of a cell, row, or column. See TableBgTarget flags for details.
func TableSetBgColorV(target TableBgTarget, color Vec4, columnN int) {
	colorArg, _ := color.wrapped()
	C.iggTableSetBgColor(C.int(target), colorArg, C.int(columnN))
}

// TableSetBgColor calls TableSetBgColorV(target, color, -1).
func TableSetBgColor(target TableBgTarget, color Vec4) {
	TableSetBgColorV(target, color, -1)
}

// TableSortSpecs is a sort specs
// Sorting specifications for a table (often handling sort specs for a single column, occasionally more)
// Obtained by calling TableGetSortSpecs()
// When SpecsDirty() == true you can sort your data. It will be true with sorting specs have changed since last call, or the first time.
// Make sure to call ClearSpecsDirty() or SetSpecsDirty(false) after sorting, else you may wastefully sort your data every frame!
type TableSortSpecs uintptr

// TableGetSortSpecs gets latest sort specs for the table (0 if not sorting).
//   - Call TableGetSortSpecs() to retrieve latest sort specs for the table. 0 when not sorting.
//   - When 'SpecsDirty() == true' you should sort your data. It will be true when sorting specs have changed
//     since last call, or the first time. Make sure to call ClearSpecsDirty() or SetSpecsDirty(false) after sorting, else you may
//     wastefully sort your data every frame!
//   - Lifetime: don't hold on this pointer over multiple frames or past any subsequent call to BeginTable().
func TableGetSortSpecs() TableSortSpecs {
	return TableSortSpecs(C.iggTableGetSortSpecs())
}

func (specs TableSortSpecs) handle() C.IggTableSortSpecs {
	return C.IggTableSortSpecs(specs)
}

// TableColumnSortSpecs is a sorting specification for one column of a table (sizeof == 12 bytes).
type TableColumnSortSpecs struct {
	ColumnUserID  uint          // User id of the column (if specified by a TableSetupColumn() call)
	ColumnIndex   int16         // Index of the column
	SortOrder     int16         // Index within parent TableSortSpecs (always stored in order starting from 0, tables sorted on a single criteria will always have a 0 here)
	SortDirection SortDirection // SortDirectionAscending or SortDirectionDescending (you can use this or SortSign, whichever is more convenient for your sort function)
}

// Specs returns columns sort spec array.
func (specs TableSortSpecs) Specs() []TableColumnSortSpecs {
	count := specs.SpecsCount()
	columnSpecs := make([]TableColumnSortSpecs, count)
	for i := 0; i < count; i++ {
		out := &C.IggTableColumnSortSpecs{}
		C.iggTableSortSpecsGetSpec(specs.handle(), C.int(i), out)

		columnSpecs[i] = TableColumnSortSpecs{
			ColumnUserID:  uint(out.ColumnUserID),
			ColumnIndex:   int16(out.ColumnIndex),
			SortOrder:     int16(out.SortOrder),
			SortDirection: SortDirection(out.SortDirection),
		}
	}
	return columnSpecs
}

// SpecsCount returns sort spec count. Most often 1 unless e.g. TableFlagsMultiSortable is enabled.
func (specs TableSortSpecs) SpecsCount() int {
	if specs == 0 {
		return 0
	}
	return int(C.iggTableSortSpecsGetSpecsCount(specs.handle()))
}

// SpecsDirty returns if specs have changed since last time! Use this to sort again.
func (specs TableSortSpecs) SpecsDirty() bool {
	if specs == 0 {
		return false
	}
	return C.iggTableSortSpecsGetSpecsDirty(specs.handle()) != 0
}

// SetSpecsDirty sets SpecsDirty value to a given value, usually used to clear the flag (set to false).
func (specs TableSortSpecs) SetSpecsDirty(value bool) {
	if specs == 0 {
		return
	}
	C.iggTableSortSpecsSetSpecsDirty(specs.handle(), castBool(value))
}

// ClearSpecsDirty calls specs.SetSpecsDirty(false).
func (specs TableSortSpecs) ClearSpecsDirty() {
	specs.SetSpecsDirty(false)
}
