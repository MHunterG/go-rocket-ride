// Code generated by sqlc. DO NOT EDIT.
// source: audit_record.sql

package sqlc

import (
	"context"
	"encoding/json"
	"time"
)

const createAuditRecord = `-- name: CreateAuditRecord :one
INSERT INTO audit_records(
	action,
	created_at,
	data,
	origin_ip,
	resource_id,
	resource_type,
	user_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, action, created_at, data, origin_ip, resource_id, resource_type, user_id
`

type CreateAuditRecordParams struct {
	Action       string
	CreatedAt    time.Time
	Data         json.RawMessage
	OriginIp     string
	ResourceID   int64
	ResourceType string
	UserID       int64
}

func (q *Queries) CreateAuditRecord(ctx context.Context, arg CreateAuditRecordParams) (AuditRecord, error) {
	row := q.db.QueryRowContext(ctx, createAuditRecord,
		arg.Action,
		arg.CreatedAt,
		arg.Data,
		arg.OriginIp,
		arg.ResourceID,
		arg.ResourceType,
		arg.UserID,
	)
	var i AuditRecord
	err := row.Scan(
		&i.ID,
		&i.Action,
		&i.CreatedAt,
		&i.Data,
		&i.OriginIp,
		&i.ResourceID,
		&i.ResourceType,
		&i.UserID,
	)
	return i, err
}
