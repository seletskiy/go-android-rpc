// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/AbsListView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type AbsListView struct {
	View
}

func NewAbsListView(id string) interface{} {
	obj := AbsListView{NewView(id).(View)}

	return obj
}

func (obj AbsListView) GetClassName() string {
	return "android.widget.AbsListView"
}

func init() {
	android.ViewTypeConstructors["android.widget.AbsListView"] = NewAbsListView
}

func (obj AbsListView) BeforeTextChanged(s_ string, start_ int, count_ int, after_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"beforeTextChanged",
		s_, start_, count_, after_,
	)
}

func (obj AbsListView) CanScrollList(direction_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"canScrollList",
		direction_,
	)
}

func (obj AbsListView) ClearChoices() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"clearChoices",
	)
}

func (obj AbsListView) ClearTextFilter() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"clearTextFilter",
	)
}

func (obj AbsListView) DeferNotifyDataSetChanged() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"deferNotifyDataSetChanged",
	)
}

func (obj AbsListView) GetCacheColorHint() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getCacheColorHint",
	)
}

func (obj AbsListView) GetCheckedItemCount() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getCheckedItemCount",
	)
}

func (obj AbsListView) GetCheckedItemPosition() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getCheckedItemPosition",
	)
}

func (obj AbsListView) GetChoiceMode() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getChoiceMode",
	)
}

func (obj AbsListView) GetListPaddingBottom() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getListPaddingBottom",
	)
}

func (obj AbsListView) GetListPaddingLeft() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getListPaddingLeft",
	)
}

func (obj AbsListView) GetListPaddingRight() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getListPaddingRight",
	)
}

func (obj AbsListView) GetListPaddingTop() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getListPaddingTop",
	)
}

func (obj AbsListView) GetSolidColor() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getSolidColor",
	)
}

func (obj AbsListView) GetTextFilter() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getTextFilter",
	)
}

func (obj AbsListView) GetTranscriptMode() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getTranscriptMode",
	)
}

func (obj AbsListView) GetVerticalScrollbarWidth() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"getVerticalScrollbarWidth",
	)
}

func (obj AbsListView) HasTextFilter() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"hasTextFilter",
	)
}

func (obj AbsListView) InvalidateViews() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"invalidateViews",
	)
}

func (obj AbsListView) IsFastScrollAlwaysVisible() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"isFastScrollAlwaysVisible",
	)
}

func (obj AbsListView) IsFastScrollEnabled() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"isFastScrollEnabled",
	)
}

func (obj AbsListView) IsItemChecked(position_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"isItemChecked",
		position_,
	)
}

func (obj AbsListView) IsScrollingCacheEnabled() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"isScrollingCacheEnabled",
	)
}

func (obj AbsListView) IsSmoothScrollbarEnabled() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"isSmoothScrollbarEnabled",
	)
}

func (obj AbsListView) IsStackFromBottom() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"isStackFromBottom",
	)
}

func (obj AbsListView) IsTextFilterEnabled() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"isTextFilterEnabled",
	)
}

func (obj AbsListView) JumpDrawablesToCurrentState() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"jumpDrawablesToCurrentState",
	)
}

func (obj AbsListView) OnCancelPendingInputEvents() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"onCancelPendingInputEvents",
	)
}

func (obj AbsListView) OnFilterComplete(count_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"onFilterComplete",
		count_,
	)
}

func (obj AbsListView) OnGlobalLayout() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"onGlobalLayout",
	)
}

func (obj AbsListView) OnRemoteAdapterConnected() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"onRemoteAdapterConnected",
	)
}

func (obj AbsListView) OnRemoteAdapterDisconnected() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"onRemoteAdapterDisconnected",
	)
}

func (obj AbsListView) OnRtlPropertiesChanged(layoutDirection_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"onRtlPropertiesChanged",
		layoutDirection_,
	)
}

func (obj AbsListView) OnTextChanged(s_ string, start_ int, before_ int, count_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"onTextChanged",
		s_, start_, before_, count_,
	)
}

func (obj AbsListView) OnTouchModeChanged(isInTouchMode_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"onTouchModeChanged",
		isInTouchMode_,
	)
}

