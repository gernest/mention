package mention

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MentionSuite struct{}

var _ = Suite(&MentionSuite{})

func (s *MentionSuite) TestGetTags(c *C) {

	sample := []struct {
		src  string
		tags []Tag
	}{
		{
			"@gernest",
			[]Tag{
				{'@', "gernest", 0},
			},
		},
		{
			"@gernest ",
			[]Tag{
				{'@', "gernest", 0},
			},
		},
		{
			"@gernest@mwanza hello",
			[]Tag{
				{'@', "gernest@mwanza", 0},
			},
		},
		{
			"@gernest @mwanza @mwanza",
			[]Tag{
				{'@', "gernest", 0},
				{'@', "mwanza", 9},
				{'@', "mwanza", 17},
			},
		},
		{
			"Hello to @gernest. Maybe we can do it together @mwanza",
			[]Tag{
				{'@', "gernest", 9},
				{'@', "mwanza", 47},
			},
		},
		{
			" @gernest @mwanza",
			[]Tag{
				{'@', "gernest", 1},
				{'@', "mwanza", 10},
			},
		},
		{
			" @gernest @mwanza ",
			[]Tag{
				{'@', "gernest", 1},
				{'@', "mwanza", 10},
			},
		},
		{
			" @gernest @mwanza @tanzania",
			[]Tag{
				{'@', "gernest", 1},
				{'@', "mwanza", 10},
				{'@', "tanzania", 18},
			},
		},
		{
			" @gernest,@mwanza/Tanzania ",
			[]Tag{
				{'@', "gernest", 1},
				{'@', "mwanza", 10},
			},
		},
		{
			"how does it feel to be rejected? @ it is @loner tt ggg sjdsj dj @linker ",
			[]Tag{
				{'@', "loner", 41},
				{'@', "linker", 64},
			},
		},
		{
			"This @gernest is @@@@ @@@ @@ @ @,, @, @mwanza,",
			[]Tag{
				{'@', "gernest", 5},
				{'@', "mwanza", 38},
			},
		},
		{
			"hello@world",
			nil,
		},
	}
	terms := []rune(",/. ")

	for _, v := range sample {
		c.Assert(GetTags('@', v.src, terms...), DeepEquals, v.tags, Commentf("Failed: %+v", v))
	}

	// use default terminators
	c.Assert(GetTags('@', "hello @test"), DeepEquals, []Tag{{'@', "test", 6}})
}
