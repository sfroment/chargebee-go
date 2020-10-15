package enum

type EventType string

const (
	EventTypePlanCreated                               EventType = "plan_created"
	EventTypePlanUpdated                               EventType = "plan_updated"
	EventTypePlanDeleted                               EventType = "plan_deleted"
	EventTypeAddonCreated                              EventType = "addon_created"
	EventTypeAddonUpdated                              EventType = "addon_updated"
	EventTypeAddonDeleted                              EventType = "addon_deleted"
	EventTypeCouponCreated                             EventType = "coupon_created"
	EventTypeCouponUpdated                             EventType = "coupon_updated"
	EventTypeCouponDeleted                             EventType = "coupon_deleted"
	EventTypeCouponSetCreated                          EventType = "coupon_set_created"
	EventTypeCouponSetUpdated                          EventType = "coupon_set_updated"
	EventTypeCouponSetDeleted                          EventType = "coupon_set_deleted"
	EventTypeCouponCodesAdded                          EventType = "coupon_codes_added"
	EventTypeCouponCodesDeleted                        EventType = "coupon_codes_deleted"
	EventTypeCouponCodesUpdated                        EventType = "coupon_codes_updated"
	EventTypeCustomerCreated                           EventType = "customer_created"
	EventTypeCustomerChanged                           EventType = "customer_changed"
	EventTypeCustomerDeleted                           EventType = "customer_deleted"
	EventTypeCustomerMovedOut                          EventType = "customer_moved_out"
	EventTypeCustomerMovedIn                           EventType = "customer_moved_in"
	EventTypePromotionalCreditsAdded                   EventType = "promotional_credits_added"
	EventTypePromotionalCreditsDeducted                EventType = "promotional_credits_deducted"
	EventTypeSubscriptionCreated                       EventType = "subscription_created"
	EventTypeSubscriptionStarted                       EventType = "subscription_started"
	EventTypeSubscriptionTrialEndReminder              EventType = "subscription_trial_end_reminder"
	EventTypeSubscriptionActivated                     EventType = "subscription_activated"
	EventTypeSubscriptionChanged                       EventType = "subscription_changed"
	EventTypeMrrUpdated                                EventType = "mrr_updated"
	EventTypeSubscriptionCancellationScheduled         EventType = "subscription_cancellation_scheduled"
	EventTypeSubscriptionCancellationReminder          EventType = "subscription_cancellation_reminder"
	EventTypeSubscriptionCancelled                     EventType = "subscription_cancelled"
	EventTypeSubscriptionReactivated                   EventType = "subscription_reactivated"
	EventTypeSubscriptionRenewed                       EventType = "subscription_renewed"
	EventTypeSubscriptionScheduledCancellationRemoved  EventType = "subscription_scheduled_cancellation_removed"
	EventTypeSubscriptionChangesScheduled              EventType = "subscription_changes_scheduled"
	EventTypeSubscriptionScheduledChangesRemoved       EventType = "subscription_scheduled_changes_removed"
	EventTypeSubscriptionShippingAddressUpdated        EventType = "subscription_shipping_address_updated"
	EventTypeSubscriptionDeleted                       EventType = "subscription_deleted"
	EventTypeSubscriptionPaused                        EventType = "subscription_paused"
	EventTypeSubscriptionPauseScheduled                EventType = "subscription_pause_scheduled"
	EventTypeSubscriptionScheduledPauseRemoved         EventType = "subscription_scheduled_pause_removed"
	EventTypeSubscriptionResumed                       EventType = "subscription_resumed"
	EventTypeSubscriptionResumptionScheduled           EventType = "subscription_resumption_scheduled"
	EventTypeSubscriptionScheduledResumptionRemoved    EventType = "subscription_scheduled_resumption_removed"
	EventTypeSubscriptionAdvanceInvoiceScheduleAdded   EventType = "subscription_advance_invoice_schedule_added"
	EventTypeSubscriptionAdvanceInvoiceScheduleUpdated EventType = "subscription_advance_invoice_schedule_updated"
	EventTypeSubscriptionAdvanceInvoiceScheduleRemoved EventType = "subscription_advance_invoice_schedule_removed"
	EventTypePendingInvoiceCreated                     EventType = "pending_invoice_created"
	EventTypePendingInvoiceUpdated                     EventType = "pending_invoice_updated"
	EventTypeInvoiceGenerated                          EventType = "invoice_generated"
	EventTypeInvoiceUpdated                            EventType = "invoice_updated"
	EventTypeInvoiceDeleted                            EventType = "invoice_deleted"
	EventTypeCreditNoteCreated                         EventType = "credit_note_created"
	EventTypeCreditNoteUpdated                         EventType = "credit_note_updated"
	EventTypeCreditNoteDeleted                         EventType = "credit_note_deleted"
	EventTypeSubscriptionRenewalReminder               EventType = "subscription_renewal_reminder"
	EventTypeTransactionCreated                        EventType = "transaction_created"
	EventTypeTransactionUpdated                        EventType = "transaction_updated"
	EventTypeTransactionDeleted                        EventType = "transaction_deleted"
	EventTypePaymentSucceeded                          EventType = "payment_succeeded"
	EventTypePaymentFailed                             EventType = "payment_failed"
	EventTypePaymentRefunded                           EventType = "payment_refunded"
	EventTypePaymentInitiated                          EventType = "payment_initiated"
	EventTypeRefundInitiated                           EventType = "refund_initiated"
	EventTypeNetdPaymentDueReminder                    EventType = "netd_payment_due_reminder"
	EventTypeAuthorizationSucceeded                    EventType = "authorization_succeeded"
	EventTypeAuthorizationVoided                       EventType = "authorization_voided"
	EventTypeCardAdded                                 EventType = "card_added"
	EventTypeCardUpdated                               EventType = "card_updated"
	EventTypeCardExpiryReminder                        EventType = "card_expiry_reminder"
	EventTypeCardExpired                               EventType = "card_expired"
	EventTypeCardDeleted                               EventType = "card_deleted"
	EventTypePaymentSourceAdded                        EventType = "payment_source_added"
	EventTypePaymentSourceUpdated                      EventType = "payment_source_updated"
	EventTypePaymentSourceDeleted                      EventType = "payment_source_deleted"
	EventTypePaymentSourceExpiring                     EventType = "payment_source_expiring"
	EventTypePaymentSourceExpired                      EventType = "payment_source_expired"
	EventTypeVirtualBankAccountAdded                   EventType = "virtual_bank_account_added"
	EventTypeVirtualBankAccountUpdated                 EventType = "virtual_bank_account_updated"
	EventTypeVirtualBankAccountDeleted                 EventType = "virtual_bank_account_deleted"
	EventTypeTokenCreated                              EventType = "token_created"
	EventTypeTokenConsumed                             EventType = "token_consumed"
	EventTypeTokenExpired                              EventType = "token_expired"
	EventTypeUnbilledChargesCreated                    EventType = "unbilled_charges_created"
	EventTypeUnbilledChargesVoided                     EventType = "unbilled_charges_voided"
	EventTypeUnbilledChargesDeleted                    EventType = "unbilled_charges_deleted"
	EventTypeUnbilledChargesInvoiced                   EventType = "unbilled_charges_invoiced"
	EventTypeOrderCreated                              EventType = "order_created"
	EventTypeOrderUpdated                              EventType = "order_updated"
	EventTypeOrderCancelled                            EventType = "order_cancelled"
	EventTypeOrderDelivered                            EventType = "order_delivered"
	EventTypeOrderReturned                             EventType = "order_returned"
	EventTypeOrderReadyToProcess                       EventType = "order_ready_to_process"
	EventTypeOrderReadyToShip                          EventType = "order_ready_to_ship"
	EventTypeOrderDeleted                              EventType = "order_deleted"
	EventTypeQuoteCreated                              EventType = "quote_created"
	EventTypeQuoteUpdated                              EventType = "quote_updated"
	EventTypeQuoteDeleted                              EventType = "quote_deleted"
	EventTypeGiftScheduled                             EventType = "gift_scheduled"
	EventTypeGiftUnclaimed                             EventType = "gift_unclaimed"
	EventTypeGiftClaimed                               EventType = "gift_claimed"
	EventTypeGiftExpired                               EventType = "gift_expired"
	EventTypeGiftCancelled                             EventType = "gift_cancelled"
	EventTypeGiftUpdated                               EventType = "gift_updated"
	EventTypeHierarchyCreated                          EventType = "hierarchy_created"
	EventTypeHierarchyDeleted                          EventType = "hierarchy_deleted"
	EventTypePaymentIntentCreated                      EventType = "payment_intent_created"
	EventTypePaymentIntentUpdated                      EventType = "payment_intent_updated"
	EventTypeContractTermCreated                       EventType = "contract_term_created"
	EventTypeContractTermRenewed                       EventType = "contract_term_renewed"
	EventTypeContractTermTerminated                    EventType = "contract_term_terminated"
	EventTypeContractTermCompleted                     EventType = "contract_term_completed"
	EventTypeContractTermCancelled                     EventType = "contract_term_cancelled"
)
