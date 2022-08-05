package main

import (
	"fmt"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

type Forum struct {
    id uuid.UUID
    Closed bool
    Subject string
    Description string
}

// StartDiscussion create a new discussion
func (f *Forum) StartDiscussion(author *Author, subject string) (*Discussion, error) {
    // check if forum is closed
    if f.Closed {
        return nil, fmt.Errorf("Forum is closed")
    }

    dId, err := uuid.NewRandom()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    // create a new discussion and return it
    return &Discussion{
        id: dId,
        ForumId: f.id,
        Author: author,
        Subject: subject,
    }, nil
}

type Discussion struct {
    id uuid.UUID
    ForumId uuid.UUID
    Author *Author
    Subject string
}

type Author struct {
    id uuid.UUID
    Name string
}


func main() {
    aId, _ := uuid.NewRandom()
    author := &Author{
        id: aId,
        Name: "John Doe",
    }

    fId, _ := uuid.NewRandom()
    forum := &Forum{
        id: fId,
        Closed: false,
        Subject: "Go",
        Description: "For discussion about Go",
    }

    d, _ := forum.StartDiscussion(author, "Go is a programming language")
    fmt.Printf("%+v\n", d)
}