package sap_api_output_formatter

type CreditMemoRequest struct {
	ConnectionKey        string `json:"connection_key"`
	Result               bool   `json:"result"`
	RedisKey             string `json:"redis_key"`
	Filepath             string `json:"filepath"`
	APISchema            string `json:"api_schema"`
	CreditMemoRequest    string `json:"credit_memo_request"`
	Deleted              bool   `json:"deleted"`
}    
    
type Header struct {
	CreditMemoRequest              string      `json:"CreditMemoRequest"`
	CreditMemoRequestType          string      `json:"CreditMemoRequestType"`
	SalesOrganization              string      `json:"SalesOrganization"`
	DistributionChannel            string      `json:"DistributionChannel"`
	OrganizationDivision           string      `json:"OrganizationDivision"`
	SalesGroup                     string      `json:"SalesGroup"`
	SalesOffice                    string      `json:"SalesOffice"`
	SalesDistrict                  string      `json:"SalesDistrict"`
	SoldToParty                    string      `json:"SoldToParty"`
	CreationDate                   string      `json:"CreationDate"`
	LastChangeDate                 string      `json:"LastChangeDate"`
	LastChangeDateTime             string      `json:"LastChangeDateTime"`
	PurchaseOrderByCustomer        string      `json:"PurchaseOrderByCustomer"`
	CustomerPurchaseOrderType      string      `json:"CustomerPurchaseOrderType"`
	CustomerPurchaseOrderDate      string      `json:"CustomerPurchaseOrderDate"`
	CreditMemoRequestDate          string      `json:"CreditMemoRequestDate"`
	TotalNetAmount                 string      `json:"TotalNetAmount"`
	TransactionCurrency            string      `json:"TransactionCurrency"`
	SDDocumentReason               string      `json:"SDDocumentReason"`
	PricingDate                    string      `json:"PricingDate"`
	CustomerTaxClassification1     string      `json:"CustomerTaxClassification1"`
	CustomerAccountAssignmentGroup string      `json:"CustomerAccountAssignmentGroup"`
	HeaderBillingBlockReason       string      `json:"HeaderBillingBlockReason"`
	IncotermsClassification        string      `json:"IncotermsClassification"`
	CustomerPaymentTerms           string      `json:"CustomerPaymentTerms"`
	PaymentMethod                  string      `json:"PaymentMethod"`
	BillingDocumentDate            string      `json:"BillingDocumentDate"`
	ReferenceSDDocument            string      `json:"ReferenceSDDocument"`
	ReferenceSDDocumentCategory    string      `json:"ReferenceSDDocumentCategory"`
	CreditMemoReqApprovalReason    string      `json:"CreditMemoReqApprovalReason"`
	SalesDocApprovalStatus         string      `json:"SalesDocApprovalStatus"`
	OverallSDProcessStatus         string      `json:"OverallSDProcessStatus"`
	TotalCreditCheckStatus         string      `json:"TotalCreditCheckStatus"`
	OverallSDDocumentRejectionSts  string      `json:"OverallSDDocumentRejectionSts"`
	OverallOrdReltdBillgStatus     string      `json:"OverallOrdReltdBillgStatus"`
    ToHeaderPartner                string      `json:"to_Partner"`
    ToItem                         string      `json:"to_Item"`	
}

type Item struct {
	CreditMemoRequest             string      `json:"CreditMemoRequest"`
	CreditMemoRequestItem         string      `json:"CreditMemoRequestItem"`
	HigherLevelItem               string      `json:"HigherLevelItem"`
	CreditMemoRequestItemCategory string      `json:"CreditMemoRequestItemCategory"`
	CreditMemoRequestItemText     string      `json:"CreditMemoRequestItemText"`
	PurchaseOrderByCustomer       string      `json:"PurchaseOrderByCustomer"`
	Material                      string      `json:"Material"`
	MaterialByCustomer            string      `json:"MaterialByCustomer"`
	PricingDate                   string      `json:"PricingDate"`
	RequestedQuantity             string      `json:"RequestedQuantity"`
	RequestedQuantityUnit         string      `json:"RequestedQuantityUnit"`
	ItemGrossWeight               string      `json:"ItemGrossWeight"`
	ItemNetWeight                 string      `json:"ItemNetWeight"`
	ItemWeightUnit                string      `json:"ItemWeightUnit"`
	ItemVolume                    string      `json:"ItemVolume"`
	ItemVolumeUnit                string      `json:"ItemVolumeUnit"`
	TransactionCurrency           string      `json:"TransactionCurrency"`
	NetAmount                     string      `json:"NetAmount"`
	MaterialGroup                 string      `json:"MaterialGroup"`
	MaterialPricingGroup          string      `json:"MaterialPricingGroup"`
	ProductTaxClassification1     string      `json:"ProductTaxClassification1"`
	MatlAccountAssignmentGroup    string      `json:"MatlAccountAssignmentGroup"`
	Batch                         string      `json:"Batch"`
	Plant                         string      `json:"Plant"`
	IncotermsClassification       string      `json:"IncotermsClassification"`
	CustomerPaymentTerms          string      `json:"CustomerPaymentTerms"`
	ItemBillingBlockReason        string      `json:"ItemBillingBlockReason"`
	SalesDocumentRjcnReason       string      `json:"SalesDocumentRjcnReason"`
	WBSElement                    string      `json:"WBSElement"`
	ProfitCenter                  string      `json:"ProfitCenter"`
	ReferenceSDDocument           string      `json:"ReferenceSDDocument"`
	ReferenceSDDocumentItem       string      `json:"ReferenceSDDocumentItem"`
	SDProcessStatus               string      `json:"SDProcessStatus"`
	OrderRelatedBillingStatus     string      `json:"OrderRelatedBillingStatus"`
	ToItemPricingElement          string      `json:"to_PricingElement"`
}

