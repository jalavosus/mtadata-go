package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/jalavosus/mtadata/internal/utils"
)

func readOutputJson[T any](filename string) (parsed []T, err error) {
	fp, fpErr := buildParsedFilePath(filename)
	if fpErr != nil {
		return nil, fpErr
	}

	f, err := utils.OpenFileRead(fp)
	if err != nil {
		return nil, err
	}

	defer func() { _ = f.Close() }()

	dataBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(dataBytes, &parsed); err != nil {
		return nil, err
	}

	return parsed, nil
}

func writeOutputJson(data any, filename string) error {
	fp, fpErr := buildParsedFilePath(filename)
	if fpErr != nil {
		return fpErr
	}

	f, err := utils.OpenFileWrite(fp)
	if err != nil {
		return err
	}

	defer func() { _ = f.Close() }()

	if err = utils.ClearFile(f); err != nil {
		return err
	}

	marshalled, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Write(marshalled)
	if err != nil {
		return err
	}

	return nil
}

func buildParsedFilePath(filename string) (string, error) {
	joined := filepath.Join("./", "data", "parsed", filename)
	fp, fpErr := filepath.Abs(joined)

	return fp, fpErr
}
