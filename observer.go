package main

import (
	"log"
)

type SubjectInterface interface {
	Register(observer ObserverInterface) bool
	Notify()
	Remove(observer ObserverInterface) bool
}

type ObserverInterface interface {
	Update(args []interface{})
}

type Subject struct {
	observers []ObserverInterface
}

type Observer struct {
	name string
}

func (self *Subject) Register(observer ObserverInterface) bool {
	self.observers = append(self.observers, observer)
	return true
}

func (self *Subject) Remove(observer ObserverInterface) bool {
	return true
}

func (self *Subject) Notify() {
	for index, observer := range self.observers {
		args := []interface{}{index}
		observer.Update(args)
	}
}

func (self *Observer) Update(args []interface{}) {
	log.Println(self.name, args)
}

func main() {
	subect := Subject{}

	observer01 := &Observer{"01"}
	observer02 := &Observer{"02"}

	subect.Register(observer01)
	subect.Register(observer02)

	subect.Notify()
}
