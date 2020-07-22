package repositories

type RepositoryResult struct {
	Result interface{} `json:"result"`
	Error  error	`json:"error"`
}
