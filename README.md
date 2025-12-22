# Env2JS

Generates JavaScript file based on `PAPERMERGE__` values in environment variables


Here is the js template file:

```
window.__PAPERMERGE_RUNTIME_CONFIG__ = {
  ocr__lang_codes: "{{ .PAPERMERGE__OCR__LANG_CODES }}",
  ocr__default_lang_code: "{{ .PAPERMERGE__OCR__DEFAULT_LANG_CODE }}",
  ocr__automatic: {{ .PAPERMERGE__OCR__AUTOMATIC }},
  auth_type: "oidc",
  oidc_logout_url: "{{ .PAPERMERGE__AUTH__OIDC_LOGOUT_URL }}",
  post_logout_redirect_uri: "{{ .PAPERMERGE__AUTH__POST_LOGOUT_REDIRECT_URI }}"
}
```

Usage:

```
  $ env2js -f core.js.templ > papermerge_runtime.js
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `PAPERMERGE__OCR__LANG_CODES` | Available OCR language codes |
| `PAPERMERGE__OCR__DEFAULT_LANG_CODE` | Default OCR language code |
| `PAPERMERGE__OCR__AUTOMATIC` | Enable automatic OCR (true/false) |
| `PAPERMERGE__AUTH__OIDC_LOGOUT_URL` | OIDC provider logout endpoint URL |
| `PAPERMERGE__AUTH__OIDC_CLIENT_ID`  | OIDC Client ID |
| `PAPERMERGE__AUTH__POST_LOGOUT_REDIRECT_URI` | URL to redirect to after OIDC logout |

