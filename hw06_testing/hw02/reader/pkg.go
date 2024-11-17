package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/overgoy/home_work_/hw06_testing/hw02/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	return data, nil
}
