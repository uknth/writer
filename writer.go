package writer

import (
	"bufio"
	"errors"
	"os"
	"sync"
	"time"
)

var errCreatingFile = errors.New("Error creating new file")
var errClosingFile = errors.New("Error closing File")
var errRenamingFile = errors.New("Error renaming File")

type RotateWriter struct {
	lock     sync.Mutex
	filename string // should be set to the actual filename
	duration int
	fp       *os.File
}

// NewWriter makes a new RotateWriter. Return nil if error occurs during setup.
func NewWriter(filename string, timeInSec int) (*RotateWriter, error) {
	// Check file before we initialize.
	return new(filename, timeInSec)
}

func new(filename string, duration int) (*RotateWriter, error) {
	w := &RotateWriter{filename: filename, duration: duration}
	err := w.Rotate()
	if err != nil {
		return nil, err
	}
	// Trigger a rotation after every x interval.
	ticker := time.NewTicker(time.Duration(int64(w.duration)) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				w.Rotate()
			}
		}
	}()
	return w, nil
}

// Write satisfies the io.Writer interface.
func (w *RotateWriter) Write(output []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.fp.Write(output)
}

func (w *RotateWriter) Close() (err error) {
	return w.fp.Close()
}

// Perform the actual act of rotating and reopening file.
func (w *RotateWriter) Rotate() (err error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	// Close existing file if open
	if w.fp != nil {
		err = w.fp.Close()
		w.fp = nil
		if err != nil {
			return errClosingFile
		}
	}
	// Rename dest file if it already exists
	_, err = os.Stat(w.filename)
	if err == nil {
		err = os.Rename(w.filename, w.filename+"."+time.Now().Format(time.RFC3339))
		if err != nil {
			return errRenamingFile
		}
	}

	// Create a file.
	w.fp, err = os.Create(w.filename)
	if err != nil {
		return errCreatingFile
	}
	return nil
}

func (w *RotateWriter) read() ([]string, error) {
	var buffer []string

	fp, err := os.Open(w.filename)
	if err != nil {
		return nil, err
	}
	// File is already open we use the scanner on it.
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		text := scanner.Text()
		buffer = append(buffer, text)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return buffer, nil
}
