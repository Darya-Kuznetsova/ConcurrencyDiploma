package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Создаём родительский контекст:
	ctx := context.Background()
	fmt.Println("Parent", ctx)

	// WithValue контекст.
	// Указываем родительский контекст, ключ и значение:
	withValue := context.WithValue(ctx, "key", "value")
	fmt.Println("With Value", withValue)
	fmt.Println(withValue.Value("key"))

	// WithCancel context
	// Указываем родительский контекст:
	withCancel, cancel := context.WithCancel(ctx)
	fmt.Println("With Cancel", withCancel)
	// Отменяем контекст:
	cancel()

	// Получаем оповещение об отмене контекста:
	fmt.Println(withCancel.Err())

	// WithDeadline контекст.
	// Вторым параметром принимает время через 3 сек от текущего:
	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	fmt.Println("With Deadline", withDeadline)
	// Посмотреть сам дедлайн:
	fmt.Println(withDeadline.Deadline())
	// По наступлению дедлайна данные будут канал получит значение
	// и, далее, отправит их:
	fmt.Println(<-withDeadline.Done())

	// WithTimeout контекст:
	// Вторым параметром принимает время равное 2 сек:
	withTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	fmt.Println("With Timeout", withTimeout)
	fmt.Println(<-withTimeout.Done())

	// WithoutCancel контекст:
	withoutCancel := context.WithoutCancel(withTimeout)
	fmt.Println(withoutCancel)

}
