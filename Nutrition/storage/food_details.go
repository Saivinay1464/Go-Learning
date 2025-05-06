package storage

import (
	"encoding/json"

	"github.com/Saivinay1464/Go-Learning/Nutrition/models"
)

func InitSampleData() {
	InitRedis()
	FoodDB := map[string]models.Food{
		"banana": {
			Name:     "Banana",
			Calories: 89,
			Serving:  100,
			Carbs:    22.8,
			Protein:  1.1,
			Fat:      0.3,
			Fiber:    1.6,
		},
		"chicken breast": {
			Name:     "Chicken Breast",
			Serving:  100,
			Calories: 165,
			Carbs:    0.0,
			Protein:  31.0,
			Fat:      3.6,
			Fiber:    2.4,
		},
	}

	for key, food := range FoodDB {
		data, _ := json.Marshal(food)
		RedisClient.Set(Ctx, key, data, 0)
	}
}
