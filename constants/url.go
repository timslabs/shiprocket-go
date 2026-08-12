package constants

const (
	VERSION_V1 = "v1"

	BASE_URL = "https://apiv2.shiprocket.in"

	// EXTERNAL_PREFIX is used as: fmt.Sprintf("/%s%s", EXTERNAL_PREFIX, PATH)
	EXTERNAL_PREFIX = "v1/external"

	// Auth
	AUTH_LOGIN_URL  = "/auth/login"
	AUTH_LOGOUT_URL = "/auth/logout"

	// Couriers
	COURIER_SERVICEABILITY_URL  = "/courier/serviceability/"
	COURIER_LIST_COUNTS_URL     = "/courier/courierListWithCounts"
	COURIER_ASSIGN_AWB_URL      = "/courier/assign/awb"
	COURIER_GENERATE_PICKUP_URL = "/courier/generate/pickup"
	COURIER_GENERATE_LABEL_URL  = "/courier/generate/label"
	COURIER_TRACK_AWB_URL       = "/courier/track/awb"
	COURIER_TRACK_SHIPMENT_URL  = "/courier/track/shipment"
	COURIER_TRACK_URL           = "/courier/track"
	COURIER_TRACK_AWBS_URL      = "/courier/track/awbs"
	MANIFESTS_GENERATE_URL      = "/manifests/generate"
	MANIFESTS_PRINT_URL         = "/manifests/print"
	CANCEL_SHIPMENT_AWBS_URL    = "/orders/cancel/shipment/awbs"

	// Orders
	ORDERS_URL                   = "/orders"
	ORDERS_CREATE_ADHOC_URL      = "/orders/create/adhoc"
	ORDERS_CREATE_URL            = "/orders/create"
	ORDERS_CREATE_RETURN_URL     = "/orders/create/return"
	ORDERS_UPDATE_ADHOC_URL      = "/orders/update/adhoc"
	ORDERS_CANCEL_URL            = "/orders/cancel"
	ORDERS_SHOW_URL              = "/orders/show"
	ORDERS_PROCESSING_RETURN_URL = "/orders/processing/return"
	ORDERS_ADDRESS_UPDATE_URL    = "/orders/address/update"
	ORDERS_ADDRESS_PICKUP_URL    = "/orders/address/pickup"
	ORDERS_FULFILL_URL           = "/orders/fulfill"
	ORDERS_MAPPING_URL           = "/orders/mapping"
	ORDERS_EXPORT_URL            = "/orders/export"
	ORDERS_IMPORT_URL            = "/orders/import"
	ORDERS_PRINT_INVOICE_URL     = "/orders/print/invoice"

	// Shipments
	SHIPMENTS_URL                = "/shipments"
	SHIPMENTS_CREATE_FORWARD_URL = "/shipments/create/forward-shipment"
	SHIPMENTS_CREATE_RETURN_URL  = "/shipments/create/return-shipment"

	// Pickup
	PICKUP_LIST_URL = "/settings/company/pickup"
	PICKUP_ADD_URL  = "/settings/company/addpickup"

	// Products
	PRODUCTS_URL        = "/products"
	PRODUCTS_SHOW_URL   = "/products/show"
	PRODUCTS_SAMPLE_URL = "/products/sample"
	PRODUCTS_IMPORT_URL = "/products/import"

	// Inventory
	INVENTORY_URL = "/inventory"

	// Listings
	LISTINGS_URL                 = "/listings"
	LISTINGS_LINK_URL            = "/listings/link"
	LISTINGS_IMPORT_URL          = "/listings/import"
	LISTINGS_EXPORT_MAPPED_URL   = "/listings/export/mapped"
	LISTINGS_EXPORT_UNMAPPED_URL = "/listings/export/unmapped"
	LISTINGS_SAMPLE_URL          = "/listings/sample"

	// Channels
	CHANNELS_URL         = "/channels"
	COUNTRIES_URL        = "/countries"
	COUNTRIES_SHOW_URL   = "/countries/show"
	POSTCODE_DETAILS_URL = "/open/postcode/details"

	// Account / billing
	ACCOUNT_WALLET_BALANCE_URL = "/account/details/wallet-balance"
	ACCOUNT_STATEMENT_URL      = "/account/details/statement"
	BILLING_DISCREPANCY_URL    = "/billing/discrepancy"

	// NDR
	NDR_ALL_URL = "/ndr/all"
	NDR_URL     = "/ndr"

	// Imports
	ERRORS_CHECK_URL = "/errors"

	// International
	INTERNATIONAL_ORDERS_CREATE_ADHOC_URL = "/international/orders/create/adhoc"
	INTERNATIONAL_ORDERS_UPDATE_ADHOC_URL = "/international/orders/update/adhoc"
	INTERNATIONAL_ORDERS_TRACK_URL        = "/international/orders/track"
	INTERNATIONAL_COURIER_SERVICEABILITY  = "/international/courier/serviceability"
	INTERNATIONAL_COURIER_ASSIGN_AWB      = "/international/courier/assign/awb"
	INTERNATIONAL_MANIFESTS_GENERATE      = "/international/manifests/generate"
	INTERNATIONAL_SHIPMENTS_FORWARD       = "/international/shipments/create/forward-shipment"
	INTERNATIONAL_ADD_BANK_DETAILS        = "/international/settings/add-bank-details"
	INTERNATIONAL_KYC                     = "/international/settings/international_kyc"

	// Warehouse (not under /v1/external)
	WAREHOUSE_PREFIX             = "v1/warehouse"
	WAREHOUSE_SRF_SERVICEABILITY = "/srf-serviceability"
)
