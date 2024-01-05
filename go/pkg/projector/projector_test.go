package projector_test

import (
	"testing"

	"github.com/josiahvehrs/polyglot/pkg/projector"
)

func getData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"fem": "is_great",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(
		&projector.Config{
			Args:      []string{},
			Operation: projector.Print,
			Pwd:       pwd,
			Config:    "Hello, Frontend Masters",
		},
		data,
	)
}

func test(t *testing.T, proj *projector.Projector, key, expected string) {
	value, ok := proj.GetValue(key)
	if !ok {
		t.Errorf("expected to find key \"%s\".", key)
	}
	if value != expected {
		t.Errorf("expected to find %s. got=%s", expected, value)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)
	test(t, proj, "foo", "bar3")
	test(t, proj, "fem", "is_great")
}

func TestSetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)
	test(t, proj, "foo", "bar3")
	proj.SetValue("foo", "baz")
	test(t, proj, "foo", "baz")

	proj.SetValue("fem", "is_super_great")
	test(t, proj, "fem", "is_super_great")

	proj = getProjector("/", data)
	test(t, proj, "fem", "is_great")
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)
	test(t, proj, "foo", "bar3")

	proj.RemoveValue("foo")
	test(t, proj, "foo", "bar2")

	proj.RemoveValue("fem")
	test(t, proj, "fem", "is_great")
}
