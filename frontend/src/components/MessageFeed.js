import React, { useState, useEffect } from 'react';
import { useAuth } from '../context/AuthContext';

function MessageFeed() {
  const [messages, setMessages] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const { user } = useAuth();

  useEffect(() => {
    fetchMessages();
  }, []);

  const fetchMessages = async () => {
    try {
      const response = await fetch('http://localhost:8080/messages', {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
      });

      if (response.ok) {
        const data = await response.json();
        setMessages(data);
      } else {
        setError('Failed to fetch messages');
      }
    } catch (err) {
      setError('Something went wrong');
    } finally {
      setLoading(false);
    }
  };

  if (loading) return <div>Loading messages...</div>;
  if (error) return <div className="error">{error}</div>;

  return (
    <div className="message-feed">
      {messages.map(message => (
        <div key={message.id} className="message">
          <div className="message-header">
            <span className="username">{message.username}</span>
            <span className="timestamp">
              {new Date(message.createdAt).toLocaleString()}
            </span>
          </div>
          <p className="message-content">{message.content}</p>
        </div>
      ))}
    </div>
  );
}

export default MessageFeed; 