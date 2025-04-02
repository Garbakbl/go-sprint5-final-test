package trainings

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return errors.New("Invalid input")
	}
	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return err
	}
	t.TrainingType = data[1]
	if t.TrainingType != "Бег" || t.TrainingType != "Ходьба" {
		return errors.New("Invalid training type")
	}
	t.Duration, err = time.ParseDuration(data[2])
	if err != nil {
		return err
	}
	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distanse := spentenergy.Distance(t.Steps)

	if t.Duration <= 0 {
		return "", errors.New("Duration is invalid")
	}

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var spentCalories float64
	switch t.TrainingType {
	case "Бег":
		spentCalories = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	case "Ходьба":

		spentCalories = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", errors.New("unknown training type")
	}

	return fmt.Sprintf(`Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\n 
				Скорость: %.2f км/ч\nСожгли калорий: %.2f\n`, t.TrainingType, t.Duration, distanse, meanSpeed, spentCalories), nil
}
