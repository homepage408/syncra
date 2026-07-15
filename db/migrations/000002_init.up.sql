CREATE INDEX idx_sessions_user_active 
ON SESSIONS (user_id, id) 
WHERE deleted_at IS NULL AND revoked_at IS NULL;