# XLSX fix API

A item manegemant api.

### Base endpoint:

`/xlsx`

### Request exemple:

```
curl -X POST \                                                                            [7]
  -F "file=@./test.xlsx;type=application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" \
  -F "sheetName=test" \
  http://localhost:8080/xlsx
```

### Response (200):

`xsls file`