func (obj AbsListView) OnWindowFocusChanged(hasWindowFocus_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"onWindowFocusChanged",
		hasWindowFocus_,
	)
}

func (obj AbsListView) PointToPosition(x_ int, y_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"pointToPosition",
		x_, y_,
	)
}

func (obj AbsListView) RequestDisallowInterceptTouchEvent(disallowIntercept_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"requestDisallowInterceptTouchEvent",
		disallowIntercept_,
	)
}

func (obj AbsListView) RequestLayout() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"requestLayout",
	)
}

func (obj AbsListView) ScrollListBy(y_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"scrollListBy",
		y_,
	)
}

func (obj AbsListView) SendAccessibilityEvent(eventType_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"sendAccessibilityEvent",
		eventType_,
	)
}

func (obj AbsListView) SetCacheColorHint(color_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setCacheColorHint",
		color_,
	)
}

func (obj AbsListView) SetChoiceMode(choiceMode_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setChoiceMode",
		choiceMode_,
	)
}

func (obj AbsListView) SetDrawSelectorOnTop(onTop_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setDrawSelectorOnTop",
		onTop_,
	)
}

func (obj AbsListView) SetFastScrollAlwaysVisible(alwaysShow_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setFastScrollAlwaysVisible",
		alwaysShow_,
	)
}

func (obj AbsListView) SetFastScrollEnabled(enabled_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setFastScrollEnabled",
		enabled_,
	)
}

func (obj AbsListView) SetFilterText(filterText_ string) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setFilterText",
		filterText_,
	)
}

func (obj AbsListView) SetFriction(friction_ float64) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setFriction",
		android.Float(friction_),
	)
}

func (obj AbsListView) SetItemChecked(position_ int, value_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setItemChecked",
		position_, value_,
	)
}

func (obj AbsListView) SetOverScrollMode(mode_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setOverScrollMode",
		mode_,
	)
}

func (obj AbsListView) SetScrollBarStyle(style_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setScrollBarStyle",
		style_,
	)
}

func (obj AbsListView) SetScrollingCacheEnabled(enabled_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setScrollingCacheEnabled",
		enabled_,
	)
}

func (obj AbsListView) SetSelectionFromTop(position_ int, y_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setSelectionFromTop",
		position_, y_,
	)
}

func (obj AbsListView) SetSelector(resID_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setSelector",
		resID_,
	)
}

func (obj AbsListView) SetSmoothScrollbarEnabled(enabled_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setSmoothScrollbarEnabled",
		enabled_,
	)
}

func (obj AbsListView) SetStackFromBottom(stackFromBottom_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setStackFromBottom",
		stackFromBottom_,
	)
}

func (obj AbsListView) SetTextFilterEnabled(textFilterEnabled_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setTextFilterEnabled",
		textFilterEnabled_,
	)
}

func (obj AbsListView) SetTranscriptMode(mode_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setTranscriptMode",
		mode_,
	)
}

func (obj AbsListView) SetVelocityScale(scale_ float64) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setVelocityScale",
		android.Float(scale_),
	)
}

func (obj AbsListView) SetVerticalScrollbarPosition(position_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"setVerticalScrollbarPosition",
		position_,
	)
}

func (obj AbsListView) SmoothScrollBy(distance_ int, duration_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"smoothScrollBy",
		distance_, duration_,
	)
}

func (obj AbsListView) SmoothScrollToPosition(position_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"smoothScrollToPosition",
		position_,
	)
}

func (obj AbsListView) SmoothScrollToPosition2ii(position_ int, boundPosition_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"smoothScrollToPosition",
		position_, boundPosition_,
	)
}

func (obj AbsListView) SmoothScrollToPositionFromTop(position_ int, offset_ int, duration_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"smoothScrollToPositionFromTop",
		position_, offset_, duration_,
	)
}

func (obj AbsListView) SmoothScrollToPositionFromTop2ii(position_ int, offset_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.AbsListView",
		"smoothScrollToPositionFromTop",
		position_, offset_,
	)
}

