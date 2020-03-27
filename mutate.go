package sanity

import (
	"context"
	"encoding/json"
	"reflect"
)

// MutateQueryResult holds the result of a query API call.
type MutateQueryResult struct {
	TransactionID string           `json:"transactionId"`
	Results       *json.RawMessage `json:"results"`
}

func (q *MutateQueryResult) unmarshal(out interface{}) error {
	if q.Results == nil {
		v := reflect.ValueOf(&out)
		if v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
			i := reflect.Indirect(v)
			i.Set(reflect.Zero(i.Type()))
		}
		return nil
	}
	return json.Unmarshal([]byte(*q.Results), out)

}

// Mutate performs a post.
// will receive the result. json.Unmarshal is used to deserialize the JSON; if the type
// supports json.Unmarshaller, then it can override the unmarshalling.
//
// On API failure, this will return an error of type *RequestError.
func (c *Client) Mutate(ctx context.Context, payload []byte, out interface{}) error {

	var resp MutateQueryResult
	if _, err := c.performPOST(ctx, "data/mutate/"+c.dataset, payload, &resp); err != nil {
		return err
	}

	return resp.unmarshal(out)
}
