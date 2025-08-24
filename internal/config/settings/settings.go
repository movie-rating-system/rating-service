package settings

import (
	"errors"
	"os"
	"reflect"
)

type Settings struct {
	HTTP HTTP `env:"HTTP"`
}

var ErrSettingsCannotBeFetched = errors.New("cannot fetch settings")

func FetchSettings() (*Settings, error) {
	appSettings := &Settings{}
	st := reflect.TypeOf(appSettings).Elem()
	if err := fetchConcreteSettings(st, ""); err != nil {
		return nil, err
	}
	return appSettings, nil
}

func fetchConcreteSettings(st reflect.Type, prefixAccumulator string) error {
	if st == nil {
		return ErrSettingsCannotBeFetched
	}

	for i := 0; i < st.NumField(); i++ {
		stField := st.Field(i)
		prefix := stField.Tag.Get("env")
		if prefix == "" {
			prefix = stField.Name
		}

		if prefixAccumulator == "" {
			prefixAccumulator = prefix
		} else {
			prefixAccumulator = prefixAccumulator + "_" + prefix
		}

		kind := stField.Type.Kind()
		switch kind {
		case reflect.Struct:
			{
				fetchConcreteSettings(stField.Type, prefixAccumulator)
			}
		case reflect.String:
			{
				val := os.Getenv(prefixAccumulator)
				setfieldVal := reflect.New(stField.Type)
				setfieldVal.Elem().SetString(val)
			}
		}
	}
	return nil
}
