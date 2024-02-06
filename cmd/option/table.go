package option

type OptTable struct {
	table string
}

func (c *OptTable) OPTION() string {
	return "t"
}

func (c *OptTable) Help() string {
	return "generating api with table, arg: table name"
}

func (c *OptTable) SetOptions(option string) error {
	c.table = option
	return nil
}

func (c *OptTable) Get() string {
	return c.table
}
