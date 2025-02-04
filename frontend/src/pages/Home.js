import React from 'react';

function Home() {
  return (
    <div className="home">
      <h1>Squawker</h1>
      <div className="message-feed">
        {/* We'll implement this next */}
        <p>Message feed coming soon...</p>
      </div>
      <div className="message-input">
        <textarea 
          placeholder="What's on your mind? (One message per day)"
          maxLength={280}
        />
        <button>Squawk!</button>
      </div>
    </div>
  );
}

export default Home; 