package redisClient

type mockRedisCfg struct {
	db       int
	hostname string
}

func (m *mockRedisCfg) DB() int {
	return m.db
}
func (m *mockRedisCfg) Hostname() string {
	return m.hostname
}
func (m *mockRedisCfg) Password() string {
	return ""
}
func newMockRedisCfg(hostname string, db int) *mockRedisCfg {
	return &mockRedisCfg{db: db, hostname: hostname}
}
