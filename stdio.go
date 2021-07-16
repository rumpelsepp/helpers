package helpers

import "os"

type StdioWrapper struct {
	Stdin  *os.File
	Stdout *os.File
}

func NewStdioWrapper() *StdioWrapper {
	return &StdioWrapper{os.Stdin, os.Stdout}
}

func (w *StdioWrapper) Read(p []byte) (int, error) {
	return w.Stdin.Read(p)
}

func (w *StdioWrapper) Write(p []byte) (int, error) {
	return w.Stdout.Write(p)
}

func (w *StdioWrapper) Close() error {
	if err := w.Stdin.Close(); err != nil {
		return err
	}
	if err := w.Stdout.Close(); err != nil {
		return err
	}
	return nil
}
