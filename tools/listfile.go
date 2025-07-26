package tools

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pararang/code-editing-agent/apis"
)

var ListFilesDefinition = apis.ToolDefinition{
	Name:        "list_files",
	Description: "List files and directories at a given path. If no path is provided, lists files in the current directory.",
	InputSchema: ListFilesInputSchema,
	Function:    ListFiles,
}

type ListFilesInput struct {
	Path string `json:"path,omitempty" jsonschema_description:"Optional relative path to list files from. Defaults to current directory if not provided."`
}

var ListFilesInputSchema = GenerateSchema[ListFilesInput]()

func ListFiles(input json.RawMessage) (string, error) {
	var listFilesInput ListFilesInput
	if err := json.Unmarshal(input, &listFilesInput); err != nil {
		return "", err
	}

	dir := "."
	if listFilesInput.Path != "" {
		dir = listFilesInput.Path
	}

	dirInfo, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return "", &os.PathError{Op: "stat", Path: dir, Err: os.ErrNotExist}
		}
		return "", err
	}

	if !dirInfo.IsDir() {
		return "", &os.PathError{Op: "stat", Path: dir, Err: os.ErrInvalid}
	}

	files := []string{}
	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// TODO: For a robust function, log the error and skip (?)
			return err
		}

		if path == dir {
			return nil
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		fmt.Println("Found file:", relPath)

		if d.IsDir() {
			files = append(files, relPath+"/")
		} else {
			files = append(files, relPath)
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	result, err := json.Marshal(files)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
