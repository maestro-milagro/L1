package main

import "fmt"

// В каждой реализации адаптера у нас есть ряд ключевых компонентов:
// Target интерфейс который определяет основную структуру метода, который мы хотим адаптировать
// Adaptee - адаптируемый класс, методы которого необходимо вызвать
// в другом месте с использованием нашего интерфейса.
// ConcreteAdapter - является оболочкой для Adaptee (включает его как атрибут)
// и содержит метод удовлетворяющий сигнатуре, которую хотим использовать в Client для вызова.
// И так у нас есть два ноутбука mac и windows, и мы хотим подключится к ним но у мы можем подключиться только с порту квадратной формы
// такой порт есть у mac, но нет у windows, следовательно нам нужен адаптер для windows
// Target в нашем случае это интерфейс компьютор с нужным нам методом insertInSquarePort
type computer interface {
	insertInSquarePort()
}
type client struct {
}

func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}

type mac struct {
}

func (m *mac) insertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}

// Adaptee windows с методом к которому нам нужен адаптер insertInCirclePort
type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}

// И сам класс оболочка адаптер вызывающий адаптируемый метод с помощью нужного нам метода insertInSquarePort
type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}
func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}
	// Таким образом мы посути вызываем нужный клиенту метод insertSquareUsbInComputer у структуры windows через адаптер windowsAdapter
	client.insertSquareUsbInComputer(windowsMachineAdapter)
}
