package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func StdinReader(data io.Reader) ([]byte, error) {
	var stdin []byte
	r := bufio.NewReader(data)
	var err error
	var line []byte
	for err == nil {
		line, err = r.ReadBytes('\n')
		stdin = append(stdin, line...)
	}
	if err != io.EOF {
		return nil, fmt.Errorf("error reading file: %s", err.Error())
	}
	return stdin, nil
}

func Reader(stdin *os.File) ([]byte, error) {
	stat, _ := stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := StdinReader(os.Stdin)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, fmt.Errorf("no data to be processed")
}
