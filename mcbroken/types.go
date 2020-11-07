package mcbroken

// Mcbroken holds ice cream machine failure rates in all avaailable cities
type Mcbroken struct {
	Cities []struct {
		City   string  `json:"city"`
		Broken float64 `json:"broken"`
	} `json:"cities"`
	Broken float64 `json:"broken"`
}
