
const(
	I2C_SLAVE_FORCE=1798    //标准i2c接口提供的设置i2c从机地址的命令号
	I2C_TIMEOUT=1794        //标准i2c接口提供的设置i2c总线操作超时命令号
	I2C_RETRIES=1793        //标准i2c接口提供的设置i2c总线重试次数的命令号
)
type FM24xxFram struct{
	Fd 		int //i2c设备接口文件句柄
	Bus 	int //i2c总线编号 从0开始
	Addr 	int //i2c从机地址
	Timeout int //超时时间
	Retry 	int  //重试次数
}
func NewFM24cxxFram(bus,addr,timeout,retry int)*FM24xxFram{
	return &FM24xxFram{
		Fd:-1,
		Bus:bus,
		Addr:addr,
		Timeout:timeout,
		Retry,retry,
	}
}
//返回尺寸大小,以字节为单位 8192
func (f *FM24xxFram)Size()int  {
	return 8192
}
//打开铁电存储器总线设备
func (f *FM24xxFram)Open()error{
	fd,err:=syscall.Open(fmt.Sprintf("/dev/i2c-%d",f.Bus),os.O_RDWR,0777)
	if err!=nil{
		return err
	}
	f.Fd = fd
	syscall.Syscall(syscall.SYS_IOCTL,uintptr(fd),I2C_SLAVE_FORCE,uintptr(f.Addr))
	syscall.Syscall(syscall.SYS_IOCTL,uintptr(fd),I2C_TIMEOUT,f.Timeout)
	syscall.Syscall(syscall.SYS_IOCTL,uintptr(fd),I2C_RETRIES,f.Retry)
	return nil
}
//关闭设备
func (f *FM24xxFram)Close(){
	if f.Fd != -1{
		syscall.Close(f.Fd)
	}
	f.Fd = -1
}

/*
往铁电某个地址写入数据
addr: 铁电的存储地址 从0 - 8K
p ：需要写入的数据内容
 */
func (f *FM24xxFram)Write(addr int,p []byte)(n int, err error){
	writeBuf :=make([]byte,2)
	writeBuf[0]=byte((addr>>8)&0xff)
	writeBuf[1]=byte(addr&0xff)
	if len(p) > 0{
		writeBuf=append(writeBuf,p...)

	}

	return syscall.Write(f.Fd,writeBuf)

	//return f.File.Write(writeBuf)
}

/**
从铁电读取某个地址的数据
addr: 铁电的存储地址
p: 存放读取的数据
 */
func (f *FM24xxFram)Read(addr int,p []byte)(n int, err error){

	//先定位到存储器的某个地址.
	if n,err=f.Write(addr,[]byte{});err!=nil{
		return
	}
	return syscall.Read(f.Fd,p)
}
