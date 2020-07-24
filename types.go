package main

type subscriber struct {
	name   string
	rate   float64
	active bool
}

type part struct {
	description string
	count       int
}

type student struct {
	name  string
	grade float64
}
