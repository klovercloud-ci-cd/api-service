package v1

// CompanyQueryOption company query param.
type CompanyQueryOption struct {
	Pagination       Pagination
	LoadRepositories string
	LoadApplications string
}

// Pagination Pagination query params
type Pagination struct {
	Page  string
	Limit string
}
