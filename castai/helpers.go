package castai

import (
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

// NewDefaultClient returns a CastAI client with typed responses and configured
// to use the provided API key.
func NewDefaultClient(
	apiKey string, opts ...ClientOption,
) (*ClientWithResponses, error) {
	apiKey = strings.TrimSpace(apiKey)
	apiKeyProvider, err := securityprovider.NewSecurityProviderApiKey(
		"header", "X-API-Key", apiKey,
	)
	if err != nil {
		return nil, err
	}
	opts = append(opts,
		WithRequestEditorFn(apiKeyProvider.Intercept),
	)
	client, err := NewClientWithResponses(DefaultServer, opts...)
	return client, err
}

const DefaultServer = "https://api.cast.ai"
