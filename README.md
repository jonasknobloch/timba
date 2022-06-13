# Timba

Timba is a dead simple CLI tool to convert [toggl track](https://track.toggl.com) CSV exports into the backup format
of the [Tim Time Tracker](https://tim.neat.software).

**Features reserved for the toggl premium tier (tasks & billing) are not supported.**

The `tim` package can be used to manually create a tim backup file. See `tim/exampe_test.go` for details.

*This project is not affiliated, associated, authorized, endorsed by, or in any way officially connected with
Neat Software Co. or any of its subsidiaries or its affiliates.*

## Usage

```
timba toggl.csv > toggl.json
```

```json
{
  "tasks": {
    "3587bc64-7911-479c-9edd-ec73fc6c5313": {
      "records": [
        {
          "start": 1654602358000,
          "end": 1654605058000
        },
        {
          "start": 1654608310000,
          "end": 1654612810000
        }
      ],
      "id": "3587bc64-7911-479c-9edd-ec73fc6c5313",
      "title": "Formale Baumsprachen",
      "updatedAt": 1655132291,
      "createdAt": 1655132291
    }
  },
  "groups": {
    "6e5afbc7-e068-405a-8105-3135d9393757": {
      "id": "6e5afbc7-e068-405a-8105-3135d9393757",
      "title": "INF-VERT6",
      "updatedAt": 1655132291631,
      "createdAt": 1655132291631
    }
  },
  "nodes": [
    {
      "id": "6e5afbc7-e068-405a-8105-3135d9393757"
    },
    {
      "id": "3587bc64-7911-479c-9edd-ec73fc6c5313",
      "parent": "6e5afbc7-e068-405a-8105-3135d9393757"
    }
  ]
}
```