type ToHeaderPartner struct {
	CreditMemoRequest   string `json:"CreditMemoRequest"`
	PartnerFunction     string `json:"PartnerFunction"`
	Customer            string `json:"Customer"`
	Supplier            string `json:"Supplier"`
}

type ToItem struct {
	CreditMemoRequest             string      `json:"CreditMemoRequest"`
	CreditMemoRequestItem         string      `json:"CreditMemoRequestItem"`
	HigherLevelItem               string      `json:"HigherLevelItem"`
	CreditMemoRequestItemCategory string      `json:"CreditMemoRequestItemCategory"`
	CreditMemoRequestItemText     string      `json:"CreditMemoRequestItemText"`
	PurchaseOrderByCustomer       string      `json:"PurchaseOrderByCustomer"`
	Material                      string      `json:"Material"`
	MaterialByCustomer            string      `json:"MaterialByCustomer"`
	PricingDate                   string      `json:"PricingDate"`
	RequestedQuantity             string      `json:"RequestedQuantity"`
	RequestedQuantityUnit         string      `json:"RequestedQuantityUnit"`
	ItemGrossWeight               string      `json:"ItemGrossWeight"`
	ItemNetWeight                 string      `json:"ItemNetWeight"`
	ItemWeightUnit                string      `json:"ItemWeightUnit"`
	ItemVolume                    string      `json:"ItemVolume"`
	ItemVolumeUnit                string      `json:"ItemVolumeUnit"`
	TransactionCurrency           string      `json:"TransactionCurrency"`
	NetAmount                     string      `json:"NetAmount"`
	MaterialGroup                 string      `json:"MaterialGroup"`
	MaterialPricingGroup          string      `json:"MaterialPricingGroup"`
	ProductTaxClassification1     string      `json:"ProductTaxClassification1"`
	MatlAccountAssignmentGroup    string      `json:"MatlAccountAssignmentGroup"`
	Batch                         string      `json:"Batch"`
	Plant                         string      `json:"Plant"`
	IncotermsClassification       string      `json:"IncotermsClassification"`
	CustomerPaymentTerms          string      `json:"CustomerPaymentTerms"`
	ItemBillingBlockReason        string      `json:"ItemBillingBlockReason"`
	SalesDocumentRjcnReason       string      `json:"SalesDocumentRjcnReason"`
	WBSElement                    string      `json:"WBSElement"`
	ProfitCenter                  string      `json:"ProfitCenter"`
	ReferenceSDDocument           string      `json:"ReferenceSDDocument"`
	ReferenceSDDocumentItem       string      `json:"ReferenceSDDocumentItem"`
	SDProcessStatus               string      `json:"SDProcessStatus"`
	OrderRelatedBillingStatus     string      `json:"OrderRelatedBillingStatus"`
	ToItemPricingElement          string      `json:"to_PricingElement"`
}

type ToItemPricingElement struct {
	CreditMemoRequest              string `json:"CreditMemoRequest"`
	CreditMemoRequestItem          string `json:"CreditMemoRequestItem"`
	PricingProcedureStep           string `json:"PricingProcedureStep"`
	PricingProcedureCounter        string `json:"PricingProcedureCounter"`
	ConditionApplication           string `json:"ConditionApplication"`
	ConditionType                  string `json:"ConditionType"`
	PricingDateTime                string `json:"PricingDateTime"`
	PriceConditionDeterminationDte string `json:"PriceConditionDeterminationDte"`
	ConditionCalculationType       string `json:"ConditionCalculationType"`
	ConditionBaseValue             string `json:"ConditionBaseValue"`
	ConditionRateValue             string `json:"ConditionRateValue"`
	ConditionCurrency              string `json:"ConditionCurrency"`
	ConditionQuantity              string `json:"ConditionQuantity"`
	ConditionQuantityUnit          string `json:"ConditionQuantityUnit"`
	ConditionToBaseQtyNmrtr        string `json:"ConditionToBaseQtyNmrtr"`
	ConditionToBaseQtyDnmntr       string `json:"ConditionToBaseQtyDnmntr"`
	ConditionCategory              string `json:"ConditionCategory"`
	PricingScaleType               string `json:"PricingScaleType"`
	ConditionRecord                string `json:"ConditionRecord"`
	ConditionSequentialNumber      string `json:"ConditionSequentialNumber"`
	TaxCode                        string `json:"TaxCode"`
	ConditionAmount                string `json:"ConditionAmount"`
	TransactionCurrency            string `json:"TransactionCurrency"`
	PricingScaleBasis              string `json:"PricingScaleBasis"`
	ConditionScaleBasisValue       string `json:"ConditionScaleBasisValue"`
	ConditionScaleBasisUnit        string `json:"ConditionScaleBasisUnit"`
	ConditionScaleBasisCurrency    string `json:"ConditionScaleBasisCurrency"`
	ConditionIsManuallyChanged     bool   `json:"ConditionIsManuallyChanged"`
}
