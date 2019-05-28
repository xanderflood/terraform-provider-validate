package validate

import (
	//"fmt"
	th "github.com/craigmonson/terraform-provider-validate/validate/test_helpers"
	//"github.com/hashicorp/terraform/helper/schema"
	"testing"
)

func TestGetCheckTypeExact(t *testing.T) {
	res := dataSourceValidateStr()
	m := make(map[string]interface{}, 3)

	m["val"] = "foo"
	m["exact"] = "bar"
	res_data := th.GetResourceData(t, m, res)
	check, err := getCheckType(res_data)

	if err != nil {
		t.Errorf("Error occurred: %e", err)
	}
	if check != "exact" {
		t.Errorf("'exact' not returned")
	}
}

func TestGetCheckTypeOneOf(t *testing.T) {
	res := dataSourceValidateStr()
	m := make(map[string]interface{}, 3)

	m["val"] = "foo"
	m["one_of"] = []string{"bar"}
	res_data := th.GetResourceData(t, m, res)
	check, err := getCheckType(res_data)

	if err != nil {
		t.Errorf("Error occurred: %e", err)
	}
	if check != "one_of" {
		t.Errorf("'one_of' not returned")
	}
}

func TestGetCheckTypeRegex(t *testing.T) {
	res := dataSourceValidateStr()
	m := make(map[string]interface{}, 3)

	m["val"] = "foo"
	m["regex"] = "bar"
	res_data := th.GetResourceData(t, m, res)
	check, err := getCheckType(res_data)

	if err != nil {
		t.Errorf("Error occurred: %e", err)
	}
	if check != "regex" {
		t.Errorf("'regex' not returned")
	}
}

func TestGetCheckTypeNone(t *testing.T) {
	res := dataSourceValidateStr()
	m := make(map[string]interface{}, 3)

	m["val"] = "foo"
	res_data := th.GetResourceData(t, m, res)
	check, err := getCheckType(res_data)

	if err == nil {
		t.Errorf("No Error: No exact, one_of, or regex selected, but didn't send error")
	}
	if check != "" {
		t.Errorf("value of check is not empty? %v", check)
	}
}

func TestExactStrPass(t *testing.T) {
	err := checkExactStr("foo", "foo")

	if err != nil {
		t.Errorf("Exact check failed: 'foo' vs 'foo'")
	}
}

func TestExactStrFail(t *testing.T) {
	err := checkExactStr("foo", "bar")

	if err == nil {
		t.Errorf("Exact check did not fail: 'foo' vs 'bar'")
	}
}

func TestCheckOneOfStrPass(t *testing.T) {
	res := dataSourceValidateStr()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["one_of"] = []string{"foo"}
	res_data := th.GetResourceData(t, m, res)

	err := checkOneOfStr("foo", res_data.Get("one_of").([]interface{}))

	if err != nil {
		t.Error("one_of check failed: 'foo' vs '[foo]'")
	}
}

func TestCheckOneOfStrFail(t *testing.T) {
	res := dataSourceValidateStr()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["one_of"] = []string{"bar"}
	res_data := th.GetResourceData(t, m, res)

	err := checkOneOfStr("foo", res_data.Get("one_of").([]interface{}))

	if err == nil {
		t.Error("one_of check did not fail: 'foo' vs '[bar]'")
	}
}

func TestCheckRegexStrPass(t *testing.T) {
	err := checkRegexStr("foo", "f..")

	if err != nil {
		t.Error("regex check failed: 'foo' vs '/foo/'")
	}
}

func TestCheckRegexStrFail(t *testing.T) {
	err := checkRegexStr("foo", "/bar/")

	if err == nil {
		t.Error("regex check did not fail: 'foo' vs '/bar/'")
	}
}

func TestCheckRegexStrBadRegex(t *testing.T) {
	err := checkRegexStr("foo", "/[0-9]++/")

	if err == nil {
		t.Error("regex with bad pattern did not fail")
	}
}
