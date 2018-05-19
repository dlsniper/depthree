package depthree

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type MyDepThree string

func (m MyDepThree) Hello() {
	fmt.Printf("Hello from MyDepThree, %s", m)
	logrus.Println("called: MyDepThree.Hello()")
}
