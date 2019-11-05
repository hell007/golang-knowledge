package BuilderPattern

type Director struct {
	Builder Builder
}

func (d Director) Create(cpu string, memory string, hardDisk string) *Computer {
	return d.Builder.SetCPU(cpu).SetMemory(memory).SetHardDisk(hardDisk).Build()
}