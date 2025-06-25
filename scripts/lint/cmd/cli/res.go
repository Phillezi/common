package cli

import (
	"fmt"

	"github.com/Phillezi/common/utils/or"
)

const (
	glint = `
        .__  .__        __   
   ____ |  | |__| _____/  |_ 
  / ___\|  | |  |/    \   __\
 / /_/  >  |_|  |   |  \  |  
 \___  /|____/__|___|  /__|  
/_____/              \/      `
)

var (
	version = "v0.0.1"
)

func banner() {
	fmt.Printf("%s\nVersion:\t\t\t\t\t%s\n", glint, or.Or(version, "not available"))
}
