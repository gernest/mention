package mention

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetTag(t *testing.T) {
	sample := []struct {
		src string
		tag []string
	}{
		{"@gernest", []string{"gernest"}},
		{"@gernest ", []string{"gernest"}},
		{"@gernest@mwanza hello", []string{"gernest", "mwanza"}},
		{"@gernest @mwanza", []string{"gernest", "mwanza"}},
		{" @gernest @mwanza", []string{"gernest", "mwanza"}},
		{" @gernest @mwanza ", []string{"gernest", "mwanza"}},
		{" @gernest @mwanza @tanzania", []string{"gernest", "mwanza", "tanzania"}},
	}
	for _, v := range sample {
		tag := GetTags('@', strings.NewReader(v.src))

		if !reflect.DeepEqual(v.tag, tag) {
			t.Errorf("expected  %v got %v for %s", v.tag, tag, v.src)
		}

	}

}
