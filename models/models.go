package models

// A Resource struct to map Resource
type Resource struct {
	CategoryId string `json:"categoryId"`
	Id string `json:"id"`
	Identifier string `json:"identifier"`
	Name string `json:"name"`
	UomId string `json:"uomId"`
}

// An Interactions struct to map Interactions
type Interactions struct {
	Id string `json:"id"`
	Quantity int `json:"quantity"`
	Resource Resource `json:"resource"`
}

// A Resources struct to map Resources
type Resources struct {
	Id string `json:"id"`
	Interactions []Interactions `json:"interactions"`
}

// A Docket struct to map the Docket
type Docket struct {
    CreatedAt string `json:"createdAt"`
	Date string `json:"date"`
	DocketNumber string `json:"docketNumber"`
	Id string `json:"id"`
	Organisation string `json:"organisation"`
	OrganisationGroup string `json:"organisationGroup"`
	Resources []Resources `json:"resources"`
	Status string `json:"status"`
	UpdatedAt string `json:"updatedAt"`
}
