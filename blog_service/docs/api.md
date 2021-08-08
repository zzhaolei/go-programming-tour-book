### 获取token
```shell
curl -X POST \
  'http://127.0.0.1:8000/auth' \
  -H 'Content-Type: application/json' \
  -d '{"app_key": "creator", "app_secret": "creator_secret"}'
```
