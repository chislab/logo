package logo

import (
	"strings"
	"time"
)

// LOGO chislab logo
const LOGO = `
     :-                              /:         
     +++-                          :sso         
    .+++++-                      :sssss.        
    :+++++++.                  -ossssss/        
    +++++++++/.              -osssssssss               CCCCCCCCCCCCC   HHHHHHHHH     HHHHHHHHH   IIIIIIIIII      SSSSSSSSSSSSSSS 
   -+++++++++++/.          -osssssssssss-           CCC::::::::::::C   H:::::::H     H:::::::H   I::::::::I    SS:::::::::::::::S
   /+++++++++++++/       .+sssssssssssss+         CC:::::::::::::::C   H:::::::H     H:::::::H   I::::::::I   S:::::SSSSSS::::::S
   ++++++++++++++++.    -ssssssssssssssss         CC:::::::::::::::C   H:::::::H     H:::::::H   I::::::::I   S:::::SSSSSS::::::S
  :+++++++++++++/.   -:   -+sssssssssssss:       C:::::CCCCCCCC::::C   HH::::::H     H::::::HH   II::::::II   S:::::S     SSSSSSS
  ++++++++++++-   .:++ss/.   :osssssssssso      C:::::C       CCCCCC     H:::::H     H:::::H       I::::I     S:::::S            
 .+++++++++/.   -/++++sssso:   -+sssssssss.    C:::::C                   H:::::H     H:::::H       I::::I     S:::::S            
 :+++++++-   .:+++++++sssssss+.   :ossssss/    C:::::C                   H::::::HHHHH::::::H       I::::I      S::::SSSS         
 +++++:.   -/+++++++++ssssssssso:   -+sssss    C:::::C                   H:::::::::::::::::H       I::::I       SS::::::SSSSS    
.+++-   .:++++++++++++ssssssssssss+.   :oss-   C:::::C                   H:::::::::::::::::H       I::::I         SSS::::::::SS  
/:.   -/++++++++++++++sssssssssssssso:   -++   C:::::C                   H::::::HHHHH::::::H       I::::I            SSSSSS::::S 
   .:+++++++++++++++++sssssssssssssssss+.      C:::::C                   H:::::H     H:::::H       I::::I                 S:::::S
   :++++++++++++++++++ssssssssssssssssss/       C:::::C       CCCCCC     H:::::H     H:::::H       I::::I                 S:::::S
     ./+++++++++++++++ssssssssssssssso-          C:::::CCCCCCCC::::C   HH::::::H     H::::::HH   II::::::II   SSSSSSS     S:::::S
        :+++++++++++++sssssssssssss/.             CC:::::::::::::::C   H:::::::H     H:::::::H   I::::::::I   S::::::SSSSSS:::::S
          ./++++++++++sssssssssso-                  CCC::::::::::::C   H:::::::H     H:::::::H   I::::::::I   S:::::::::::::::SS 
             :++++++++ssssssss/.                       CCCCCCCCCCCCC   HHHHHHHHH     HHHHHHHHH   IIIIIIIIII    SSSSSSSSSSSSSSS   
               ./+++++ssssso-                   
                  :+++sss/.                     
                    ./+-                       
`

// Print print chislab logo
func Print(d time.Duration) {
	for _, line := range strings.Split(LOGO, "\n") {
		println(line)
		time.Sleep(d)
	}
}
