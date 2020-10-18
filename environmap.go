package environmap

// EnvironMap represents environment variables in map structure.
type EnvironMap map[string]string

// Merge imports environment variables from another map.
// Values from given map will take precedence over value in this map.
func (m EnvironMap) Merge(envVars map[string]string) {
	for k, v := range envVars {
		m[k] = v
	}
}

// ToStrings transforms key-value pairs into string slice. Each pair (k, v) will
// translate into string in `key=value` form.
func (m EnvironMap) ToStrings() (envList []string) {
	l := len(m)
	envList = make([]string, 0, l)
	for k, v := range m {
		aux := k + "=" + v
		envList = append(envList, aux)
	}
	return envList
}
