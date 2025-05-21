package repository

import (
	"database/sql"
	"fmt"
)

const (
	addChatQuery           = `insert into chats (chat_id) values ($1)`
	removeChatQuery        = `delete from chats where chat_id = $1`
	addLinkQuery           = `insert into links (link) values ($1)`
	removeLinkQuery        = `delete from links where link = $1`
	addChatLinkQuery       = `insert into chat_link (chat_id, link_id) values ($1, $2)`
	removeChatLinkQuery    = `delete from chat_link where chat_id = $1 and link_id = $2`
	getLinksQuery          = `select link from chats join links using(link_id) where chat_id = $1`
	getIdByLinkQuery       = `select link_id from links where link = $1`
	getAllLinksQuery       = `select link from links`
	isChatPresentQuery     = `select chat_id from chats where chat_id = $1`
	isLinkPresentQuery     = `select link_id from links where link = $1`
	isChatLinkPresentQuery = `select chat_id from chat_link where chat_id = $1 and link_id = $2`
)

type PostgresqlLinkTracker struct {
	db *sql.DB
}

func NewPostgresqlLinkTracker(connStr string) (*PostgresqlLinkTracker, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening connection to postgresql link tracker: %w", err)
	}
	return &PostgresqlLinkTracker{db: db}, nil
}

func (s *PostgresqlLinkTracker) AddChat(chatId int64) error {
	_, err := s.db.Exec(addChatQuery, chatId)
	if err != nil {
		return fmt.Errorf("error adding chat: %w", err)
	}
	return nil
}

func (s *PostgresqlLinkTracker) RemoveChat(chatId int64) error {
	_, err := s.db.Exec(removeChatQuery, chatId)
	if err != nil {
		return fmt.Errorf("error removing chat: %w", err)
	}
	return nil
}

func (s *PostgresqlLinkTracker) AddLink(link string) error {
	_, err := s.db.Exec(addLinkQuery, link)
	if err != nil {
		return fmt.Errorf("error adding link: %w", err)
	}
	return nil
}

func (s *PostgresqlLinkTracker) RemoveLink(link string) error {
	_, err := s.db.Exec(removeLinkQuery, link)
	if err != nil {
		return fmt.Errorf("error removing link: %w", err)
	}
	return nil
}

func (s *PostgresqlLinkTracker) AddChatLink(chatId int64, linkId int64) error {
	_, err := s.db.Exec(addChatLinkQuery, chatId, linkId)
	if err != nil {
		return fmt.Errorf("error adding chat link: %w", err)
	}
	return nil
}

func (s *PostgresqlLinkTracker) RemoveChatLink(chatId int64, linkId int64) error {
	_, err := s.db.Exec(removeChatLinkQuery, chatId, linkId)
	if err != nil {
		return fmt.Errorf("error removing chat link: %w", err)
	}
	return nil
}

func (s *PostgresqlLinkTracker) GetLinks(chatId int64) ([]string, error) {
	rows, err := s.db.Query(getLinksQuery, chatId)
	if err != nil {
		return nil, fmt.Errorf("error getting links: %w", err)
	}
	defer rows.Close()

	links := make([]string, 0)
	for rows.Next() {
		var link string
		if err := rows.Scan(&link); err != nil {
			return nil, fmt.Errorf("error getting links: %w", err)
		}
		links = append(links, link)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error getting links: %w", err)
	}
	return links, nil
}

func (s *PostgresqlLinkTracker) GetIdByLink(link string) (int64, error) {
	row := s.db.QueryRow(getIdByLinkQuery, link)
	var linkId int64
	if err := row.Scan(&linkId); err != nil {
		return 0, fmt.Errorf("error getting link linkId: %w", err)
	}
	return linkId, nil
}

func (s *PostgresqlLinkTracker) IsChatPresent(chatId int64) (bool, error) {
	rows, err := s.db.Query(isChatPresentQuery, chatId)
	if err != nil {
		return false, fmt.Errorf("error checking if chat exists: %w", err)
	}
	defer rows.Close()

	return rows.Next(), nil
}

func (s *PostgresqlLinkTracker) IsLinkPresent(link string) (bool, error) {
	rows, err := s.db.Query(isLinkPresentQuery, link)
	if err != nil {
		return false, fmt.Errorf("error checking if link exists: %w", err)
	}
	defer rows.Close()
	return rows.Next(), nil
}

func (s *PostgresqlLinkTracker) IsChatLinkPresent(chatId int64, linkId int64) (bool, error) {
	rows, err := s.db.Query(isChatLinkPresentQuery, chatId, linkId)
	if err != nil {
		return false, fmt.Errorf("error checking if chatlink exists: %w", err)
	}
	defer rows.Close()
	return rows.Next(), nil
}

func (s *PostgresqlLinkTracker) GetAllLinks() ([]string, error) {
	rows, err := s.db.Query(getAllLinksQuery)
	if err != nil {
		return nil, fmt.Errorf("error getting all links: %w", err)
	}
	defer rows.Close()
	links := make([]string, 0)
	for rows.Next() {
		var link string
		if err := rows.Scan(&link); err != nil {
			return nil, fmt.Errorf("error getting all links: %w", err)
		}
		links = append(links, link)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error getting all links: %w", err)
	}
	return links, nil
}
