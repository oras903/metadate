package  main
import ("github.com/r3labs/sse"
        "fmt"
        "encoding/json"
//       "github.com/widuu/gojson"
        "bytes"
         "os/exec"
        "github.com/buger/jsonparser")

type Status_Update_Event struct {    
   SlaveId string `form:"slaveId" json:"slaveId"`
   TaskId string `form:"taskId" json:"taskId"`
   TaskStatus string `form:"taskStatus" json:"taskStatus"`
   Message string `form:"message" json:"message"`
   AppId string `form:"appId" json:"appId"`
   Host string `form:"host" json:"host"`
//   IpAddresses [] ipaddress `form:"ipAddresses" json:"ipAddresses"`
   IpAddresses [] ipaddress 
   Ports string `form:"ports" json:"ports"`
   Version string `form:"version" json:"version"`
   EventType  string `form:"eventType" json:"eventType"`    
   Timestamp string `form:"timestamp" json:"timestamp"`
}

type ipaddress struct {
    IpAddress string  `form:"ipAddress" json:"ipAddress"`
    Protocol  string `form:"protocol" json:"protocol"`
}



func main() {
      client := sse.NewClient("http://192.168.20.11:8088/v2/events?event_type=status_update_event")

    client.Subscribe("status_update_event", func(msg *sse.Event) {
        // Got some data!
//        fmt.Println(string(msg.Data))
     fmt.Println("start")
/*     json1 := `{"slaveId":"e697c3c5-6163-42ab-9f13-977287217ea3-S0","taskId":"example.288558c0-eb78-11e7-b648-02422e407dd9","taskStatus":"TASK_KILLED","message":"Command terminated with signal Terminated","appId":"/example","host":"slave","ipAddresses":[{"ipAddress":"192.168.20.10","protocol":"IPv4"}],"ports":[31016],"version":"2017-12-28T02:38:16.484Z","eventType":"status_update_event","timestamp":"2017-12-28T02:40:09.592Z"}`        



json2 := `{"from":"en","to":"zh","ipAddresses":[{"ipAddress":"192.168.20.10","dst":"uuyt"},{"src":"tomorrow","dst":"ooiuu"}]}`

json2 := `{"from":"en","to":"zh","ipAddresses":[{"ipAddress":"192.168.20.10","dst":"uuyt"}],"version":[12123234]}`
*/
        var s Status_Update_Event
        json.Unmarshal(msg.Data, &s)
 //      fmt.Println(s.IpAddresses[0].IpAddress)
       fmt.Println(s.EventType)
       //fmt.Println(s.TaskStatus)
       fmt.Println(s.TaskId)
       fmt.Println(s.Ports)
       // fmt.Println(gojson.Json(string(msg.Data)).Get("ipAddresses").Get("ipAddress").Tostring())
//       fmt.Println("1")
//       fmt.Println(gojson.Json(json1).Getindex(7).Getindex(1).Getindex(1).Get("ipAddress").Tostring())
//       fmt.Println(gojson.Json(json2).Getindex(3).Getindex(1).Getindex(1).Get("ipAddress").Tostring())
//         fmt.Println(gojson.Json(json1).Get("ipAddresses").Arrayindex(1))  
        
//        var  st  string 
       ss, err := jsonparser.GetString(msg.Data,"ipAddresses","[0]", "ipAddress")
       sss, err := jsonparser.GetInt(msg.Data,"ports","[0]")
       fmt.Println(ss)
       fmt.Println(err)
       fmt.Println(sss)
       
       if (s.EventType == "status_update_event") && (s.TaskStatus == "TASK_FAILED" || s.TaskStatus == "TASK_KILLING" || s.TaskStatus == "TASK_KILLED" || s.TaskStatus == "TASK_LOST") {
                sc := " -X DELETE http://192.168.20.11:8500/v1/kv/"
                sc += s.TaskId
                
                c := exec.Command("curl","-XDELETE","http://192.168.20.11:8500/v1/kv/nginx/"+s.TaskId)
                if err := c.Run(); err != nil {
		 fmt.Println(c.Stderr)
	        } 
                fmt.Println(c.Stdout)
            }

      if (s.EventType == "status_update_event") && (s.TaskStatus == "TASK_RUNNING") {
                  
                 c_run := exec.Command("curl","-XPUT","-d",ss+":"+fmt.Sprint(sss),"http://192.168.20.11:8500/v1/kv/nginx/"+s.TaskId)
                 var out bytes.Buffer
                 var stderr bytes.Buffer
                 c_run.Stdout = &out
                 c_run.Stderr = &stderr
                 if err_run := c_run.Run(); err_run != nil {
  //             fmt.Println(string(c_run.Stderr))
                     fmt.Println(err_run.Error())
                     fmt.Println(stderr.String())
                 } 
                 fmt.Println(out.String())
                  
               }

       fmt.Println("stop----------------")
       fmt.Println("++++")
       fmt.Println("++++")
       fmt.Println("++++")
       fmt.Println("++++") 
       fmt.Println("++++")
      
       

 })

}
