package cNT24rG4F

// XiaoV2NodeMode 桥接原内联读取 0x50A2AE0 与 0x50A2AF0，保留精确比较。
func XiaoV2NodeMode() bool {
	return nodeConfigurationAt50A2AE0.EtRWBSEhLQc == "xiaov2board" && nodeConfigurationAt50A2AE0.AG_mSFklcvj == "v2node"
}
