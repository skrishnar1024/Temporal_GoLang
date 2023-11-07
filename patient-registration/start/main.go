package main

import (
	"context"
	"log"
	"patient-registration/app"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "Patient_Creation-108",
		TaskQueue: app.RegistrationTaskQueue,
	}

	var data app.PatientDetails

	data.Name = "Alex"
	data.Age = "30"
	data.Contact = "123456789"
	data.Address = "Texas"

	we, err := c.ExecuteWorkflow(context.Background(), options, app.PatientRegistrationWorkflow, data)

	if err != nil {
		log.Fatalln("Unable to complete workflow\n", err)
	}

	var workflowOutput app.PatientDetails
	err = we.Get(context.Background(), &workflowOutput)

	log.Fatalln("workflow workflowOutput: ", workflowOutput)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

}
