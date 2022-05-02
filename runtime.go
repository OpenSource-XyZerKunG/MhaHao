package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/OpenSource-XyZerKunG/MhaHao-Runtime/module"
	"github.com/pkg/browser"
	"github.com/ttacon/chalk"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s - %s\n", chalk.Magenta.Color("MHAHAO"), chalk.Yellow.Color("mhahao ./hellworld.hao"))
		os.Exit(1)
	}

	fmt.Printf("%s %s, %s!\n", chalk.White.Color("Using"), chalk.Magenta.Color(module.Name), module.Version)
	dogAddress, err := os.Getwd()

	if err != nil {
		fmt.Println(chalk.Red.Color("Mha Ha Mai Jeor!"))
		os.Exit(1)
	}

	dogMailBox := path.Join(dogAddress, os.Args[1])
	fmt.Printf("%s - %s\n", chalk.Magenta.Color(module.Name), chalk.Green.Color(dogMailBox))
	dogMail, err := ioutil.ReadFile(dogMailBox)

	if err != nil {
		fmt.Println(chalk.Red.Color("Hong Hong!"))
		os.Exit(1)
	}

	partOfDog := strings.Split(string(dogMail), "\n")
	dogForWhat := strings.Split(partOfDog[0], " ")

	if len(dogForWhat) != 2 {
		fmt.Println(chalk.Red.Color("Hong Hong!"))
		os.Exit(1)
	}

	dogGoBruh := strings.TrimSpace(dogForWhat[1])
	positiveDog := []string{}

	for dogIndex := 1; dogIndex < len(partOfDog); dogIndex++ {
		dogCMD := strings.TrimSpace(partOfDog[dogIndex])

		if dogCMD == "" {
			hao(dogIndex, true)
			for {
			}
		}

		if !strings.HasPrefix(dogCMD, "//") && !strings.HasPrefix(dogCMD, "#") {
			positiveDog = append(positiveDog, dogCMD)
		} else {
			haoTaeMaiKud(dogIndex)
		}
	}

	dogFlashlight := false

	for dogIndex := 0; dogIndex < len(positiveDog); dogIndex++ {
		niceDOG := positiveDog[dogIndex]

		if strings.HasSuffix(niceDOG, ":") {
			partOfHeadDog := strings.Split(niceDOG, ":")[0]

			if dogFlashlight {
				hao(dogIndex, true)
			}

			if partOfHeadDog == dogGoBruh {
				dogFlashlight = true
				println(fmt.Sprintf("%s - %s %s", chalk.Magenta.Color(module.Name), chalk.Red.Color("Jump"), partOfHeadDog))
			}
		} else if dogFlashlight {
			partOfCMDog := strings.Split(niceDOG, ",")
			theCMDog := strings.ToUpper(partOfCMDog[0])

			if strings.HasPrefix(theCMDog, "MHAGIVEUP") {
				println(fmt.Sprintf("%s - %s %s", chalk.Magenta.Color(module.Name), chalk.Green.Color("The dog sleep!"), chalk.Yellow.Color("(●__●)")))
				os.Exit(0)
			}

			if strings.HasPrefix(theCMDog, "KRADOD") {
				dogIndex = -1
				dogFlashlight = false
				dogGoBruh = strings.TrimSpace(partOfCMDog[1])
				continue
			}

			if strings.HasPrefix(theCMDog, "BARK") {
				print(strings.ReplaceAll(strings.TrimSpace(strings.Join(partOfCMDog[1:], "")), "<<", "\n"))
				continue
			}

			println(theCMDog)
		}
	}
}

func conditionContain(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func haoTaeMaiKud(haoTime int) {
	var haoTimemiwa = []string{}
	for i := 0; i < haoTime+1; i++ {
		haoTimemiwa = append(haoTimemiwa, "Hong")
	}
	haoYee := strings.Join(haoTimemiwa, " ")
	fmt.Printf("%s - %s%s!\n", chalk.Magenta.Color(module.Name), chalk.Blue.Color(haoYee), chalk.Yellow.Color(" Tae Mai Kud"))
}

func hao(haoTime int, roll bool) {
	for {
		go func() {
			for {
				go func() {
					for {
						var haoTimemiwa = []string{}
						for i := 0; i < (haoTime+4)*(haoTime+3^4)+1; i++ {
							if roll {
								haoTimemiwa = append(haoTimemiwa, "Never gonna give you up")
								continue
							}

							haoTimemiwa = append(haoTimemiwa, "Hong")
						}

						if roll {
							browser.OpenURL("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
						}

						haoYee := chalk.Red.Color(strings.Join(haoTimemiwa, " "))
						fmt.Printf("%s!\n", haoYee)
					}
				}()
			}
		}()
	}
}
