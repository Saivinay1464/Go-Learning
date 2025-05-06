package models

type Food struct {
	Name     string  `json:"name"`
	Serving  int     `json:"serving"`
	Calories int     `json:"calories"`
	Fat      float64 `json:"fat"`
	Protein  float64 `json:"protein"`
	Carbs    float64 `json:"carbs"`
	Fiber    float64 `json:"fiber"`
}
