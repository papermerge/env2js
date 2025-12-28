# Env2JS

Generates JavaScript file based on `PM_` values in environment variables


Here is the js template file:

```
window.__PAPERMERGE_RUNTIME_CONFIG__ = {
  auth_type: "oidc",
  oidc_logout_url: "{{ .PM_OIDC_LOGOUT_URL }}",
  post_logout_redirect_uri: "{{ .PM_POST_LOGOUT_REDIRECT_URI }}"
}
```

Usage:

```
  $ env2js -f core.js.templ > papermerge_runtime.js
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `PM_OIDC_LOGOUT_URL` | OIDC provider logout endpoint URL |
| `PM_OIDC_CLIENT_ID`  | OIDC Client ID |
| `PM_POST_LOGOUT_REDIRECT_URI` | URL to redirect to after OIDC logout |

