package v1

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
