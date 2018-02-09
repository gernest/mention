package mention

import (
	"fmt"
)

func ExampleGetTags_mention() {
	msg := " hello @gernest"
	tags := GetTags('@', msg)
	fmt.Println(tags)

	//Output:
	//[gernest]
}

func ExampleGetTags_hashtag() {
	msg := " viva la #tanzania"
	tags := GetTags('#', msg)
	fmt.Println(tags)

	//Output:
	//[tanzania]
}
