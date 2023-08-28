package fecs

// componentMask is a mask indicating which component types are currently
// assigned to the containing target.
type componentMask struct {
	masks [4]uint64
}

func (m *componentMask) Set(t ComponentType) {
	switch true {
	case t > 192:
		m.masks[3] |= 1 << (t - 193)
	case t > 128:
		m.masks[2] |= 1 << (t - 129)
	case t > 64:
		m.masks[1] |= 1 << (t - 65)
	default:
		m.masks[0] |= 1 << (t - 1)
	}
}

func (m *componentMask) Has(t ComponentType) bool {
	switch true {
	case t > 192:
		return m.masks[3]&1<<(t-193) > 0
	case t > 128:
		return m.masks[2]&1<<(t-129) > 0
	case t > 64:
		return m.masks[1]&1<<(t-65) > 0
	default:
		return m.masks[0]&1<<(t-1) > 0
	}
}

func (m *componentMask) Unset(t ComponentType) {
	switch true {
	case t > 192:
		m.masks[3] &= ^(1 << (t - 193))
	case t > 128:
		m.masks[2] &= ^(1 << (t - 129))
	case t > 64:
		m.masks[1] &= ^(1 << (t - 65))
	default:
		m.masks[0] &= ^(1 << (t - 1))
	}
}

func (m *componentMask) Reset() {
	m.masks = [4]uint64{}
}
