package gwell

// DiscoveryState returns the reusable bootstrap state for future sessions.
// It can be cached by callers to avoid repeating server discovery and full
// device enumeration on subsequent connections.
func (s *Session) DiscoveryState() *DiscoveryResult {
	if s == nil {
		return nil
	}
	devices := make([]DeviceInfo, len(s.devices))
	copy(devices, s.devices)
	return &DiscoveryResult{
		Devices:    devices,
		ServerAddr: s.serverAddr,
	}
}
