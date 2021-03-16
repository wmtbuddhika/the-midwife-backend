package scheduler

import (
	"back-end/modules/database"
	"back-end/modules/models"
	"back-end/modules/notification"
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func StartScheduler()  {
	c := cron.New()
	fmt.Println("Starting Scheduler ... ")

	c.AddFunc("0 0 * * *", func() {
		vaccines, err := database.GetChildrenVaccineByDate(time.Now())

		if err == nil {
			for _, vaccine := range vaccines {
				message := fmt.Sprintf("<h1>Vaccine Notification</h1>" +
					"<p>Your Child, %s %s is need to get following Vaccine. <br>" +
					"Please visit your Midwife Officer. <br><br> " +
					"Vaccine Code : %s <br><br>" +
					"Vaccine Name : %s </p><br><br>" +
					"Regards, <br>" +
					"<strong>Midwife System</strong>", vaccine.ChildFirstName, vaccine.ChildLastName, vaccine.VaccineCode, vaccine.VaccineName)
				text := "Your Child, %s %s is need to get following Vaccine. \n Vaccine Code : %s \n Vaccine Name : %s \nRegards, \nMidwife System"
				notification.SendEmail(models.SiteEmail, []string{vaccine.Email}, models.VaccineReminderSubject, message)
				notification.SendSMS(vaccine.PhoneNo, text)
			}
		}
	})
	//c.AddFunc("@every 1m", func() {
	//	fmt.Println("Every 1 minute ")
	//})
	c.Start()
}
