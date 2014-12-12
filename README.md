urlregex
======

express-like named url parameters extracting from url

## Example
```Go
package main

import (
	"fmt"
	"github.com/ivpusic/urlregex"
)

func main() {
  	// express-like url named params
	reg := urlregex.Pattern("some/:name/path/:other/")
	
	// extract named params values from given url string
	res, err := reg.Match("some/123/path/456/")

  	// err returned -> no matches
	if err != nil {
		fmt.Println("no matches")
		return
	}

  // we have matches
	fmt.Println("found matches")
	for k, v := range res {
		fmt.Println(k + ": " + v)
	}
}
```

This will output:
```
found matches
name: 123
other: 456
```

You can also pass url pattern without names parameters, and later check if given url matches.
```Go
package main

import (
	"fmt"
	"github.com/ivpusic/urlregex"
)

func main() {
	reg := urlregex.Pattern("some/123/path/456")
	
	// in this case let's pass something invalid
	_, err := reg.Match("some/123/path/invalid")

	if err != nil {
		fmt.Println("not matched")
		return
	}

	fmt.Println("matched")
}
```
This will output
```
not matched
```
However if we passed ``some/123/path/456`` to ``Match`` method, then it would output ``matched``.

#### Access actual go regex instance
```Go
package main

import (
	"fmt"
	"github.com/ivpusic/urlregex"
)

func main() {
	reg := urlregex.Pattern("some/:name/path/:other/")
	// native generated *Regex instance
	fmt.Println(reg.Regex)
}

```

# License
MIT
