package attachments

import "time"

type AttachmentDB struct {
	attachments map[string][]string //

}

func ConnectAttachmentDB() *AttachmentDB {
	attachments := map[string][]string{
		"sess_1":  {"doc1.pdf", "image1.png"},
		"sess_2":  {"document.docx"},
		"sess_3":  {"dock3.pdf", "source.mp3"},
		"sess_5":  {"movie.mp4"},
		"sess_10": {"file.txt"},
	}
	return &AttachmentDB{attachments: attachments}
}

func (db *AttachmentDB) GetAttachments(sessionID string) []string {
	if sessionID == "" {
		return []string{}
	}
	time.Sleep(300 * time.Millisecond)
	return db.attachments[sessionID]
}
