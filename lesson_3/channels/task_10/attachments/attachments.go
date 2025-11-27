package attachments

import "time"

type Attachments struct {
	SessionId string
	Files     []string
}

type AttDB struct {
	files map[string][]string
}

func ConnectAttachmentDB() *AttDB {
	return &AttDB{
		files: map[string][]string{
			"sess_1":  {"doc1.pdf", "image1.png"},
			"sess_2":  {"document.docx"},
			"sess_3":  {"dock3.pdf", "source.mp3"},
			"sess_5":  {"movie.mp4"},
			"sess_10": {"file.txt"},
		},
	}
}

func (db *AttDB) GetAttachments(sessionID string) *Attachments {

	time.Sleep(300 * time.Millisecond)
	return &Attachments{
		SessionId: sessionID,
		Files:     db.files[sessionID],
	}
}
