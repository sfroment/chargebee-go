package usage

import (
	"fmt"
	"github.com/sfroment/chargebee-go"
	"github.com/sfroment/chargebee-go/models/usage"
	"net/url"
)

func Create(id string, params *usage.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/usages", url.PathEscape(id)), params)
}
func Retrieve(id string, params *usage.RetrieveRequestParams) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v/usages", url.PathEscape(id)), params)
}
func Delete(id string, params *usage.DeleteRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/delete_usage", url.PathEscape(id)), params)
}
func List(params *usage.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/usages"), params)
}
func Pdf(params *usage.PdfRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/usages/pdf"), params)
}
