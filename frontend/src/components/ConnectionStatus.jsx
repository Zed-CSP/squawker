import { useState, useEffect } from 'react';

export default function ConnectionStatus({ isConnected, onClick }) {
    const [cursorVisible, setCursorVisible] = useState(true);
    
    useEffect(() => {
        const cursorInterval = setInterval(() => {
            setCursorVisible(prev => !prev);
        }, 1000);
        
        return () => clearInterval(cursorInterval);
    }, []);

    return (
        <div 
            className="connection-status" 
            style={{
                padding: '10px',
                margin: '10px 0',
                borderRadius: '4px',
                backgroundColor: isConnected ? 'rgba(0, 255, 0, 0.1)' : 'rgba(255, 0, 0, 0.1)',
                border: `1px solid ${isConnected ? 'var(--neon-blue)' : 'var(--neon-pink)'}`,
                color: isConnected ? 'var(--neon-blue)' : 'var(--neon-pink)',
                textAlign: 'center',
                fontFamily: 'monospace',
                cursor: 'pointer'
            }} 
            onClick={onClick}
        >
            {isConnected ? (
                <span>LIVE CONNECTION ESTABLISHED!{cursorVisible ? '▇' : '_'}</span>
            ) : (
                <span>CONNECTING TO SERVER...{cursorVisible ? '▇' : '_'}</span>
            )}
        </div>
    );
} 