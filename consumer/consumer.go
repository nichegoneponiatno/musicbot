package consumer

import (
	"log"
	telegram "musicbot/client"
	"time"
)

type Consumer struct {
	telegram *telegram.Client
	offset   int
	limit    int
}

func New(offset, limit int, telegram *telegram.Client) *Consumer {
	return &Consumer{
		offset:   offset,
		limit:    limit,
		telegram: telegram,
	}
}

func (c *Consumer) WorkingCycle() error {
	for {
		events, err := c.telegram.Updates(c.offset, c.limit)
		if err != nil {
			log.Println(err)

			continue
		}

		if len(events) == 0 {
			continue
		}

		for i := range events {
			c.telegram.SendMessages(events[i].Message.Chat.ID, events[i].Message.Text)
			if err != nil {
				log.Println(err)

				continue
			}
		}

		c.offset = events[len(events)-1].ID + 1

		time.Sleep(5 * time.Second)

	}
}
