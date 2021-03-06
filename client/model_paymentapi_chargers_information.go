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

type PaymentapiChargersInformation struct {
	BearerCode string `json:"bearer_code,omitempty"`
	SenderCharges []ChargersInformationSenderCharges `json:"sender_charges,omitempty"`
	ReceiverChargesAmount string `json:"receiver_charges_amount,omitempty"`
	ReceiverChargesCurrency string `json:"receiver_charges_currency,omitempty"`
}
