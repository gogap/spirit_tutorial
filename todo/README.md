Todo Component
==============

#### Run Todo Component

- stand-alone api server

```bash
spirit-tool run -c 01.spirit_api_todo.conf -s source.json -v 5 -e SPIRIT_ENV="$PWD/todo.env"
```

- api request to mns queue, and response to api-call-back

> You should have ali-mns account (http://www.aliyun.com/product/mns), similar with amazon sqs, and create following queues: `api-call-back` `todo-task-delete` `todo-task-finish` `todo-task-get` `todo-task-list` `todo-task-new`

```bash
spirit-tool run -c 02.spirit_mns_api.conf -s source.json -v 5 -e SPIRIT_ENV="$PWD/todo.env"
spirit-tool run -c 02.spirit_mns_todo.conf -s source.json -v 5 -e SPIRIT_ENV="$PWD/todo.env"
```

- advance usage of mns queue, the api request will send to mns queue, and it will pass the delivery to next queue.

```bash
spirit-tool run -c 03.spirit_mns_api.conf -s source.json -v 5 -e SPIRIT_ENV="$PWD/todo.env"
spirit-tool run -c 03.spirit_mns_todo.conf -s source.json -v 5 -e SPIRIT_ENV="$PWD/todo.env"
```

### API

#### Create Task

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

#### Finish Task

```bash
curl -X POST -H "X-API: api.task.finish" -H "Content-Type: application/json"  -d '{
    "id": "Vk85tnHnFLdU1qVN"
}' 'http://127.0.0.1:8080/v1'
```

```bash
{
  "code": 0,
  "message": "",
  "result": null
}
```

#### List Task

```bash
curl -X POST -H "X-API: api.task.list" -H "Content-Type: application/json" 'http://127.0.0.1:8080/v1'
```

```bash
{
    "code": 0,
    "message": "",
    "result": [{
        "id": "Vk85tnHnFLdU1qVN",
        "is_finished": false,
        "title": "task 1",
        "description": "This Is My First Task",
        "create_time": "2015-11-20T23:18:14.281447693+08:00",
        "update_time": "2015-11-20T23:18:14.281447715+08:00",
        "finished_time": "0001-01-01T00:00:00Z"
    }
}]
```

#### Delete Task

```bash
curl -X POST -H "X-API: api.task.delete" -H "Content-Type: application/json"  -d '{
    "id": "Vk85tnHnFLdU1qVN"
}' 'http://127.0.0.1:8080/v1'
```

```bash
{
  "code": 0,
  "message": "",
  "result": null
}
```