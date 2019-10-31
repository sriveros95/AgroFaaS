package unit

// Datarate represents a unit of data rate (in bits per second, bit/s)
type Datarate Unit

// ...
const (
	// base 10 (SI prefixes)
	BitPerSecond      Datarate = 1e0
	KilobitPerSecond           = BitPerSecond * 1e3
	MegabitPerSecond           = BitPerSecond * 1e6
	GigabitPerSecond           = BitPerSecond * 1e9
	TerabitPerSecond           = BitPerSecond * 1e12
	PetabitPerSecond           = BitPerSecond * 1e15
	ExabitPerSecond            = BitPerSecond * 1e18
	ZettabitPerSecond          = BitPerSecond * 1e21
	YottabitPerSecond          = BitPerSecond * 1e24

	BytePerSecond      = BitPerSecond * 8
	KilobytePerSecond  = BytePerSecond * 1e3
	MegabytePerSecond  = BytePerSecond * 1e6
	GigabytePerSecond  = BytePerSecond * 1e9
	TerabytePerSecond  = BytePerSecond * 1e12
	PetabytePerSecond  = BytePerSecond * 1e15
	ExabytePerSecond   = BytePerSecond * 1e18
	ZettabytePerSecond = BytePerSecond * 1e21
	YottabytePerSecond = BytePerSecond * 1e24

	// base 2 (IEC prefixes)
	KibibitPerSecond = BitPerSecond * 1024
	MebibitPerSecond = KibibitPerSecond * 1024
	GibibitPerSecond = MebibitPerSecond * 1024
	TebibitPerSecond = GibibitPerSecond * 1024
	PebibitPerSecond = TebibitPerSecond * 1024
	ExbibitPerSecond = PebibitPerSecond * 1024
	ZebibitPerSecond = ExbibitPerSecond * 1024
	YobibitPerSecond = ZebibitPerSecond * 1024

	KibibytePerSecond = BytePerSecond * 1024
	MebibytePerSecond = KibibytePerSecond * 1024
	GibibytePerSecond = MebibytePerSecond * 1024
	TebibytePerSecond = GibibytePerSecond * 1024
	PebibytePerSecond = TebibytePerSecond * 1024
	ExbibytePerSecond = PebibytePerSecond * 1024
	ZebibytePerSecond = ExbibytePerSecond * 1024
	YobibytePerSecond = ZebibytePerSecond * 1024
)

// BitsPerSecond returns the data rate in bit/s
func (b Datarate) BitsPerSecond() float64 {
	return float64(b)
}

// KilobitsPerSecond returns the data rate in kbit/s
func (b Datarate) KilobitsPerSecond() float64 {
	return float64(b / KilobitPerSecond)
}

// MegabitsPerSecond returns the data rate in Mbit/s
func (b Datarate) MegabitsPerSecond() float64 {
	return float64(b / MegabitPerSecond)
}

// GigabitsPerSecond returns the data rate in Gbit/s
func (b Datarate) GigabitsPerSecond() float64 {
	return float64(b / GigabitPerSecond)
}

// TerabitsPerSecond returns the data rate in Tbit/s
func (b Datarate) TerabitsPerSecond() float64 {
	return float64(b / TerabitPerSecond)
}

// PetabitsPerSecond returns the data rate in Pbit/s
func (b Datarate) PetabitsPerSecond() float64 {
	return float64(b / PetabitPerSecond)
}

// ExabitsPerSecond returns the data rate in Ebit/s
func (b Datarate) ExabitsPerSecond() float64 {
	return float64(b / ExabitPerSecond)
}

// ZettabitsPerSecond returns the data rate in Zbit/s
func (b Datarate) ZettabitsPerSecond() float64 {
	return float64(b / ZettabitPerSecond)
}

// YottabitsPerSecond returns the data rate in Ybit/s
func (b Datarate) YottabitsPerSecond() float64 {
	return float64(b / YottabitPerSecond)
}

// BytesPerSecond returns the data rate in B/s
func (b Datarate) BytesPerSecond() float64 {
	return float64(b / BytePerSecond)
}

