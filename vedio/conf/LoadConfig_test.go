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
	LoadConfig()
	fmt.Println(ConfigYaml)
}
