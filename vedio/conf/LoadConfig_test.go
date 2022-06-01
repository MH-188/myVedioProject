/**
* @Author: 18209
* @Description:
* @File:  LoadConfigTest
* @Version: 1.0.0
* @Date: 2022/5/27 20:36
 */

package conf

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ConfigYaml)
}
