package Pkg

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	
	"github.com/golang/snappy"
)

func Compact(name, path string) error {
	list, err := ListAll(path)
	if err != nil {
		return err
	}
	
	b64 := base64.StdEncoding.EncodeToString([]byte(list))
	compressed := snappy.Encode(nil, []byte(b64))

	ierr := ioutil.WriteFile(name, []byte(compressed), os.ModePerm)
	if ierr != nil {
		return ierr
	}
	
    return  nil
}