package internal

type Place struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	Location []float64 `json:"location"`
}
