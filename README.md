# ticket_prog
Программа генерирует список билетов на полёт к другой планете

В таблице билетов четыре столбца:

    Компания, что предоставляет услуги;
    Продолжительность полёта в днях;
    Тип поездки туда и обратно или в один конец;
    Цена в миллионах долларов.

Для каждого билета случайным образом выбирается компания предоставляющая услуги;
Также случайным образом выбирается скорость корабля, и в зависимости от неё рассчитывается время полета и цена билета;
Чем быстрее корабль, тем больше цена; Линейная зависимость в виде (200 - продолжительность полёта);

Скорость варьируется в пределах от 160 до 300 тысяч км/ч;
Перед выводом на экран список билетов сортируется по возрастанию времени полёта
