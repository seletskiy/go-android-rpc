// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/CalendarView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type CalendarView struct {
	View
}

func NewCalendarView(id string) interface{} {
	obj := CalendarView{NewView(id).(View)}

	return obj
}

func (obj CalendarView) GetClassName() string {
	return "android.widget.CalendarView"
}

func init() {
	android.ViewTypeConstructors["android.widget.CalendarView"] = NewCalendarView
}

func (obj CalendarView) GetDateTextAppearance() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"getDateTextAppearance",
	)
}

func (obj CalendarView) GetFirstDayOfWeek() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"getFirstDayOfWeek",
	)
}

func (obj CalendarView) GetFocusedMonthDateColor() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"getFocusedMonthDateColor",
	)
}

func (obj CalendarView) GetSelectedWeekBackgroundColor() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"getSelectedWeekBackgroundColor",
	)
}

func (obj CalendarView) GetShowWeekNumber() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"getShowWeekNumber",
	)
}

func (obj CalendarView) GetShownWeekCount() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"getShownWeekCount",
	)
}

func (obj CalendarView) GetUnfocusedMonthDateColor() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"getUnfocusedMonthDateColor",
	)
}

func (obj CalendarView) GetWeekDayTextAppearance() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"getWeekDayTextAppearance",
	)
}

func (obj CalendarView) GetWeekNumberColor() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"getWeekNumberColor",
	)
}

func (obj CalendarView) GetWeekSeparatorLineColor() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"getWeekSeparatorLineColor",
	)
}

func (obj CalendarView) IsEnabled() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"isEnabled",
	)
}

func (obj CalendarView) SetDateTextAppearance(resourceId_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setDateTextAppearance",
		resourceId_,
	)
}

func (obj CalendarView) SetEnabled(enabled_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setEnabled",
		enabled_,
	)
}

func (obj CalendarView) SetFirstDayOfWeek(firstDayOfWeek_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setFirstDayOfWeek",
		firstDayOfWeek_,
	)
}

func (obj CalendarView) SetFocusedMonthDateColor(color_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setFocusedMonthDateColor",
		color_,
	)
}

func (obj CalendarView) SetSelectedDateVerticalBar(resourceId_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setSelectedDateVerticalBar",
		resourceId_,
	)
}

func (obj CalendarView) SetSelectedWeekBackgroundColor(color_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setSelectedWeekBackgroundColor",
		color_,
	)
}

func (obj CalendarView) SetShowWeekNumber(showWeekNumber_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setShowWeekNumber",
		showWeekNumber_,
	)
}

func (obj CalendarView) SetShownWeekCount(count_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setShownWeekCount",
		count_,
	)
}

func (obj CalendarView) SetUnfocusedMonthDateColor(color_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setUnfocusedMonthDateColor",
		color_,
	)
}

func (obj CalendarView) SetWeekDayTextAppearance(resourceId_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setWeekDayTextAppearance",
		resourceId_,
	)
}

func (obj CalendarView) SetWeekNumberColor(color_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setWeekNumberColor",
		color_,
	)
}

func (obj CalendarView) SetWeekSeparatorLineColor(color_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CalendarView",
		"setWeekSeparatorLineColor",
		color_,
	)
}

