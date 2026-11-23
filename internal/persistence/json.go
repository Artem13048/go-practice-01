package persistence

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"example.com/go-practice-01/internal/tasks"
)

func Load(path string) (tasks.State, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return tasks.State{NextID: 1}, nil
		}
		return tasks.State{}, fmt.Errorf("read %s: %w", path, err)
	}

	var state tasks.State
	dec := json.NewDecoder(bytesReader(data))
	dec.DisallowUnknownFields()

	if err := dec.Decode(&state); err != nil {
		return tasks.State{}, fmt.Errorf("decode %s: %w", path, err)
	}

	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		if err == nil {
			return tasks.State{}, fmt.Errorf("decode %s: unexpected second JSON document", path)
		}
		return tasks.State{}, fmt.Errorf("decode %s: trailing data: %w", path, err)
	}

	store := tasks.NewMemoryStore()
	if err := store.Restore(state); err != nil {
		return tasks.State{}, fmt.Errorf("validate %s: %w", path, err)
	}

	return state, nil
}

func Save(path string, state tasks.State) error {
	// Валидация до любой записи.
	store := tasks.NewMemoryStore()
	if err := store.Restore(state); err != nil {
		return fmt.Errorf("validate state: %w", err)
	}

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal state: %w", err)
	}
	data = append(data, '\n')

	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, ".tasks-*.tmp")
	if err != nil {
		return fmt.Errorf("create temp file in %s: %w", dir, err)
	}
	tmpName := tmp.Name()

	defer func() {
		if tmpName != "" {
			_ = os.Remove(tmpName)
		}
	}()

	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("write temp file: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("sync temp file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("close temp file: %w", err)
	}

	if err := os.Rename(tmpName, path); err != nil {
		return fmt.Errorf("rename temp to %s: %w", path, err)
	}
	tmpName = "" // успешно — defer не удаляет

	return nil
}

func bytesReader(data []byte) io.Reader {
	return &byteReader{data: data}
}

type byteReader struct {
	data []byte
	pos  int
}

func (r *byteReader) Read(p []byte) (int, error) {
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}
	n := copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}
