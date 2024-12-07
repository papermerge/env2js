# Env2JS

Generates JavaScript file based on `PAPERMERGE__` values in environment variables


Here is the js template files:

```
window.__PAPERMERGE_RUNTIME_CONFIG__ = {
  ocr__lang_codes: "{{ .PAPERMERGE__OCR__LANG_CODES }}",
  ocr__default_lang_code: "{{ .PAPERMERGE__OCR__DEFAULT_LANG_CODE }}",
  ocr__automatic: {{ .PAPERMERGE__OCR__AUTOMATIC }}
}
```

Usage:

```
  $ env2js -f core.js.templ > papermerge_runtime.js
```

