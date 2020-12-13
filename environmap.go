package environmap

import (
	"os"
)

// EnvironMap represents environment variables in map structure.
type EnvironMap map[string]string

// ParseEnviron convert given environment variable string slice into EnvironMap.
func ParseEnviron(environs []string) (m EnvironMap) {
	envMap := make(map[string]string)
	for _, a := range environs {
		l := len(a)
		for i := 0; i < l; i++ {
			if a[i] == '=' {
				k := a[:i]
				v := a[i+1:]
				envMap[k] = v
				break
			}
		}
	}
	return EnvironMap(envMap)
}

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

func defaultApplyRuntimeEnvironCheck(envKey, envValue string) (shouldApply bool) {
	return (envValue == "")
}

// ApplyRuntimeEnviron replace value from environment variables when given fnShouldApply returns true.
func (m EnvironMap) ApplyRuntimeEnviron(fnShouldApply func(envKey, envValue string) (shouldApply bool)) {
	runtimeEnv := (map[string]string)(ParseEnviron(os.Environ()))
	if fnShouldApply == nil {
		fnShouldApply = defaultApplyRuntimeEnvironCheck
	}
	for rtK, rtV := range runtimeEnv {
		localV, ok := m[rtK]
		if !ok {
			continue
		}
		if fnShouldApply(rtK, localV) {
			m[rtK] = rtV
		}
	}
	return
}
