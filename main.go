package main

import (
	"fmt"
	"sort"
)

type Student struct {
	ID         int
	name       string
	surname    string
	patronymic string
	Grades     []int
}

func (s Student) FullName() string {
	return s.surname + " " + s.name + " " + s.patronymic
}

func (s *Student) AddGrade(grade int) {
	if grade >= 1 && grade <= 5 {
		s.Grades = append(s.Grades, grade)
	} else {
		fmt.Println("Оценка должна быть от 1 до 5.")
	}
}

func (s *Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0.0
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

type Gradebook struct {
	students map[int]*Student
	nextID   int
}

func NewGradebook() *Gradebook {
	return &Gradebook{
		students: make(map[int]*Student),
		nextID:   1,
	}
}

func (g *Gradebook) AddStudent(name, surname, patronymic string) int {
	id := g.nextID
	g.nextID++
	student := &Student{
		ID:         id,
		name:       name,
		surname:    surname,
		patronymic: patronymic,
		Grades:     []int{},
	}
	g.students[id] = student
	fmt.Printf("Студент '%s' добавлен с ID %d\n", student.FullName(), id)
	return id
}

func (g *Gradebook) AddGradeToStudent(id int, grade int) {
	if student, exists := g.students[id]; exists {
		student.AddGrade(grade)
	} else {
		fmt.Println("Студент с таким ID не найден.")
	}
}

func (g *Gradebook) PrintAllStudents() {
	fmt.Println("\nСписок всех студентов:")
	for _, student := range g.students {
		avg := student.AverageGrade()
		fmt.Printf("ID: %d, Имя: %s, Оценки: %v, Средний балл: %.2f\n",
			student.ID, student.FullName(), student.Grades, avg)
	}
}

func (g *Gradebook) StudentsBelowGrade(threshold float64) []*Student {
	var result []*Student
	for _, student := range g.students {
		if student.AverageGrade() < threshold {
			result = append(result, student)
		}
	}
	return result
}

func (g *Gradebook) SortByAverage(ascending bool) []*Student {
	students := make([]*Student, 0, len(g.students))
	for _, s := range g.students {
		students = append(students, s)
	}

	if ascending {
		sort.Slice(students, func(i, j int) bool {
			return students[i].AverageGrade() < students[j].AverageGrade()
		})
	} else {
		sort.Slice(students, func(i, j int) bool {
			return students[i].AverageGrade() > students[j].AverageGrade()
		})
	}

	return students
}

func main() {
	gradebook := NewGradebook()

	for {
		fmt.Println("\n--- Учет успеваемости студентов ---")
		fmt.Println("1. Добавить студента")
		fmt.Println("2. Добавить оценку студенту")
		fmt.Println("3. Показать всех студентов")
		fmt.Println("4. Показать студентов с баллом ниже заданного")
		fmt.Println("5. Сортировка по среднему баллу (по убыванию)")
		fmt.Println("6. Сортировка по среднему баллу (по возрастанию)")
		fmt.Println("7. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Ошибка ввода. Пожалуйста, введите число от 1 до 7.")
			continue
		}

		switch choice {
		case 1:
			var name, surname, patronymic string
			fmt.Println("Введите ФИО (Фамилия Имя Отчество): ")
			n, err := fmt.Scanln(&surname, &name, &patronymic)
			if err != nil || n != 3 {
				fmt.Println("Ошибка при вводе ФИО. Убедитесь, что вы ввели три слова, разделенных пробелом.")
				var discard string
				fmt.Scanln(&discard)
				continue
			}
			gradebook.AddStudent(name, surname, patronymic)

		case 2:
			var id int
			var grade int
			gradebook.PrintAllStudents()
			for {
				fmt.Print("Введите ID студента: ")
				_, err := fmt.Scanln(&id)
				if err != nil {
					fmt.Println("Ошибка ввода ID. Пожалуйста, введите целое число.")
					var discard string
					fmt.Scanln(&discard)
					continue
				}

				if _, exists := gradebook.students[id]; !exists {
					fmt.Printf("Студент с ID %d не найден. Попробуйте снова.\n", id)
					continue
				}

				break
			}
			fmt.Print("Введите оценку (1-5): ")
			_, err := fmt.Scanln(&grade)
			if err != nil {
				fmt.Println("Ошибка ввода оценки. Пожалуйста, введите целое число от 1 до 5.")
				var discard string
				fmt.Scanln(&discard)
				continue
			}

			if grade < 1 || grade > 5 {
				fmt.Println("Оценка должна быть от 1 до 5.")
				continue
			}
			gradebook.AddGradeToStudent(id, grade)

		case 3:
			gradebook.PrintAllStudents()

		case 4:
			var threshold float64
			fmt.Print("Введите пороговое значение (например, 4): ")
			n, err := fmt.Scanln(&threshold)
			if err != nil || n != 1 {
				fmt.Println("Ошибка ввода порога. Пожалуйста, введите число.")
				var discard string
				fmt.Scanln(&discard)
				continue
			}
			students := gradebook.StudentsBelowGrade(threshold)
			fmt.Printf("\nСтуденты с баллом ниже %.2f:\n", threshold)
			for _, s := range students {
				fmt.Printf("ID: %d, Имя: %s, Средний балл: %.2f\n", s.ID, s.FullName(), s.AverageGrade())
			}

		case 5:
			sorted := gradebook.SortByAverage(false)
			fmt.Println("Сортировка по убыванию среднего балла:")
			for _, s := range sorted {
				fmt.Printf("ID: %d, Имя: %s, Средний балл: %.2f\n", s.ID, s.FullName(), s.AverageGrade())
			}

		case 6:
			sorted := gradebook.SortByAverage(true)
			fmt.Println("Сортировка по возрастанию среднего балла:")
			for _, s := range sorted {
				fmt.Printf("ID: %d, Имя: %s, Средний балл: %.2f\n", s.ID, s.FullName(), s.AverageGrade())
			}

		case 7:
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}
