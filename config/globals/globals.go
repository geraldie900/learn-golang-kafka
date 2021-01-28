package globals

import "context"

// constant variable
const (
	BranchColletion                     = "branches"
	AliasBranchCollection               = "b"
	AgentCollection                     = "agents"
	AliasAgentCollection                = "a"
	MerchantCollection                  = "merchants"
	AliasMerchantCollection             = "m"
	MerchantAdminCollection             = "merchant_admins"
	AliasMerchantAdminColletion         = "ma"
	UserCollection                      = "users"
	AliasUserCollection                 = "u"
	RoleCollection                      = "roles"
	AliasRoleCollection                 = "r"
	CustomerRegistrationCollection      = "customer_registrations"
	AliasCustomerRegistrationCollection = "cr"
	ImageCollection                     = "images"
	AliasImageCollection                = "i"
)

// Ctx globally for accessing context.Background()
var (
	Ctx = context.Background()
)
