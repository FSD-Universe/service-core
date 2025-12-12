// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

// Package config
package config

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"half-nothing.cn/service-core/interfaces/global"
)

func CreateFileWithContent(filePath string, content []byte) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, global.DefaultDirectoryPermission); err != nil {
		return err
	}
	return os.WriteFile(filePath, content, global.DefaultFilePermissions)
}

func ReadOrDownloadFile(filePath, url string) ([]byte, error) {
	if content, err := os.ReadFile(filePath); err == nil {
		return content, nil
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("file read error: %w", err)
	}

	fmt.Printf("%s not found, downloading from %s", filePath, url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	fmt.Printf("Connection established with %s", url)

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error: %w", err)
	}

	fmt.Printf("%s successfully downloaded, %d bytes", filePath, len(content))

	if err := CreateFileWithContent(filePath, content); err != nil {
		return nil, fmt.Errorf("file write error: %w", err)
	}

	return content, nil
}

func CheckPoint(port uint) error {
	if port <= 0 {
		return errors.New("port must be greater than 0")
	}
	if port > 65535 {
		return errors.New("port must be less than 65535")
	}
	return nil
}
