package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Создаем структуру в которой будет храниться вся информация билета
type Ticket struct {
	Company  string // Компания предоставляющая услуги полетов
	Duration int    // Длительность полёта
	TripType string // Тип билета (В один конец/Туда обратно)
	Price    int    // Цена билета
}

// Создаем функцию для вывода информации о всех имеющихся билетах
func PrintTickets(AllTickets []Ticket) {
	//-------------------|123456712345671234567|1123456781|123123456789123|112345678901|
	fmt.Println("№   |       Company       | Duration |   Trip Type   | Price, mln$|")
	for i, data := range AllTickets {
		fmt.Printf("%4d|%21s|%10d|%15s|%12d|\n", i+1, data.Company, data.Duration, data.TripType, data.Price)
	}
}

// Функция добавления билета в слайс
func AddToTickets(AllTick *[]Ticket, comp string, dur int, tt string, pr int) {

	var dat Ticket
	dat.Duration = dur
	dat.TripType = tt
	dat.Price = pr
	dat.Company = comp

	*AllTick = append(*AllTick, dat)
}

func main() {

	// Создаём список всех компаний предоставляющих услуги
	spaceline := map[int]string{
		1: "Virgin Galactics",
		2: "SpaceX",
		3: "Space Adventures",
		4: "NASA",
	}

	// Тип билета (В один конец/Туда обратно)
	ticketType := map[int]string{
		1: "One-way",
		2: "Round-trip",
	}

	// Срез, в котором будет хранится вся информация о билетах
	Alltickets := make([]Ticket, 0, 10)

	rand.Seed(time.Now().UnixNano())
	// Добавим первую запись в срез
	//AddToTickets(&Alltickets, spaceline[1], 50+rand.Intn(12), ticketType[1], 36+rand.Intn(14))

	// Заполнение Alltickets
	kol := rand.Intn(50)
	fmt.Printf("В продаже появилось %d билетов:\n", kol)

	for i := 0; i < kol; i++ {
		tt := 1 + rand.Intn(2)
		if tt == 1 {
			AddToTickets(&Alltickets, spaceline[1+rand.Intn(3)], 50+rand.Intn(12), ticketType[tt], 36+rand.Intn(14))
		} else {
			AddToTickets(&Alltickets, spaceline[1+rand.Intn(3)], 50+rand.Intn(12), ticketType[2], 2*(36+rand.Intn(14)))
		}
	}
	PrintTickets(Alltickets)
}
