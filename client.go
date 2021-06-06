package deepl

type Client struct {
	AuthenticationKey string
}

// NewClient returns new DeepL API client with an authentication key.
func NewClient(authnKey string) *Client {
	return &Client{
		AuthenticationKey: authnKey,
	}
}
