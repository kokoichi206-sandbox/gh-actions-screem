## from server to client

### type 1

status の更新

```json
{
    "type": 1
}
```

## from client to server

``` json
{"protocol":"json","version":1}
```

### start

1229 は runId っぽい

``` json
{
  "arguments": ["42d3ad83-59f8-4e45-8fa8-c22808f1810c", 1229],
  "target": "WatchRunStepsProgressAsync",
  "type": 1
}
```
