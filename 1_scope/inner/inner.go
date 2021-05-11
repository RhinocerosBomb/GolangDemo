package inner

// names starting with lower cases are private and names starting with upper cases are public
type privateStruct struct {
	privVal int
	PubVal int
}

type PublicStruct struct {
	privVal int
	PubVal int
}

func NewPrivate() privateStruct {
	return privateStruct{}
}

func NewPublic() PublicStruct {
	return PublicStruct{
		privVal: 10,
	}
}

// Same as:
//func SetPriv(pub PublicStruct, val int) PublicStruct
func (pub PublicStruct) SetPriv(val int) PublicStruct {
	pub.privVal = val
	return pub
}

