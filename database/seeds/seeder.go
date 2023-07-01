package seed

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Seed struct {
	db *gorm.DB
}

type Method struct {
	Name      string // actually name
	Timestamp string
}

func seed(s Seed, seedMethodName string) error {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)

	if !m.IsValid() {
		return errors.New("invalid seed")
	}

	err := m.Call(nil)
	if !err[0].IsNil() {
		log.Fatalf("Cannot seed %v , got an error %v", seedMethodName, err[0])
	}

	log.Println("✔️Seed", seedMethodName, "succeed")

	return nil
}

func Execute(db *gorm.DB, seedMethodNames ...string) error {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	var seedMethods []Method
	seeds := make(map[string]reflect.Method)

	for i := 0; i < seedType.NumMethod(); i++ {
		method := seedType.Method(i)

		name := strings.Split(method.Name, "Seed")

		seedMethod := Method{
			Name:      name[0],
			Timestamp: name[1],
		}

		seedMethods = append(seedMethods, seedMethod)
		seeds[name[0]] = method
	}

	sort.Slice(seedMethods, func(p, q int) bool {
		timestamp1, err := strconv.Atoi(seedMethods[p].Timestamp)
		if err != nil {
			log.Fatalln(err)
		}

		timestamp2, err := strconv.Atoi(seedMethods[q].Timestamp)
		if err != nil {
			log.Fatalln(err)
		}

		return timestamp1 < timestamp2
	})

	// Execute all
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		for _, seedMethod := range seedMethods {
			err := seed(s, seeds[seedMethod.Name].Name)
			if err != nil {
				return err
			}
		}
	}

	// Execute only the given names
	for _, item := range seedMethodNames {
		err := seed(s, seeds[item].Name)
		if err != nil {
			return err
		}
	}
	return nil
}
