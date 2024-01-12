package provider

import "reflect"

// Name returns the name of the provider.
func Name(provider Provider) string {
	return reflect.TypeOf(provider).Elem().Name()
}

// Elem returns the name of the element type of the provider.
func Elem(provider Provider) string {
	return reflect.TypeOf(provider).Elem().String()
}
