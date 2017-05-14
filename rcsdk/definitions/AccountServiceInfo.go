package definitions

type AccountServiceInfo struct {
	Brand           BrandInfo       `json:"brand,omitempty"`
	ServicePlan     ServicePlanInfo `json:"servicePlan,omitempty"`
	BillingPlan     BillingPlanInfo `json:"billingPlan,omitempty"`
	ServiceFeatures `json:"serviceFeatures,omitempty"`
	Limits          AccountLimits `json:"limits,omitempty"`
	Uri             string        `json:"uri,omitempty"`
	ServicePlanName string        `json:"servicePlanName,omitempty"`
}