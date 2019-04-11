/*
 * FinTech api
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Contact: none@example.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Payment model
type PaymentapiPayment struct {
	Type_ string `json:"type,omitempty"`
	Id string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
	OrganisationId string `json:"organisation_id"`
	Attributes *PaymentapiPaymentAttributes `json:"attributes,omitempty"`
}