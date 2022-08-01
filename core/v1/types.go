package v1

import "time"

// CompanyQueryOption company query param.
type CompanyQueryOption struct {
	Pagination       Pagination
	LoadRepositories string
	LoadApplications string
}

// ResourceQueryOption contains resource query options
type ResourceQueryOption struct {
	Pagination    Pagination
	AscendingSort string
}

// Pagination Pagination query params
type Pagination struct {
	Page  string
	Limit string
}

// RepositoryQueryOption repository query option
type RepositoryQueryOption struct {
	Pagination       Pagination
	LoadApplications string
}

// ProcessQueryOption process query option
type ProcessQueryOption struct {
	Pagination Pagination
	Step       string
}

// Process Process struct
type Process struct {
	ProcessId    string                 `bson:"process_id" json:"process_id"`
	CompanyId    string                 `bson:"company_id" json:"company_id"`
	AppId        string                 `bson:"app_id" json:"app_id"`
	RepositoryId string                 `bson:"repository_id" json:"repository_id"`
	CommitId     string                 `bson:"commit_id" json:"commit_id"`
	Data         map[string]interface{} `bson:"data" json:"data"`
	CreatedAt    time.Time              `bson:"created_at" json:"created_at"`
	Branch       string                 `bson:"branch" json:"branch"`
}
