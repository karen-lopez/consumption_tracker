package dtos

type MeterData struct {
	MeterID            int    `json:"meter_id"`
	Address            string `json:"address"`
	Active             []int  `json:"active"`
	ReactiveInductive  []int  `json:"reactive_inductive"`
	ReactiveCapacitive []int  `json:"reactive_capacitive"`
	Exported           []int  `json:"exported"`
}

type ConsumptionResponse struct {
	Period    []string    `json:"period"`
	DataGraph []MeterData `json:"data_graph"`
}
