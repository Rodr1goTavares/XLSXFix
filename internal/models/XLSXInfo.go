package models

import "mime/multipart"

type XLSXFileInfo struct {
  InputFile multipart.File;
  SheetName string;
}
