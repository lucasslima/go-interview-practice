// Package challenge8 contains the solution for Challenge 8: Chat Server with Channels.
package challenge8

import (
	"errors"
	"sync"
	// Add any other necessary imports
)

// Client represents a connected chat client
type Client struct {
	// TODO: Implement this struct
	// Hint: username, message channel, mutex, disconnected flag
	Username     string
	MessageBus   chan string
	Disconnected bool
	mutex        sync.Mutex
}

// Send sends a message to the client
func (c *Client) Send(message string) {
	// TODO: Implement this method
	// Hint: thread-safe, non-blocking send
	if c.Disconnected {
		return
	}
	c.mutex.Lock()
	c.MessageBus <- message
	c.mutex.Unlock()
}

// Receive returns the next message for the client (blocking)
func (c *Client) Receive() string {
	// TODO: Implement this method
	// Hint: read from channel, handle closed channel
	if c.Disconnected {
		return ""
	}
	message, ok := <-c.MessageBus
	if !ok {
		return ""
	}
	return message
}

// ChatServer manages client connections and message routing
type ChatServer struct {
	// TODO: Implement this struct
	// Hint: clients map, mutex
	Clients  map[string]*Client
	serverMu sync.Mutex
}

// NewChatServer creates a new chat server instance
func NewChatServer() *ChatServer {
	// TODO: Implement this function
	return &ChatServer{
		Clients: make(map[string]*Client),
	}
}

// Connect adds a new client to the chat server
func (s *ChatServer) Connect(username string) (*Client, error) {
	// TODO: Implement this method
	// Hint: check username, create client, add to map
	if _, exists := s.Clients[username]; exists {
		return nil, ErrUsernameAlreadyTaken
	}
	c := &Client{
		Username:     username,
		Disconnected: false,
		MessageBus:   make(chan string),
	}
	s.Clients[username] = c
	return c, nil
}

// Disconnect removes a client from the chat server
func (s *ChatServer) Disconnect(client *Client) {
	// TODO: Implement this method
	// Hint: remove from map, close channels
	s.serverMu.Lock()
	client.mutex.Lock()
	delete(s.Clients, client.Username)
	client.Disconnected = true
	close(client.MessageBus)
	s.serverMu.Unlock()
	client.mutex.Unlock()
}

// Broadcast sends a message to all connected clients
func (s *ChatServer) Broadcast(sender *Client, message string) {
	// TODO: Implement this method
	// Hint: format message, send to all clients
}

// PrivateMessage sends a message to a specific client
func (s *ChatServer) PrivateMessage(sender *Client, recipient string, message string) error {
	// TODO: Implement this method
	// Hint: find recipient, check errors, send message
	if sender.Disconnected {
		return ErrClientDisconnected
	}
	if c, connected := s.Clients[recipient]; connected {
		c.Send(message)
	} else {
		return ErrRecipientNotFound
	}
	return nil
}

// Common errors that can be returned by the Chat Server
var (
	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrRecipientNotFound    = errors.New("recipient not found")
	ErrClientDisconnected   = errors.New("client disconnected")
	// Add more error types as needed
)
