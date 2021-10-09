deepl-sdk-go
===

This is a an unofficial Go SDK for using the DeepL API.

# Usage
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmichimani%2Fdeepl-sdk-go.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmichimani%2Fdeepl-sdk-go?ref=badge_shield)


```bash
go get github.com/michimani/deepl-sdk-go
```

# Sample

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/michimani/deepl-sdk-go"
	"github.com/michimani/deepl-sdk-go/params"
	"github.com/michimani/deepl-sdk-go/types"
)

func main() {
	client, err := deepl.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	text := []string{
		"こんにちは",
		"これはサンプルテキストです。",
	}
	params := &params.TranslateTextParams{
		TargetLang: types.TargetLangEN,
		Text:       text,
	}

	res, errRes, err := c.TranslateText(context.TODO(), params)

	if err != nil {
		fmt.Println(err)
	}

	if errRes != nil {
		fmt.Println("ErrorResponse", errRes.Message)
	}

	for i := range res.Translations {
		fmt.Printf("%s -> %s\n", text[i], res.Translations[i].Text)
	}
}
```

```bash
$ DEEPL_API_AUTHN_KEY="your-authn-key" DEEPL_API_PLAN="free" go run main.go

こんにちは -> hello
これはサンプルテキストです。 -> This is a sample text.
```


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmichimani%2Fdeepl-sdk-go.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmichimani%2Fdeepl-sdk-go?ref=badge_large)