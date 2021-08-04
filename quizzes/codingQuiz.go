package quizzes

import (
	"github.com/xavierhazzardadmin/GoQuiz/Utils"
)

var correctQuestions int = 0

func CodingQuestion1() {

	Utils.AssertAnswer("How many bits are there in a byte?", 8, &correctQuestions)
}
