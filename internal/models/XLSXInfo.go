package models

import (
	"bytes"
)

type XLSXFileInfo struct {
  InputFile *bytes.Buffer;
  SheetName string;
}
