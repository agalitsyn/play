package main

import (
	"fmt"
	"time"

	"github.com/reugn/go-quartz/quartz"
)

func main() {
	ch := make(chan Event)

	go func() {
		time.Sleep(time.Second)
		ch <- Event{msg: "first"}
	}()

	sched := quartz.NewStdScheduler()
	sched.Start()
	sched.ScheduleJob(&SimpleJob{"Simple job 1"}, quartz.NewSimpleTrigger(time.Second*3))
	sched.ScheduleJob(&EventJob{"Background job 1", ch}, quartz.NewRunOnceTrigger(0))
	sched.ScheduleJob(&EventJob{"Background job 2", ch}, quartz.NewRunOnceTrigger(0))
	fmt.Printf("%+v\n", sched.GetJobKeys())

	time.Sleep(time.Second * 5)
	fmt.Printf("%+v\n", sched.GetJobKeys())
	sched.Stop()
	close(ch)
}

type Event struct {
	msg string
}

type EventJob struct {
	desc string
	ch   chan Event
}

func (j *EventJob) Description() string {
	return j.desc
}

func (j *EventJob) Key() int {
	return quartz.HashCode(j.Description())
}

func (j *EventJob) Execute() {
	go func() {
		for ev := range j.ch {
			fmt.Println("Executing " + j.Description() + " " + ev.msg)
		}
	}()
}

type SimpleJob struct {
	desc string
}

func (j *SimpleJob) Description() string {
	return j.desc
}

func (j *SimpleJob) Key() int {
	return quartz.HashCode(j.Description())
}

func (j *SimpleJob) Execute() {
	fmt.Println("Executing " + j.Description())
}
