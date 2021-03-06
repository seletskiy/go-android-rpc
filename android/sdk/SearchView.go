// It is autogenerated bindings for Android sdk class.
//
// See documentation about methods on: https://developer.android.com//reference/android/widget/SearchView.html
package sdk

import "github.com/seletskiy/go-android-rpc/android"

type SearchView struct {
	View
}

func NewSearchView(id string) interface{} {
	obj := SearchView{NewView(id).(View)}

	return obj
}

func (obj SearchView) GetClassName() string {
	return "android.widget.SearchView"
}

func init() {
	android.ViewTypeConstructors["android.widget.SearchView"] = NewSearchView
}

func (obj SearchView) GetImeOptions() (int, error) {
	return return_int(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"getImeOptions",
	))
}

func (obj SearchView) GetInputType() (int, error) {
	return return_int(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"getInputType",
	))
}

func (obj SearchView) GetMaxWidth() (int, error) {
	return return_int(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"getMaxWidth",
	))
}

func (obj SearchView) GetQuery() (string, error) {
	return return_string(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"getQuery",
	))
}

func (obj SearchView) GetQueryHint() (string, error) {
	return return_string(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"getQueryHint",
	))
}

func (obj SearchView) IsIconfiedByDefault() (bool, error) {
	return return_bool(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"isIconfiedByDefault",
	))
}

func (obj SearchView) IsIconified() (bool, error) {
	return return_bool(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"isIconified",
	))
}

func (obj SearchView) IsQueryRefinementEnabled() (bool, error) {
	return return_bool(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"isQueryRefinementEnabled",
	))
}

func (obj SearchView) IsSubmitButtonEnabled() (bool, error) {
	return return_bool(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"isSubmitButtonEnabled",
	))
}

func (obj SearchView) OnActionViewCollapsed() error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"onActionViewCollapsed",
	))
}

func (obj SearchView) OnActionViewExpanded() error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"onActionViewExpanded",
	))
}

func (obj SearchView) OnWindowFocusChanged(hasWindowFocus_ bool) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"onWindowFocusChanged",
		hasWindowFocus_,
	))
}

func (obj SearchView) SetIconified(iconify_ bool) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"setIconified",
		iconify_,
	))
}

func (obj SearchView) SetIconifiedByDefault(iconified_ bool) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"setIconifiedByDefault",
		iconified_,
	))
}

func (obj SearchView) SetImeOptions(imeOptions_ int) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"setImeOptions",
		imeOptions_,
	))
}

func (obj SearchView) SetInputType(inputType_ int) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"setInputType",
		inputType_,
	))
}

func (obj SearchView) SetMaxWidth(maxpixels_ int) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"setMaxWidth",
		maxpixels_,
	))
}

func (obj SearchView) SetQuery(query_ string, submit_ bool) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"setQuery",
		query_, submit_,
	))
}

func (obj SearchView) SetQueryHint(hint_ string) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"setQueryHint",
		hint_,
	))
}

func (obj SearchView) SetQueryRefinementEnabled(enable_ bool) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"setQueryRefinementEnabled",
		enable_,
	))
}

func (obj SearchView) SetSubmitButtonEnabled(enabled_ bool) error {
	return return_error(android.CallViewMethod(
		obj.GetId(),
		"android.widget.SearchView",
		"setSubmitButtonEnabled",
		enabled_,
	))
}

