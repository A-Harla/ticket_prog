package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const (
	distance = 62100000 // Примерное расстояние от Земли до Марса в км
)

// Структуру в которой будет храниться вся информация билета
type Ticket struct {
	Company  string // Компания предоставляющая услуги полетов
	Duration int    // Длительность полёта
	TripType string // Тип билета (В один конец/Туда обратно)
	Price    int    // Цена билета
}

// Функцию для вывода информации о всех имеющихся билетах
func PrintTickets(AllTickets *[]Ticket) {
	fmt.Println("№   |       Company       | Duration |   Trip Type   | Price, mln$|")
	for i, data := range *AllTickets {
		fmt.Printf("%4d|%21s|%10d|%15s|%12d|\n", i+1, data.Company, data.Duration, data.TripType, data.Price)
	}
}

// Функция добавления билета в слайс
func AddToTickets(AllTick *[]Ticket, comp string, dur int, tt string, pr int) {

	// Собираем все переданные данные в одну структуру, тип которой соответствует ячейке слайса
	var dat Ticket
	dat.Duration = dur
	dat.TripType = tt
	dat.Price = pr
	dat.Company = comp

	// добавляем информацию в слайс
	*AllTick = append(*AllTick, dat)
}

// Вычисление продолжительность полёта
func GetDuration() (dur int) {
	speed := (rand.Intn(15) + 15) * 10000 // Скорость корабля в тысячах км/ч
	dur = int(math.Ceil(float64(distance) / (3.6 * float64(speed))))
	return dur
}

// Сортировка готового списка билетов по увеличению длительности полёта
func SortTickets(AllTick *[]Ticket) {
	sort.Slice(*AllTick, func(i, j int) bool {
		return (*AllTick)[i].Duration < (*AllTick)[j].Duration
	})
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

	// Добавляется для генерации случайных чисел далее в программе
	rand.Seed(time.Now().UnixNano())

	// Заполнение Alltickets
	kol := rand.Intn(15) + 60
	fmt.Printf("В продаже появилось %d билетов:\n", kol)

	for i := 0; i < kol; i++ {
		tt := 1 + rand.Intn(2) // тип поездки (в одну сторону / туда и обратно)
		dur := GetDuration()   // продолжительность полёта в одну сторону

		if tt == 1 { // если полет в олну сторону, то продолжительность равна dur
			AddToTickets(&Alltickets, spaceline[1+rand.Intn(3)], dur, ticketType[tt], 200-dur) // чем меньше время полёта, тем выше цена
		} else { // иначе полёт туда и обратно -> продолжительность и цена увеличивается в два раза
			AddToTickets(&Alltickets, spaceline[1+rand.Intn(3)], 2*dur, ticketType[2], 2*(200-dur))
		}
	}
	SortTickets(&Alltickets) // Сортировка билетов по увеличению длительности полёта
	PrintTickets(&Alltickets)
}
