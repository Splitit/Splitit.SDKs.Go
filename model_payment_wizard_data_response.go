/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// PaymentWizardDataResponse struct for PaymentWizardDataResponse
type PaymentWizardDataResponse struct {
	RequestedNumberOfInstallments string `json:"RequestedNumberOfInstallments,omitempty"`
	SuccessExitURL string `json:"SuccessExitURL,omitempty"`
	ErrorExitURL string `json:"ErrorExitURL,omitempty"`
	CancelExitURL string `json:"CancelExitURL,omitempty"`
	SuccessAsyncUrl string `json:"SuccessAsyncUrl,omitempty"`
	ViewName string `json:"ViewName,omitempty"`
	IsOpenedInIframe bool `json:"IsOpenedInIframe"`
	PaymentFormMessage string `json:"PaymentFormMessage,omitempty"`
	ShowAddressElements string `json:"ShowAddressElements,omitempty"`
	CurrencyDisplay *Currency `json:"CurrencyDisplay,omitempty"`
	ForceDisplayImportantNotes bool `json:"ForceDisplayImportantNotes"`
	ShowShopperDetailsExpendedOnStart bool `json:"ShowShopperDetailsExpendedOnStart"`
	ShowPaymentScheduleRequiredCredit bool `json:"ShowPaymentScheduleRequiredCredit"`
	IsShopperEmailMandatory bool `json:"IsShopperEmailMandatory"`
	IsShopperPhoneMandatory bool `json:"IsShopperPhoneMandatory"`
	NumberOfInstallmentsSelectionsOption string `json:"NumberOfInstallmentsSelectionsOption,omitempty"`
	AddressIsReadonly bool `json:"AddressIsReadonly"`
	LogoURL string `json:"LogoURL,omitempty"`
	PrivacyPolicyUrl string `json:"PrivacyPolicyUrl,omitempty"`
	TermsAndConditionsUrl string `json:"TermsAndConditionsUrl,omitempty"`
	LearnMoreUrl string `json:"LearnMoreUrl,omitempty"`
	PaymentFormMessages *[]PaymentFormMessage `json:"PaymentFormMessages,omitempty"`
}
