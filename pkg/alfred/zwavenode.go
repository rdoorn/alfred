package alfred

type NodeInterface interface {
	Sha() string // uniq ID to identify node with
}

type ZwaveNode struct {
	ID           byte
	BasicType    byte
	GenericType  byte
	SpecificType byte
}

func NewZwaveNode(id byte) *ZwaveNode {
	return &ZwaveNode{
		ID: id,
	}
}

func SetType() {

}
