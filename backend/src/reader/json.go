package reader

import (
	"encoding/json"
	"io"
	"marketplace/backend/src/data"
	"os"
)

type Data struct {
	Item []data.Product `json:"products"`
}

func (d *Data) Parse(file *os.File) error {
	reader, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(reader, d)

	if err != nil {
		return err
	}

	return nil
}
