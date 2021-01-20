# user-service
A user-service written in GO which returns a collection of dummy users.

Endpoints | Status | Body
------------ | -------------| -------------
`/health` | `200` | N/A
`/getUserSummary` | `200`, `400`, `500` | `[]`, `[{ id: string, name: string, avatar: string}]`
