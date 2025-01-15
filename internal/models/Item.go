package models

import "mime/multipart"

type XSLFileInfo struct {
  InputFile multipart.File;
  SheetName string;
}
