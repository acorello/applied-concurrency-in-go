package orderstatus

type OrderStatus string

const (
	New       OrderStatus = "New"
	Completed OrderStatus = "Completed"
	Rejected  OrderStatus = "Rejected"
)
