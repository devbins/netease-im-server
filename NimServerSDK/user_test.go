package nimserversdk

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {

	res, err := RefreshToken("hbs")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

}
