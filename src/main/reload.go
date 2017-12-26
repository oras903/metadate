package main 


import (
    "github.com/gin-gonic/gin"
	"fmt"
	"os/exec"
)



type Status_Update_Event struct {    
   EventType     string `form:"eventType" json:"eventType"`    
   Timestamp string `form:"timestamp" json:"timestamp"`
   SlaveId string `form:"slaveId" json:"slaveId"`
   TaskId string `form:"taskId" json:"taskId"`
   TaskStatus string `form:"taskStatus" json:"taskStatus"`
   AppId string `form:"appId" json:"appId"`
   Host string `form:"host" json:"host"`
   Ports string `form:"ports" json:"ports"`
   Version string `form:"version" json:"version"`
}



func main() {    

   router := gin.Default()
   
   router.POST("/callback", func(c *gin.Context) {
   
       var update Status_Update_Event
       
       if c.Bind(&update) != nil {
           
            if update.TaskStatus == "TASK_FAILED" || update.TaskStatus == "TASK_KILLING" || update.TaskStatus == "TASK_KILLED" || update.TaskStatus == "TASK_LOST" {
                s := " -X DELETE http://192.168.20.11:8500/v1/kv/"
                s += update.TaskId
                
               	f, err := exec.Command("curl",s).Output()
	              if err == nil {
	                  c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	                	c.JSON(200, gin.H{"status": "delete sacling nginx"})
	              }
	              fmt.Println(string(f))
	             
               
              
            }
           
           
               if update.TaskStatus == "TASK_RUNNING" {
                  h := " -X PUT -d '"
                  h += update.Host
                  h += ":"
                  h += update.Ports
                  h += "' http://192.168.20.11:8500/v1/kv/nginx/"
                  h += update.TaskId
                  
                  
                  
               	 f_run, err_run := exec.Command("curl",h).Output()
	               if err_run == nil {
	                   c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	                	 c.JSON(200, gin.H{"status": "add sacling nginx"})
	               }
	               fmt.Println(string(f_run))                  
            
                
               }
             
            
       
       } 
      
   })
   
   
  router.Run(":8898") // listen and serve on 0.0.0.0:8898


}