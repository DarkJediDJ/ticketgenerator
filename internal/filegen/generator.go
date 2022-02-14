package ticket

import (
	"fmt"
	"log"
	"os"

	pb "github.com/darkjedidj/ticketgenerator/api/proto"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func Generate(in *pb.TicketRequset) error {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("Ticket ID: %v", in.Id))
		})
	})
	m.Row(11, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("Ticket Time: %v", in.Time))
		})
	})
	m.Row(12, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("Ticket Price: %v", in.Price))
		})
	})
	m.Row(13, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("Ticket Seat: %v", in.Seat))
		})
	})
	m.Row(13, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("Ticket Title: %s", in.Title))
		})
	})
	err := m.OutputFileAndClose(fmt.Sprintf("internal/tickets/%v.pdf", in.Id))
	if err != nil {
		return err
	}
	return nil
}

func Delete(id int64) {
	err := os.Remove(fmt.Sprintf("internal/tickets/%v.pdf", id))
	if err != nil {
		log.Fatalf("%v", err)
	}
}
