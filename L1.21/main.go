package main

import "fmt"

type SOSBrigade interface {
	ServeTea()
}

type Mikuru struct{}

func (t *Mikuru) ClassifiedInformation() string {
	return "Секретные сведения:)"
}

type MikuruAdapter struct {
	*Mikuru
}

func (a *MikuruAdapter) ServeTea() {
	secret := a.ClassifiedInformation()
	fmt.Printf("Микуру подает чай и испуганно шепчет: «%s»\n", secret)
}

func main() {
	mikuruFromFuture := &Mikuru{}

	adapter := &MikuruAdapter{
		mikuruFromFuture,
	}

	adapter.ServeTea()
}
