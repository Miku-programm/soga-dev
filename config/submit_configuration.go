package config

// SubmitTrafficThresholds reads the original float64 globals 0x50a29d0
// and 0x50a29d8. Their conversion to byte thresholds belongs to the caller.
func SubmitTrafficThresholds() (aliveIP, traffic float64) {
	return configAt50A2900.KI4eIM8J, configAt50A2900.DhXdMvN
}
