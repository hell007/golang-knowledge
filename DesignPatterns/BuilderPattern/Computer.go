package BuilderPattern

type Computer struct {
	CPU      string
	Memory   string
	HardDisk string
}


func (c *Computer) SetCPU(cpu string) {
	c.CPU = cpu
}

func (c *Computer) GetCPU() string {
	return c.CPU
}

func (c *Computer) SetMemory(memory string) {
	c.Memory = memory
}

func (c *Computer) GetMemory() string {
	return c.Memory
}

func (c *Computer) SetHardDisk(hardDisk string) {
	c.HardDisk = hardDisk
}

func (c *Computer) GetHardDisk() string {
	return c.HardDisk
}