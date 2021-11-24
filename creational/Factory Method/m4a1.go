// Concrete product
package main

type m4a1 struct {
	gun
}

func newM4a1() iGun {
	return &m4a1{
		gun: gun{
			name:  "M4A1 gun",
			power: 1,
		},
	}
}
