package cNT24rG4F

// DatabaseConfiguration 暴露原 SQL 初始化器读取的全局字段，非单节点副本。
func DatabaseConfiguration() (host string, port int, name, user, password string) {
	c := &nodeConfigurationAt50A2AE0
	return c.ApfXLZLmFXh, c.Mk5cDiQzaj, c.CsBHZVUxTO, c.ZGaSmv, c.ASSy13OjR6aJ
}

// 原初始化器在 sql.Open 成功之后读取连接池参数。
func DatabasePoolConfiguration() (idleMinutes, idleConnections, openConnections int) {
	c := &nodeConfigurationAt50A2AE0
	return c.NawkeUTY_dJT, c.EUexnK81w_, c.MpUqmS0
}