// KilobytesPerSecond returns the data rate in kB/s
func (b Datarate) KilobytesPerSecond() float64 {
	return float64(b / KilobytePerSecond)
}

// MegabytesPerSecond returns the data rate in MB/s
func (b Datarate) MegabytesPerSecond() float64 {
	return float64(b / MegabytePerSecond)
}

// GigabytesPerSecond returns the data rate in GB/s
func (b Datarate) GigabytesPerSecond() float64 {
	return float64(b / GigabytePerSecond)
}

// TerabytesPerSecond returns the data rate in TB/s
func (b Datarate) TerabytesPerSecond() float64 {
	return float64(b / TerabytePerSecond)
}

// PetabytesPerSecond returns the data rate in PB/s
func (b Datarate) PetabytesPerSecond() float64 {
	return float64(b / PetabytePerSecond)
}

// ExabytesPerSecond returns the data rate in EB/s
func (b Datarate) ExabytesPerSecond() float64 {
	return float64(b / ExabytePerSecond)
}

// ZettabytesPerSecond returns the data rate in ZB/s
func (b Datarate) ZettabytesPerSecond() float64 {
	return float64(b / ZettabytePerSecond)
}

// YottabytesPerSecond returns the data rate in YB/s
func (b Datarate) YottabytesPerSecond() float64 {
	return float64(b / YottabytePerSecond)
}

// KibibitsPerSecond returns the data rate in Kibit/s
func (b Datarate) KibibitsPerSecond() float64 {
	return float64(b / KibibitPerSecond)
}

// MebibitsPerSecond returns the data rate in Mibit/s
func (b Datarate) MebibitsPerSecond() float64 {
	return float64(b / MebibitPerSecond)
}

// GibibitsPerSecond returns the data rate in Gibit/s
func (b Datarate) GibibitsPerSecond() float64 {
	return float64(b / GibibitPerSecond)
}

// TebibitsPerSecond returns the data rate in Tibit/s
func (b Datarate) TebibitsPerSecond() float64 {
	return float64(b / TebibitPerSecond)
}

// PebibitsPerSecond returns the data rate in Pibit/s
func (b Datarate) PebibitsPerSecond() float64 {
	return float64(b / PebibitPerSecond)
}

// ExbibitsPerSecond returns the data rate in Eibit/s
func (b Datarate) ExbibitsPerSecond() float64 {
	return float64(b / ExbibitPerSecond)
}

// ZebibitsPerSecond returns the data rate in Zibit/s
func (b Datarate) ZebibitsPerSecond() float64 {
	return float64(b / ZebibitPerSecond)
}

// YobibitsPerSecond returns the data rate in Yibit/s
func (b Datarate) YobibitsPerSecond() float64 {
	return float64(b / YobibitPerSecond)
}

// KibibytesPerSecond returns the data rate in KiB/s
func (b Datarate) KibibytesPerSecond() float64 {
	return float64(b / KibibytePerSecond)
}

// MebibytesPerSecond returns the data rate in MiB/s
func (b Datarate) MebibytesPerSecond() float64 {
	return float64(b / MebibytePerSecond)
}

// GibibytesPerSecond returns the data rate in GiB/s
func (b Datarate) GibibytesPerSecond() float64 {
	return float64(b / GibibytePerSecond)
}

// TebibytesPerSecond returns the data rate in TiB/s
func (b Datarate) TebibytesPerSecond() float64 {
	return float64(b / TebibytePerSecond)
}

// PebibytesPerSecond returns the data rate in PiB/s
func (b Datarate) PebibytesPerSecond() float64 {
	return float64(b / PebibytePerSecond)
}

// ExbibytesPerSecond returns the data rate in EiB/s
func (b Datarate) ExbibytesPerSecond() float64 {
	return float64(b / ExbibytePerSecond)
}

// ZebibytesPerSecond returns the data rate in ZiB/s
func (b Datarate) ZebibytesPerSecond() float64 {
	return float64(b / ZebibytePerSecond)
}

// YobibytesPerSecond returns the data rate in YiB/s
func (b Datarate) YobibytesPerSecond() float64 {
	return float64(b / YobibytePerSecond)
}
