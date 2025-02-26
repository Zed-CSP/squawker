import React from 'react';
import './MessageCard.css';

const MessageCard = ({ message }) => {
  return (
    <div className="message-card">
      <div className="message-header">
        <div className="user-info">
          <span className="username">@{message.username}</span>
          <p className="content" data-text={message.content}>{message.content}</p>
          <span className="time">{new Date(message.createdAt).toLocaleTimeString()}</span>
        </div>

      </div>
    </div>
  );
};

export default MessageCard; 