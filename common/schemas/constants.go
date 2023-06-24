package schemas

const (
	StatusQuotationCreated          = "QUOTATION_CREATED"
	StatusTransactionInitiated      = "TRANSACTION_INITIATED"
	StatusTransactionCancelled      = "TRANSACTION_CANCELLED"
	StatusTransactionExpired        = "TRANSACTION_EXPIRED"
	StatusTransactionPaid           = "TRANSACTION_PAID"
	StatusDataVerificationNeeded    = "DATA_VERIFICATION_NEEDED"
	StatusDataVerificationSubmitted = "DATA_VERIFICATION_SUBMITTED"
	StatusDataVerificationInProcess = "DATA_VERIFICATION_IN_PROCESS"
	StatusDataVerified              = "DATA_VERIFIED"
	StatusIssuingPolicy             = "ISSUING_POLICY"
	StatusPolicyIssued              = "POLICY_ISSUED"
	StatusPolicyRejected            = "POLICY_REJECTED"
	StatusPolicyActive              = "POLICY_ACTIVE"
	StatusPolicyRejectedQoala       = "POLICY_REJECTED_QOALA"
	StatusPolicyRejectedInsurance   = "POLICY_REJECTED_INSURANCE"
	StatusPolicyResubmitDocument    = "POLICY_RESUBMIT_DOCUMENT"
	StatusPolicyWaitingPayment      = "POLICY_WAITING_PAYMENT"
	StatusPolicyPaid                = "POLICY_PAID"
	StatusPolicyActivatedQoala      = "POLICY_ACTIVATED_QOALA"
	StatusPolicyActivatedInsurance  = "POLICY_ACTIVATED_INSURANCE"
	StatusPolicyEndorsed            = "POLICY_ENDORSED"
	StatusPolicyExpired             = "POLICY_EXPIRED"
	StatusPolicyClaimed             = "POLICY_CLAIMED"
	StatusPolicyClosed              = "POLICY_CLOSED"
	StatusPolicyRenewed             = "POLICY_RENEWED"
	StatusPolicyCancelled           = "POLICY_CANCELLED"
	StatusRefund                    = "REFUND"
	StatusTransactionForTesting     = "TRANSACTION_FOR_TESTING"

	StatusClaimInitiated                        = "CLAIM_INITIATE"
	StatusClaimApproved                         = "CLAIM_STATUS_APPROVED"
	StatusClaimInsuranceApprove                 = "INSURANCE_CLAIM_APPROVE"
	StatusClaimInsuranceRejected                = "INSURANCE_CLAIM_REJECT"
	StatusClaimInsuranceClaimPaid               = "INSURANCE_CLAIM_PAID"
	StatusClaimInsuranceAskDetail               = "INSURANCE_ASK_DETAIL"
	StatusClaimInsuranceAskResubmitCost         = "INSURANCE_CLAIM_ASK_RESUBMIT_COST"
	StatusClaimInsuranceClaimReview             = "INSURANCE_CLAIM_REVIEW"
	StatusClaimInsuranceResubmitDocReqInsurance = "INSURANCE_CLAIM_RESUBMIT_DOCUMENT_REQ_INSURANCE"
	StatusClaimQoalaApproved                    = "QOALA_CLAIM_APPROVE"
	StatusClaimQoalaRejected                    = "QOALA_CLAIM_REJECT"
	StatusClaimQoalaAskDetail                   = "QOALA_ASK_DETAIL"
	StatusClaimQoalaResubmitDocument            = "QOALA_CLAIM_RESUBMIT_DOCUMENT"
	StatusClaimQoalaClaimPaid                   = "QOALA_CLAIM_PAID"
	StatusClaimCustomerResubmitDocument         = "CUSTOMER_RESUBMIT_DOCUMENT"
	StatusClaimCustomerResubmitDocReqInsurance  = "CUSTOMER_RESUBMIT_DOCUMENT_REQ_INSURANCE"
	StatusClaimCancelled                        = "CLAIM_CANCELLED"
	StatusClaimExpire                           = "CLAIM_EXPIRE"
	StatusClaimWaitingExcessFeeConfirm          = "CLAIM_WAITING_EXCESS_FEE_CONFIRM"
	StatusClaimResubmitDocumentReqInsurance     = "QOALA_CLAIM_RESUBMIT_DOCUMENT_REQ_INSURANCE"
	StatusClaimResubmitDocumentReqQoala         = "QOALA_CLAIM_RESUBMIT_DOCUMENT_REQ_QOALA"
	StatusClaimResubmitDocumentAdditional       = "QOALA_CLAIM_RESUBMIT_DOCUMENT_ADDITIONAL"
	StatusClaimCustomerAgreeClaimAmount         = "CUSTOMER_AGREE_CLAIM_AMOUNT"
	StatusClaimCustomerRejectClaimAmount        = "CUSTOMER_REJECT_CLAIM_AMOUNT"
	StatusClaimActivityLog                      = "ACTIVITY_LOG"
	StatusClaimComment                          = "COMMENT"

	StateClaimInitiated = "INITIATED"
	StateClaimApproved  = "APPROVED"
	StateClaimRejected  = "REJECTED"
	StateClaimDone      = "DONE"
	StateClaimPending   = "PENDING"
	StateClaimActive    = "ACTIVE"
	StateClaimComplete  = "COMPLETE"

	LABEL_BUTTON_CLAIM_APPROVE        = "Approve"
	LABEL_BUTTON_CLAIM_REJECT         = "Reject"
	LABEL_BUTTON_CLAIM_ASK_DETAIL     = "Ask for Document"
	LABEL_BUTTON_CLAIM_INSURANCE_PAID = "Mark as Paid"
	LABEL_BUTTON_CLAIM_QOALA_PAID     = "Mark as Paid"
	LABEL_BUTTON_CLAIM_NOTES          = "Notes"
	LABEL_BUTTON_CLAIM_AMOUNT         = "Amount"
	LABEL_BUTTON_ADD_LOG              = "Add Log"
	LABEL_BUTTON_COMMENT              = "Comment"
	LABEL_BUTTON_ACTIVITY_LOG         = "Update Status"
	LABEL_BUTTON_SUBMIT_DOCUMENT      = "Submit Document"

	StatusInvoicePending      = "INVOICE_PENDING"
	StatusInvoiceCreated      = "INVOICE_CREATED"
	StatusInvoiceCancelled    = "INVOICE_CANCELLED"
	StatusInvoicePaid         = "INVOICE_PAID"
	StatusInvoiceExpired      = "INVOICE_EXPIRED"
	StatusInvoiceFundTransfer = "INVOICE_FUND_TRANSFERED"

	StateExpired = "EXPIRED"
	StateSuccess = "SUCCESS"
)
