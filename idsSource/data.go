package idsSource



type Infects struct {
	Filters struct {
		Filter []struct {
			ID          string `json:"id"`
			Rule        string `json:"rule"`
			Description string `json:"description"`
			Tags        struct {
				Tag []string `json:"tag"`
			} `json:"tags"`
			Impact string `json:"impact"`
		} `json:"filter"`
	} `json:"filters"`
}
