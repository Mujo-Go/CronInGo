package main
import (
    "fmt"
    "time"
)

type job struct{
    Month,Day,Weekday  int8
    Hour,Minute,Second int8
    Task               func(time.Time)
}

const ANY = -1

var jobs []job

func NewCronJob(month,day,weekday,hour,minute,second int8,task func(time.Time)){
    cj:=job{month,day,weekday,hour,minute,second,task}
    jobs=append(jobs,cj)
}

func NewMonthlyJob(day,hour,minute,second int8,task func(time.Time)){
    NewCronJob(ANY,day,ANY,hour,minute,second,task)
}

func NewWeeklyJob(weekday,hour,minute,second int8,task func(time.Time)){
    NewCronJob(ANY,ANY,weekday,hour,minute,second,task)
}

func NewDailyJob(hour,minute,second int8,task func(time.Time)){
    NewCronJob(ANY,ANY,ANY,hour,minute,second,task)
}

func NewHourJob(minute,second int8,task func(time.Time)){
    NewCronJob(ANY,ANY,ANY,ANY,minute,second,task)
}

func NewMinuteJob(second int8,task func(time.Time)){
    NewCronJob(ANY,ANY,ANY,ANY,ANY,second,task)
}

func (cj job) Matches(t time.Time)(ok bool){
    ok =(cj.Month==ANY || cj.Month==int8(t.Month()))    &&
        (cj.Day==ANY   || cj.Day==int8(t.Day()))        &&
        (cj.Weekday==ANY||cj.Weekday==int8(t.Weekday()))&&
        (cj.Hour==ANY  || cj.Hour==int8(t.Hour()))      &&
        (cj.Minute==ANY|| cj.Minute==int8(t.Minute()))  &&
        (cj.Second==ANY|| cj.Second==int8(t.Second()))
    return ok
}

func Task(t time.Time){
    fmt.Println(t.String())
}

func processJob(){
    for {
        now := time.Now()
        for _,j:=range jobs{
            if j.Matches(now){
                go j.Task(now)
            }
        }
        time.Sleep(time.Second)
    }
}

func main(){
    NewDailyJob(ANY,ANY,2,Task)
    NewDailyJob(ANY,ANY,4,Task)
    NewDailyJob(ANY,ANY,12,Task)
    NewDailyJob(ANY,ANY,14,Task)
    NewDailyJob(ANY,ANY,22,Task)
    NewDailyJob(ANY,ANY,24,Task)
    NewDailyJob(ANY,ANY,32,Task)
    NewDailyJob(ANY,ANY,34,Task)
    NewDailyJob(ANY,ANY,42,Task)
    NewDailyJob(ANY,ANY,44,Task)
    processJob()
}
