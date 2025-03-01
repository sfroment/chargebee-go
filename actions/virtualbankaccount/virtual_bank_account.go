package virtualbankaccount

import (
	"fmt"
	"github.com/sfroment/chargebee-go"
	"github.com/sfroment/chargebee-go/models/virtualbankaccount"
	"net/url"
)

func CreateUsingPermanentToken(params *virtualbankaccount.CreateUsingPermanentTokenRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts/create_using_permanent_token"), params)
}
func Create(params *virtualbankaccount.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts"), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/virtual_bank_accounts/%v", url.PathEscape(id)), nil)
}
func List(params *virtualbankaccount.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/virtual_bank_accounts"), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts/%v/delete", url.PathEscape(id)), nil)
}
func DeleteLocal(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/virtual_bank_accounts/%v/delete_local", url.PathEscape(id)), nil)
}
