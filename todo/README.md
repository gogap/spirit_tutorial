Todo Component
==============

```bash
spirit-tool run -c spirit.conf -s source.json -v 5
```

#### Create Task

 
```json
{
    "title": "task 1",
    "description": "This Is My First Task"
}
```

```bash
curl -X POST -H "X-API: api.task.new" -H "Content-Type: application/json"  -d '{
    "title": "task 1",
    "description": "This Is My First Task"
}' 'http://127.0.0.1:8080/v1'
```

```bash
{"code":0,"message":"","result":"Vk85tnHnFLdU1qVN"}
```

#### Get Task

 
```json
{
    "id": "Vk85tnHnFLdU1qVN",
}
```

```bash
curl -X POST -H "X-API: api.task.get" -H "Content-Type: application/json"  -d '{
    "id": "Vk85tnHnFLdU1qVN"
}' 'http://127.0.0.1:8080/v1'
```

```bash
{
    "code": 0,
    "message": "",
    "result": {
        "id": "Vk85tnHnFLdU1qVN",
        "is_finished": false,
        "title": "task 1",
        "description": "This Is My First Task",
        "create_time": "2015-11-20T23:18:14.281447693+08:00",
        "update_time": "2015-11-20T23:18:14.281447715+08:00",
        "finished_time": "0001-01-01T00:00:00Z"
    }
}
```