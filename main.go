package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func main() {
	// Получаем количество мониторов
	numScreens := screenshot.NumActiveDisplays()
	if numScreens <= 0 {
		fmt.Println("Нет активных дисплеев")
		return
	}

	// Захватываем скриншот с первого монитора
	img, err := screenshot.CaptureDisplay(0)
	if err != nil {
		fmt.Println("Ошибка при захвате экрана:", err)
		return
	}

	// Создаем файл для сохранения скриншота
	file, err := os.Create("screenshot.png")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	// Сохраняем изображение в формате PNG
	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Ошибка при сохранении изображения:", err)
		return
	}

	fmt.Println("Скриншот сохранен как screenshot.png")
}
