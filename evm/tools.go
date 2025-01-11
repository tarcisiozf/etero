package evm

import "reflect"

func assemble(code ...any) []byte {
	out := make([]byte, 0)

	for _, c := range code {
		v := reflect.ValueOf(c)
		switch v.Kind() {
		case reflect.Int:
			out = append(out, byte(v.Int()))
		case reflect.Ptr:
			ix := c.(*Instruction)
			out = append(out, ix.opcode)
		default:
			panic("unsupported type")
		}
	}

	return out
}
