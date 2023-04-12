package config

func ValueOf(name string) string {

	item, err := Fetch(&FetchParam{Name: name})

	if err != nil {
		return ""
	}

	return item.Value

}

func ValuesOf(module string) map[string]string {

	values := map[string]string{}

	items, err := FetchAll(&FetchAllParam{Module: module})

	if err != nil {
		return values
	}

	for _, item := range items {
		values[item.Name] = item.Value
	}
	return values

}
