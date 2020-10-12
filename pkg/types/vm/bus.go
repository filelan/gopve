package vm

type Bus string

const (
	BusIDE    Bus = "ide"
	BusSATA   Bus = "sata"
	BusSCSI   Bus = "scsi"
	BusVirtIO Bus = "virtio"
)

func (obj Bus) String() string {
	return string(obj)
}
