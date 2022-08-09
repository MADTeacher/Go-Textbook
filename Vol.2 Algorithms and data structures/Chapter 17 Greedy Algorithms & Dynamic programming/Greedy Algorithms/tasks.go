package main

import (
	"fmt"
	"sort"
)

type Task struct {
	ID          int
	Deadline    int
	Profit      int
	Description string
}

// Для сортировки по прибыльности
type TaskProfit []Task

func (t TaskProfit) Len() int {
	return len(t)
}

func (t TaskProfit) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TaskProfit) Less(i, j int) bool {
	return t[i].Profit > t[j].Profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ScheduleTasks(tasks []Task, maxDay int) (int, []Task) {
	profit := 0
	slot := make([]*Task, maxDay)
	for i := 0; i < maxDay; i++ {
		slot[i] = nil
	}

	sort.Sort(TaskProfit(tasks))
	for idx, task := range tasks {
		for j := task.Deadline - 1; j >= 0; j-- {
			if j < maxDay && slot[j] == nil {
				slot[j] = &tasks[idx]
				profit += task.Profit
				break
			}
		}
	}

	resultTasks := []Task{}
	for _, it := range slot {
		if it != nil {
			resultTasks = append(resultTasks, *it)
		}
	}
	return profit, resultTasks
}

func main() {
	tasks := []Task{
		{1, 9, 15, "Уборка"},
		{2, 2, 2, "Запилить баг"},
		{3, 5, 18, "Пофиксить баг"},
		{4, 7, 1, "Сварить кофе"},
		{5, 4, 25, "Участие в совещании"},
		{6, 2, 20, "One-to-one"},
		{7, 5, 8, "Ревью кода"},
		{8, 7, 10, "Погладить одежду"},
		{9, 4, 12, "Разработать фичу"},
		{10, 3, 5, "Кричать: 'За что???'"},
	}
	profit, optimalTasks := ScheduleTasks(tasks, 15)
	fmt.Printf("Profit: %d\n", profit)
	fmt.Println("Optimal tasks:")
	for _, it := range optimalTasks {
		fmt.Printf("%+v\n", it)
	}
}
