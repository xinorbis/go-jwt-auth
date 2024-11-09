package api_key

type ApiKey struct {
	ID     uint64
	UserID uint64
	Key    string
}

func New(userId uint64, apiKey string) ApiKey {
	apiKeyDTO := ApiKey{}
	apiKeyDTO.UserID = userId
	apiKeyDTO.Key = apiKey

	return apiKeyDTO
}
