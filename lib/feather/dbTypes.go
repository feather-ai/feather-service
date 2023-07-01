package feather

type String string

/*
* Implement SQL Scan API for NullableString
 */
func (v *String) Scan(src interface{}) error {
	if src == nil {
		*v = ""
		return nil
	}
	*v = String(src.(string))
	return nil
}

func (v String) String() string {
	return string(v)
}
