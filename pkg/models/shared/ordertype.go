// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// OrderType - The type of order.
type OrderType string

const (
	OrderTypeDrink      OrderType = "drink"
	OrderTypeIngredient OrderType = "ingredient"
)

func (e OrderType) ToPointer() *OrderType {
	return &e
}

func (e *OrderType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "drink":
		fallthrough
	case "ingredient":
		*e = OrderType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrderType: %v", v)
	}
}
