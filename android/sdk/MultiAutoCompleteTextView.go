// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/MultiAutoCompleteTextView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type MultiAutoCompleteTextView struct {
	View
}

func NewMultiAutoCompleteTextView(id string) interface{} {
	obj := MultiAutoCompleteTextView{NewView(id).(View)}

	return obj
}

func (obj MultiAutoCompleteTextView) GetClassName() string {
	return "android.widget.MultiAutoCompleteTextView"
}

func init() {
	android.ViewTypeConstructors["android.widget.MultiAutoCompleteTextView"] = NewMultiAutoCompleteTextView
}

func (obj MultiAutoCompleteTextView) EnoughToFilter() (bool, error) {
	return return_bool(android.CallViewMethod(
		obj.GetId(),
		"android.widget.MultiAutoCompleteTextView",
		"enoughToFilter",
	))
}

func (obj MultiAutoCompleteTextView) PerformValidation() error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.MultiAutoCompleteTextView",
		"performValidation",
	))
}

