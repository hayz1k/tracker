package status

const (
	StatusCreated            = "Created"
	StatusSortingCenter      = "Sorting Center"
	StatusIntermediateCenter = "Intermediate Center"
	StatusInTransit          = "In Transit"
	StatusArrivedDestination = "Arrived at Destination"
	StatusOutForDelivery     = "Out for Delivery"
	StatusDelivered          = "Delivered"
	StatusAwaitingPickup     = "Awaiting Pickup"
	StatusReturnedToSender   = "Returned to Sender"
	StatusCustomsHold        = "Customs Hold"
)

var Chains = [][]string{
	// цепочка 1: стандартная доставка
	{StatusCreated, StatusSortingCenter, StatusIntermediateCenter, StatusInTransit, StatusArrivedDestination, StatusOutForDelivery, StatusDelivered, StatusDelivered, StatusDelivered, StatusDelivered},

	// цепочка 2: доставка с ожиданием на сортировке
	{StatusCreated, StatusSortingCenter, StatusSortingCenter, StatusInTransit, StatusArrivedDestination, StatusOutForDelivery, StatusDelivered, StatusDelivered, StatusDelivered, StatusDelivered},

	// цепочка 3: международная доставка
	{StatusCreated, StatusSortingCenter, StatusInTransit, StatusIntermediateCenter, StatusCustomsHold, StatusInTransit, StatusArrivedDestination, StatusOutForDelivery, StatusDelivered, StatusDelivered},

	// цепочка 4: самовывоз
	{StatusCreated, StatusSortingCenter, StatusInTransit, StatusArrivedDestination, StatusAwaitingPickup, StatusDelivered, StatusDelivered, StatusDelivered, StatusDelivered, StatusDelivered},

	// цепочка 5: возврат
	{StatusCreated, StatusSortingCenter, StatusInTransit, StatusArrivedDestination, StatusOutForDelivery, StatusReturnedToSender, StatusReturnedToSender, StatusReturnedToSender, StatusReturnedToSender, StatusReturnedToSender},
}
