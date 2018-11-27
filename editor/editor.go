package editor

import (
	"os"
	"os/exec"
	"io"
	"io/ioutil"
)

func Call(editor string, target []byte) ([]byte, error) {
	return callEditor(editor, target)
}

func callEditor(editor string, target []byte) ([]byte, error) {
	f, err := ioutil.TempFile(os.TempDir(), "editor-go-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(f.Name())
	tfile := f.Name()
	if err := f.Close(); err != nil {
		return nil, err
	}

	if err := writeFile(tfile, &target); err != nil {
		return nil, err
	}
	if err := execEditor(tfile, editor); err != nil {
		return nil, err
	}

	var ret []byte
	if err := readFile(tfile, &ret) ;err != nil {
		return nil, err
	}
	return ret, nil
}

func execEditor(path string, editor string) error {

	cmd := exec.Command(editor, path)
        cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func writeFile(path string, tag *[]byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(*tag)
	return nil
}

func readFile(path string, ret *[]byte) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 200)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		*ret = append(*ret, buf[:n]...)
	}
	return nil
}
