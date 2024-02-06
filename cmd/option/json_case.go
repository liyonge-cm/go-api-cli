package option

import "errors"

type OptJsonCase struct {
	jsonCase string
}

func (c *OptJsonCase) OPTION() string {
	return "j"
}

func (c *OptJsonCase) Help() string {
	return "json case, option: camel or snake, default: snake"
}

func (c *OptJsonCase) SetOptions(option string) error {
	if option == "" {
		return nil
	}
	if option != "camel" && option != "snake" {
		return errors.New("json case must be camel or snake")
	}

	c.jsonCase = option
	return nil
}

func (c *OptJsonCase) Get() string {
	return c.jsonCase
}
