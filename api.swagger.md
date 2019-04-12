# FinTech api
## Version: 1.0

**Contact information:**  
FinTech api back  
https://github.com/jozuenoone/rest-api  
none@example.com  

### Security
**ApiKeyAuth**  

|apiKey|*API Key*|
|---|---|
|Name|X-API-Key|
|In|header|

### /v1/transaction/payments

#### GET
##### Summary:

Get or List payments

##### Description:

If empty payment ID is sent, full list of payments is returned.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | query |  | No | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [paymentapiPaymentServiceResponse](#paymentapipaymentserviceresponse) |
| 403 | Returned when the user does not have permission to access the resource. | [paymentapiError](#paymentapierror) |
| 404 | Returned when the resource does not exist. | [paymentapiError](#paymentapierror) |
| 500 | Returned when the internal server error occur. | [paymentapiError](#paymentapierror) |

#### POST
##### Summary:

Create payment

##### Description:

Creates new payment and returns full payment data with special fields attached eg. ID, organization etc.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| body | body |  | Yes | [paymentapiPaymentAttributes](#paymentapipaymentattributes) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [paymentapiPaymentServiceResponse](#paymentapipaymentserviceresponse) |
| 403 | Returned when the user does not have permission to access the resource. | [paymentapiError](#paymentapierror) |
| 404 | Returned when the resource does not exist. | [paymentapiError](#paymentapierror) |
| 500 | Returned when the internal server error occur. | [paymentapiError](#paymentapierror) |

### /v1/transaction/payments/{id}

#### GET
##### Summary:

Get or List payments

##### Description:

If empty payment ID is sent, full list of payments is returned.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [paymentapiPaymentServiceResponse](#paymentapipaymentserviceresponse) |
| 403 | Returned when the user does not have permission to access the resource. | [paymentapiError](#paymentapierror) |
| 404 | Returned when the resource does not exist. | [paymentapiError](#paymentapierror) |
| 500 | Returned when the internal server error occur. | [paymentapiError](#paymentapierror) |

#### DELETE
##### Summary:

Delete payment

##### Description:

Deletes payment by given ID.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [paymentapiPaymentRequest](#paymentapipaymentrequest) |
| 403 | Returned when the user does not have permission to access the resource. | [paymentapiError](#paymentapierror) |
| 404 | Returned when the resource does not exist. | [paymentapiError](#paymentapierror) |
| 500 | Returned when the internal server error occur. | [paymentapiError](#paymentapierror) |

#### PATCH
##### Summary:

Updates payment

##### Description:

Updates payment based on sent specification.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path |  | Yes | string |
| body | body |  | Yes | [paymentapiPaymentAttributes](#paymentapipaymentattributes) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | A successful response. | [paymentapiPaymentServiceResponse](#paymentapipaymentserviceresponse) |
| 403 | Returned when the user does not have permission to access the resource. | [paymentapiError](#paymentapierror) |
| 404 | Returned when the resource does not exist. | [paymentapiError](#paymentapierror) |
| 500 | Returned when the internal server error occur. | [paymentapiError](#paymentapierror) |

### Models


#### ChargersInformationSenderCharges

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| amount | string |  | No |
| currency | string |  | No |

#### paymentapiBeneficiaryParty

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| account_name | string |  | No |
| account_number | string |  | No |
| account_number_code | string |  | No |
| account_type | integer |  | No |
| address | string |  | No |
| bank_id | string |  | No |
| bank_id_code | string |  | No |
| name | string |  | No |

#### paymentapiChargersInformation

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| bearer_code | string |  | No |
| sender_charges | [ [ChargersInformationSenderCharges](#chargersinformationsendercharges) ] |  | No |
| receiver_charges_amount | string |  | No |
| receiver_charges_currency | string |  | No |

#### paymentapiDebtorParty

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| account_name | string |  | No |
| account_number | string |  | No |
| account_number_code | string |  | No |
| address | string |  | No |
| bank_id | string |  | No |
| bank_id_code | string |  | No |
| name | string |  | No |

#### paymentapiError

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| error | string |  | No |
| code | integer |  | No |

#### paymentapiFx

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| contract_reference | string |  | No |
| exchange_rate | string |  | No |
| original_amount | string |  | No |
| original_currency | string |  | No |

#### paymentapiPayment

Payment model

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| type | string |  | No |
| id | string |  | No |
| version | string |  | No |
| organisation_id | string |  | Yes |
| attributes | [paymentapiPaymentAttributes](#paymentapipaymentattributes) |  | No |

#### paymentapiPaymentAttributes

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| amount | string |  | No |
| currency | string |  | No |
| numeric_reference | string |  | No |
| payment_id | string |  | No |
| payment_purpose | string |  | No |
| payment_scheme | string |  | No |
| payment_type | string |  | No |
| processing_date | string |  | No |
| reference | string |  | No |
| scheme_payment_sub_type | string |  | No |
| scheme_payment_type | string |  | No |
| end_to_end_reference | string |  | No |
| beneficiary_party | [paymentapiBeneficiaryParty](#paymentapibeneficiaryparty) |  | No |
| charges_information | [paymentapiChargersInformation](#paymentapichargersinformation) |  | No |
| fx | [paymentapiFx](#paymentapifx) |  | No |
| debtor_party | [paymentapiDebtorParty](#paymentapidebtorparty) |  | No |
| sponsor_party | [paymentapiSponsorParty](#paymentapisponsorparty) |  | No |

#### paymentapiPaymentRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | No |

#### paymentapiPaymentServiceResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data | [ [paymentapiPayment](#paymentapipayment) ] |  | No |

#### paymentapiSponsorParty

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| account_number | string |  | No |
| bank_id | string |  | No |
| bank_id_code | string |  | No |