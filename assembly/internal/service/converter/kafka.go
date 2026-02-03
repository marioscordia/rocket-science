package converter

import (
	"encoding/json"

	"github.com/marioscordia/rocket-science/assembly/internal/model"
)

func DecodeOrderMsg(msg []byte) (*model.Order, error) {
	var order model.Order
	err := json.Unmarshal(msg, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
