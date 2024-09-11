package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var heroClasses = map[string]string{
	"warrior": "warrior",
	"mage":    "mage",
	"healer":  "healer",
}

var classRandomAttack = map[string][2]int{
	"warrior": {3, 5},
	"mage":    {5, 10},
	"healer":  {-3, -1},
}
var classRandomDefence = map[string][2]int{
	"warrior": {5, 10},
	"mage":    {-2, 2},
	"healer":  {2, 5},
}

var selectedClassMessages = map[string]string{
	"warrior": "Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.",
	"mage":    "Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.",
	"healer":  "Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.",
}

var classTrainingLeadMessage = map[string]string{
	"warrior": "Воитель - отличный боец ближнего боя.",
	"mage":    "Маг - превосходный укротитель стихий.",
	"healer":  "Лекарь - чародей, способный исцелять раны.",
}

var classSkill = map[string]Skill{
	"warrior": {name: "Выносливость", points: 105},
	"mage":    {name: "Атака", points: 45},
	"healer":  {name: "Защита", points: 40},
}

func main() {
	var hero = NewHero()

	fmt.Println(
		"Приветствую тебя, искатель приключений!\n",
		"Прежде чем начать игру...\n",
		"...назови себя: ",
	)

	fmt.Scanf("%s\n", &hero.name)

	fmt.Printf("Здравствуй, %s\n", hero.name)

	fmt.Printf(
		"Сейчас твоя выносливость — %d, атака — %d и защита — %d.\n",
		hero.stats.stamina, hero.stats.damage, hero.stats.defense,
	)

	fmt.Println(
		"Ты можешь выбрать один из трёх путей силы:\n",
		"Воитель, Маг, Лекарь",
	)

	hero.class = selectHeroClass()
	hero.skill = classSkill[hero.class]

	fmt.Println(startTraining(hero))
}

func selectHeroClass() string {
	var answer string
	var selectedClass string

	for answer != "y" {
		fmt.Print(
			"Введи название персонажа, за которого хочешь играть: " +
				"Воитель — warrior, Маг — mage, Лекарь — healer: ",
		)

		fmt.Scanf("%s\n", &selectedClass)

		var message, isExist = selectedClassMessages[selectedClass]

		if isExist {
			fmt.Println(message)
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &answer)

		answer = strings.ToLower(answer)
	}

	return selectedClass
}

func startTraining(hero Hero) string {
	var randomAttack = classRandomAttack[hero.class]
	var randomDefence = classRandomDefence[hero.class]
	var leadTrainingMessage = classTrainingLeadMessage[hero.class]

	fmt.Printf("%s, ты %s.\n", hero.name, leadTrainingMessage)

	fmt.Println("Потренируйся управлять своими навыками.")

	fmt.Println(
		"Введи одну из команд: " +
			"attack — чтобы атаковать противника,\n" +
			"defence — чтобы блокировать атаку противника,\n" +
			"special — чтобы использовать свою суперсилу.",
	)

	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var command string

	for command != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &command)

		var _, isExist = heroClasses[hero.class]

		if !isExist {
			fmt.Println("неизвестный класс персонажа")
			continue
		}

		if command == "attack" {
			var damage = hero.attack() + random(randomAttack[0], randomAttack[1])

			fmt.Printf("%s нанес урон противнику равный %d.\n", hero.name, damage)
		}

		if command == "defence" {
			var blockedDamage = hero.block() + random(randomDefence[0], randomDefence[1])

			fmt.Printf("%s блокировал %d урона.\n", hero.name, blockedDamage)
		}

		if command == "special" {
			fmt.Printf(
				"%s применил специальное умение `%s %d`\n",
				hero.name, hero.skill.name, hero.skill.points,
			)
		}
	}

	return "тренировка окончена"
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

type Hero struct {
	class string
	name  string
	skill Skill
	stats Stats
}

func (hero Hero) attack() int {
	return hero.stats.damage
}

func (hero Hero) block() int {
	return hero.stats.defense
}

func NewHero() Hero {
	return Hero{
		stats: Stats{
			stamina: 80,
			damage:  5,
			defense: 10,
		},
	}
}

type Skill struct {
	name   string
	points int
}

type Stats struct {
	stamina int
	defense int
	damage  int
}
