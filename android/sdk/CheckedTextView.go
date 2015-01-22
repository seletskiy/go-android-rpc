// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/CheckedTextView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type CheckedTextView struct {
	View
}

func NewCheckedTextView(id string) interface{} {
	obj := CheckedTextView{NewView(id).(View)}

	return obj
}

func (obj CheckedTextView) GetClassName() string {
	return "android.widget.CheckedTextView"
}

func init() {
	android.ViewTypeConstructors["android.widget.CheckedTextView"] = NewCheckedTextView
}

func (obj CheckedTextView) IsChecked() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CheckedTextView",
		"isChecked",
	)
}

func (obj CheckedTextView) JumpDrawablesToCurrentState() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CheckedTextView",
		"jumpDrawablesToCurrentState",
	)
}

func (obj CheckedTextView) OnRtlPropertiesChanged(layoutDirection_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CheckedTextView",
		"onRtlPropertiesChanged",
		layoutDirection_,
	)
}

func (obj CheckedTextView) SetCheckMarkDrawable(resid_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CheckedTextView",
		"setCheckMarkDrawable",
		resid_,
	)
}

func (obj CheckedTextView) SetChecked(checked_ bool) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CheckedTextView",
		"setChecked",
		checked_,
	)
}

func (obj CheckedTextView) SetVisibility(visibility_ int) {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CheckedTextView",
		"setVisibility",
		visibility_,
	)
}

func (obj CheckedTextView) Toggle() {
	android.CallViewMethod(
		obj.GetId(),
		"android.widget.CheckedTextView",
		"toggle",
	)
}

