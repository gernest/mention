package mention

import (
	"reflect"
	"testing"
)

func TestGetTags(t *testing.T) {

	sample := []struct {
		src string
		tag []string
	}{
		{"@gernest", []string{"gernest"}},
		{"@gernest ", []string{"gernest"}},
		{"@gernest@mwanza hello", []string{"gernest@mwanza"}},
		{"@gernest @mwanza", []string{"gernest", "mwanza"}},
		{"Hello to @gernest. Maybe we can do it together @mwanza", []string{"gernest", "mwanza"}},
		{" @gernest @mwanza", []string{"gernest", "mwanza"}},
		{" @gernest @mwanza ", []string{"gernest", "mwanza"}},
		{" @gernest @mwanza @tanzania", []string{"gernest", "mwanza", "tanzania"}},
		{" @gernest,@mwanza/Tanzania ", []string{"gernest", "mwanza"}},
		{"how does it feel to be rejected? @ it is @loner tt ggg sjdsj dj @linker ", []string{"loner", "linker"}},
		{"This @gernest is @@@@ @@@ @@ @ @,, @, @mwanza,", []string{"gernest", "mwanza"}},
		{"hello@world", nil},
	}
	terms := []rune(",/. ")

	for _, v := range sample {
		tag := GetTags('@', v.src, terms...)
		if !reflect.DeepEqual(v.tag, tag) {
			t.Errorf("expected  %v got %v for %s", v.tag, tag, v.src)
		}

	}
}
