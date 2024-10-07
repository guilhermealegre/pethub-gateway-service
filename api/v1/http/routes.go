package http

import (
	"net/http"

	infra "bitbucket.org/asadventure/be-infrastructure-lib/http"
)

var (
	GroupV1  = infra.NewGroup("v1")
	GroupV1P = GroupV1.Group("p")

	// Gateway
	GatewayAlive                  = GroupV1.NewEndpoint("/alive", http.MethodGet)
	GatewayPublicAlive            = GroupV1P.NewEndpoint("/alive/gateway", http.MethodGet)
	GatewayPublicAccesssClearance = GroupV1P.NewEndpoint("/access-clearance", http.MethodGet)

	// Docs
	GroupV1Documentation = GroupV1P.Group("documentation")
	SwaggerDocs          = GroupV1Documentation.NewEndpoint("/:service/docs", http.MethodGet)
	SwaggerSwagger       = GroupV1Documentation.NewEndpoint("/:service/swagger", http.MethodGet)
	SwaggerJson          = GroupV1Documentation.NewEndpoint("/:service/swagger.json", http.MethodGet)

	// Fallback
	PublicAlive = GroupV1P.NewEndpoint("/alive/:service", http.MethodGet)

	// Store
	GroupV1Store  = GroupV1.Group("store")
	GroupV1PStore = GroupV1P.Group("store")
	StoreByID     = GroupV1Store.NewEndpoint("/:id", http.MethodGet)
	StoreList     = GroupV1PStore.NewEndpoint("/list", http.MethodGet)
	StoreByIP     = GroupV1PStore.NewEndpoint("", http.MethodGet)

	// User
	GroupV1User          = GroupV1.Group("auth")
	UserLogOff           = GroupV1User.NewEndpoint("/logoff", http.MethodPut)
	GroupV1PUser         = GroupV1P.Group("auth")
	UserScannerList      = GroupV1PUser.NewEndpoint("/scanner/list", http.MethodGet)
	UserScannerLogin     = GroupV1PUser.NewEndpoint("/scanner/login", http.MethodPost)
	UserScannerAuthorize = GroupV1PUser.NewEndpoint("/scanner/authorize", http.MethodGet)
	UserByID             = GroupV1User.NewEndpoint("/:id", http.MethodGet)

	// Customer
	GroupV1Customer       = GroupV1.Group("customer")
	CustomerByID          = GroupV1Customer.NewEndpoint("/:id", http.MethodGet)
	CustomerByCard        = GroupV1Customer.NewEndpoint("/card/:uid", http.MethodGet)
	CustomerTitles        = GroupV1Customer.NewEndpoint("/titles", http.MethodGet)
	CustomerUpdate        = GroupV1Customer.NewEndpoint("/:id", http.MethodPut)
	CustomerSearchGeneral = GroupV1Customer.NewEndpoint("/search", http.MethodPost)
	CustomerAddressList   = GroupV1Customer.NewEndpoint("/:id/addresses", http.MethodGet)
	AddCustomerAddress    = GroupV1Customer.NewEndpoint("/:id/address", http.MethodPost)
	EditCustomerAddress   = GroupV1Customer.NewEndpoint("/:id/address/:aid", http.MethodPut)

	// Order
	GroupV1Order                         = GroupV1.Group("order")
	GroupV1POrder                        = GroupV1P.Group("order")
	OrderList                            = GroupV1Order.NewEndpoint("/list", http.MethodGet)
	OrderReadyForPickupList              = GroupV1Order.NewEndpoint("/status/ready-for-pickup/list", http.MethodGet)
	OrderProductServiceList              = GroupV1Order.NewEndpoint("/product/service/list", http.MethodPost)
	OrderCreate                          = GroupV1Order.NewEndpoint("", http.MethodPost)
	OrderAddProduct                      = GroupV1Order.NewEndpoint("/:oid/product", http.MethodPost)
	OrderSaveDelivery                    = GroupV1Order.NewEndpoint("/:oid/delivery", http.MethodPost)
	OrderFinish                          = GroupV1Order.NewEndpoint("/:oid/finish", http.MethodPost)
	OrderRemoveProduct                   = GroupV1Order.NewEndpoint("/:oid/product/:pid", http.MethodDelete)
	OrderUpdateProductService            = GroupV1Order.NewEndpoint("/:oid/product-service/:psid", http.MethodPut)
	OrderServices                        = GroupV1Order.NewEndpoint("/services", http.MethodGet)
	ServiceStatusCommentList             = GroupV1Order.NewEndpoint("/service/:sid/status-comment/list", http.MethodGet)
	ProductTypes                         = GroupV1Order.NewEndpoint("/product-types", http.MethodGet)
	ProductSubTypes                      = GroupV1Order.NewEndpoint("/service/:sid/product-type/:pid/subtypes/details", http.MethodGet)
	DeliveryOption                       = GroupV1Order.NewEndpoint("/:oid/delivery", http.MethodGet)
	OrderDetails                         = GroupV1Order.NewEndpoint("/:oid/details", http.MethodGet)
	OrderStatus                          = GroupV1Order.NewEndpoint("/:oid/status", http.MethodPost)
	ItemWashReadyForShipmentToWh         = GroupV1Order.NewEndpoint("/:oid/product-service/:psid/wash/ready-for-shipment-to-warehouse", http.MethodPut)
	ItemWashArrivedWarehouse             = GroupV1Order.NewEndpoint("/:oid/product-service/:psid/wash/arrived-warehouse", http.MethodPut)
	ItemWashBeingProcessed               = GroupV1Order.NewEndpoint("/:oid/product-service/:psid/wash/being-processed", http.MethodPut)
	ItemWashHoldBeforeService            = GroupV1Order.NewEndpoint("/:oid/product-service/:psid/wash/hold-before-service", http.MethodPut)
	ItemWashProcessed                    = GroupV1Order.NewEndpoint("/:oid/product-service/:psid/wash/processed", http.MethodPut)
	ItemWashHoldAfterService             = GroupV1Order.NewEndpoint("/:oid/product-service/:psid/wash/hold-after-service", http.MethodPut)
	ItemWashOnTheWayToStore              = GroupV1Order.NewEndpoint("/:oid/product-service/:psid/wash/on-the-way-to-store", http.MethodPut)
	ItemWashReadyForShipmentToCustomer   = GroupV1Order.NewEndpoint("/:oid/product-service/:psid/wash/ready-for-shipment", http.MethodPut)
	ItemWashReturnUnsuccessfulItem       = GroupV1Order.NewEndpoint("/:oid/product-service/:psid/wash/return-unsuccessful-item", http.MethodPut)
	ItemWashReadyForPickup               = GroupV1Order.NewEndpoint("/:oid/product-service/:psid/wash/ready-for-pickup", http.MethodPut)
	ItemWashCanceled                     = GroupV1Order.NewEndpoint("/:oid/product-service/:psid/wash/canceled", http.MethodPut)
	OrderCustomerHistory                 = GroupV1Order.NewEndpoint("/customer/:id/history", http.MethodGet)
	StoreOrderProductLabelExist          = GroupV1Order.NewEndpoint("/product/store/label/:lid/exists", http.MethodGet)
	StoreAssignOrderProductItemToBox     = GroupV1Order.NewEndpoint("/product/store/label/:lid/box/:bid", http.MethodPost)
	WarehouseOrderProductLabelExist      = GroupV1Order.NewEndpoint("/product/warehouse/label/:lid/exists", http.MethodGet)
	WarehouseAssignOrderProductItemToBox = GroupV1Order.NewEndpoint("/product/warehouse/label/:lid/box/:bid", http.MethodPost)
	OrderProductLabel                    = GroupV1Order.NewEndpoint("/product/:pid/label", http.MethodGet)
	OrderProductServiceByLabel           = GroupV1POrder.NewEndpoint("/product/service/label/:lid", http.MethodGet)
	PublicItemMarkAsPaid                 = GroupV1POrder.NewEndpoint("/product/paid", http.MethodPost)
	OrderStatusOnTheWayToTheCustomer     = GroupV1Order.NewEndpoint("/:oid/on-the-way-to-the-customer", http.MethodPut)
	OrderGetOrderProductServiceComments  = GroupV1Order.NewEndpoint("/product/service/:id/comments", http.MethodGet)

	// Uploader
	GroupV1Uploader  = GroupV1.Group("uploader")
	GroupV1PUploader = GroupV1P.Group("uploader")
	ImageUpload      = GroupV1Uploader.NewEndpoint("/image/upload", http.MethodPost)
	ImagePUpload     = GroupV1PUploader.NewEndpoint("/image/upload", http.MethodPost)

	// Logging
	GroupV1Logging     = GroupV1.Group("logging")
	GroupV1PLogging    = GroupV1P.Group("logging")
	LoggingCreateFeLog = GroupV1PLogging.NewEndpoint("/log", http.MethodPost)
)
