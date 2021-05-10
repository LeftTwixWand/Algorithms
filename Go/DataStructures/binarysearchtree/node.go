package binarysearchtree

import "time"

type Node struct {
	CarNumber        int
	RegistrationDate time.Time
	Name             string
	Surname          string
	Left             *Node
	Right            *Node
}
