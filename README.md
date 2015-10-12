# mention

mention parses twitter like mentions and hashtags like @gernest and #Tanzania from text input.

# Motivation
I have an idea that I'm implementing, its my attempt to try solving information flow in my country( Tanzania). I nneded to figure out how to compute directions, and heck the simplest approach is to use mentions and hashtags.

You can benefit from `mention` by reading the source code. I have made it simple, and a bit clear for anyone who wants to use `bufio.Scanner` in their project.


# Installation

	go get github.com/gernest/mention
	

# Usage

`mention` is flexible, menaning it is not only limited to `@` and `#` tags. You can choose whatever tag you like and mention will take it from there.

## twitter like mentions

For instance you have the following message

```
hello @gernesti I would like to follow you on twitter
```

And you want to know aho was mentioed in the text.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/gernest/mention"
)

func main() {
	message := "hello @gernest I would like to follow you on twitter"

	tags := mention.GetTags('@', strings.NewReader(message))

	fmt.Println(tags)
}
```

If you run the above example it will print `[gernesti]` is the stdout.

## twitter like hashtags

For instance you have the following message

```
how does it feel to be rejected? #loner
```

And you want to know the hashtags

```go
package main

import (
	"fmt"
	"strings"

	"github.com/gernest/mention"
)

func main() {
	message := "how does it feel to be rejected? #loner"

	tags := mention.GetTags('#', strings.NewReader(message))

	fmt.Println(tags)
}
```

If you run the above example it will print `[loner]` in the stdout.

# The API
mention exposes only one function `GetTags(char rune, src io.Reader) []string`

The first argument `char` is the prefix for your tag, this can be `@` or `#` or whatever unicode character you prefer.Don't be worried by its type `rune` it is just your normal characters but in single quotes. see the examples for more information.

The second argument is the source of the input which can be from texts.

# Contributing

Start with clicking the star button to make the author and his neighbors happy. Then fork the repository and submit a pull request for whatever change you want to be added to this project.

If you have any questions, just open an issue.

# Author
Geofrey Ernest <geofreyernest@live.com>

Twitter  : [@gernesti](https://twitter.com/gernesti)

Facebook : [Geofrey Ernest](https://www.facebook.com/geofrey.ernest.35)


# Licence

This project is released under the MIT licence. See [LICENCE](LICENCE) for more details.
