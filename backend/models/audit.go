package models

import (
	"database/sql"
	"time"
)

type AuditLog struct {
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	UserID    *int      `json:"user_id,omitempty"`
	Path      string    `json:"path"`
}

func GetAuditLogs(db *sql.DB, limit int) ([]AuditLog, error) {
	rows, err := db.Query(`
		SELECT timestamp, ip, user_id, path
		FROM audit_log
		ORDER BY timestamp DESC
		LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []AuditLog
	for rows.Next() {
		var log AuditLog
		if err := rows.Scan(&log.Timestamp, &log.IP, &log.UserID, &log.Path); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
