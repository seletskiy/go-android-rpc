package sdk

import "errors"

// such funny namings because of code-generation, which call method with name
// `return_{{.ReturnType}}`.

func return_error(data map[string]interface{}) error {
	if _, ok := data["error"]; ok {
		return errors.New(data["error"].(string))
	} else {
		return nil
	}
}

func return_int(data map[string]interface{}) (int, error) {
	if err := return_error(data); err != nil {
		return 0, err
	}

	if _, ok := data["result"]; ok {
		return data["result"].(int), nil
	}

	return 0, errors.New(`response missing "result" field`)
}

func return_float64(data map[string]interface{}) (float64, error) {
	if err := return_error(data); err != nil {
		return 0, err
	}

	if _, ok := data["result"]; ok {
		return data["result"].(float64), nil
	}

	return 0, errors.New(`response missing "result" field`)
}

func return_bool(data map[string]interface{}) (bool, error) {
	if err := return_error(data); err != nil {
		return false, err
	}

	if _, ok := data["result"]; ok {
		return data["result"].(bool), nil
	}

	return false, errors.New(`response missing "result" field`)
}

func return_string(data map[string]interface{}) (string, error) {
	if err := return_error(data); err != nil {
		return "", err
	}

	if _, ok := data["result"]; ok {
		return data["result"].(string), nil
	}

	return "", errors.New(`response missing "result" field`)
}
