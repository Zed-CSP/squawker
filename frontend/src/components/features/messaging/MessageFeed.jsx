import MessageCard from './MessageCard';

export default function MessageFeed({ messages, error, onRetry }) {
    return (
        <div className="feed-container">
            <div className="messages-container">
                {error && (
                    <div className="error">
                        {error}
                        <button onClick={onRetry}>Try Again</button>
                    </div>
                )}
                
                {!error && (
                    <>
                        {messages.length === 0 ? (
                            <div className="no-messages">
                                No messages yet. Be the first to post!
                            </div>
                        ) : (
                            <div className="message-list">
                                {messages.map((message) => (
                                    <MessageCard key={message.id} message={message} />
                                ))}
                            </div>
                        )}
                    </>
                )}
            </div>
        </div>
    );
} 