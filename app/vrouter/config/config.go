package config

type Config struct {
	ID       int64
	System   System
	Security Security
}

type System struct {
	Sysctl map[string]string
}

type Security struct {
	AddressBook AddressBook
}

type AddressBook struct {
	// Address    map[string]Address
	AddressSet map[string]AddressSet
}

type AddressType uint8

const AddressTypeIPv4 AddressType = 4
const AddressTypeIPv6 AddressType = 6

// type Address struct {
// 	Type    AddressType
// 	Address string
// }

type AddressSet struct {
	Type    AddressType
	Address string
}
